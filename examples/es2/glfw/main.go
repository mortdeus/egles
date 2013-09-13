package main

import (
	"flag"
	"fmt"
	glfw "github.com/go-gl/glfw3"
	gl "github.com/mortdeus/egles/es2"
	mgl "github.com/mortdeus/mathgl"
	"runtime"
	"sync/atomic"
	"time"
)

var debug = flag.Bool("d", false, "Prints FPS to stdout")

func init() {
	runtime.LockOSThread()
	runtime.GOMAXPROCS(2)
}

var (
	Title         = "egles"
	Width, Height = 640, 480

	POSITION = [8]float32{
		0.0, 0.0, 1.0, 0.0,
		0.0, 1.0, 1.0, 1.0}
	COLOR = [16]float32{
		1.0, 0.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
		0.0, 0.0, 1.0, 1.0,
		1.0, 1.0, 0.0, 1.0,
	}
	attrPos, attrColor, uMVP int
	ModelView, Projection    mgl.Mat4f

	MVP = make(chan mgl.Mat4f, 1)
	fps = new(uint64)
)

func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}
func main() {
	flag.Parse()
	glfw.SetErrorCallback(errorCallback)
	if !glfw.Init() {
		println("glfw init failure")
	}
	defer glfw.Terminate()
	glfw.WindowHint(glfw.ClientApi, glfw.OpenglEsApi)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)
	window, err := glfw.CreateWindow(Width, Height, Title, nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	window.SetSizeCallback(reshape)
	window.SetKeyCallback(keyEvent)
	gl.Viewport(0, 0, Width, Height)
	initScene()

	if *debug {
		go func() {
			tick := time.Tick(1 * time.Second)
			for {
				<-tick
				fmt.Printf("FPS:%v\r", atomic.LoadUint64(fps))
				atomic.StoreUint64(fps, 0)
			}
		}()
	}
	for !window.ShouldClose() {
		drawScene()
		window.SwapBuffers()
		glfw.PollEvents()
		atomic.AddUint64(fps, 1)
	}
}
func initScene() {
	halfW, halfH := float32(Width)*0.5, float32(Height)*0.5

	Projection = mgl.Ortho2D(-halfW, halfW, -halfH, halfH).Mul4(
		mgl.Translate3D(-halfW, -halfH, 0))

	ModelView = mgl.Ident4f().Mul4(mgl.Scale3D(100, 100, 0))
	MVP <- Projection.Mul4(ModelView)

	gl.Disable(gl.DEPTH_TEST)
	gl.DepthMask(false)

	program := Program(FragmentShader(fsh), VertexShader(vsh))
	gl.UseProgram(uint(program))
	uMVP = gl.GetUniformLocation(program, "uMVP")
	attrPos = gl.GetAttribLocation(program, "pos")
	attrColor = gl.GetAttribLocation(program, "color")

	gl.EnableVertexAttribArray(uint(attrPos))
	gl.EnableVertexAttribArray(uint(attrColor))
	gl.ClearColor(0.5, 0.5, 0.5, 1.0)
	gl.VertexAttribPointer(uint(attrPos), 2, gl.FLOAT, false, 0, uintptr(gl.Void(&POSITION[0])))
	gl.VertexAttribPointer(uint(attrColor), 4, gl.FLOAT, false, 0, uintptr(gl.Void(&COLOR[0])))

}
func drawScene() {
	select {
	case m := <-MVP:
		gl.UniformMatrix4fv(uMVP, 1, false, m[:])
	default:
		break
	}
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
	gl.Flush()
	gl.Finish()
}
func reshape(w *glfw.Window, width, height int) {
	gl.Viewport(0, 0, width, height)
	halfW, halfH := float32(width)*0.5, float32(height)*0.5
	Projection = mgl.Ortho2D(-halfW, halfW, -halfH, halfH).Mul4(
		mgl.Translate3D(-halfW, -halfH, 0))

	go func(m mgl.Mat4f) { MVP <- m }(Projection.Mul4(ModelView))
}

func keyEvent(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch key {
	case glfw.KeyEscape:
		w.SetShouldClose(true)
	case glfw.KeyLeft:
		if mods == glfw.ModShift {
			ModelView = ModelView.Mul4(mgl.HomogRotate3DY(5))
		} else {
			ModelView = ModelView.Mul4(mgl.Translate3D(-0.1, 0, 0))
		}
	case glfw.KeyRight:
		if mods == glfw.ModShift {
			ModelView = ModelView.Mul4(mgl.HomogRotate3DY(-5))
		} else {
			ModelView = ModelView.Mul4(mgl.Translate3D(0.1, 0, 0))
		}
	case glfw.KeyUp:
		if mods == glfw.ModShift {
			ModelView = ModelView.Mul4(mgl.HomogRotate3DX(5))
		} else {
			ModelView = ModelView.Mul4(mgl.Translate3D(0, 0.1, 0))
		}
	case glfw.KeyDown:
		if mods == glfw.ModShift {
			ModelView = ModelView.Mul4(mgl.HomogRotate3DX(-5))
		} else {
			ModelView = ModelView.Mul4(mgl.Translate3D(0, -0.1, 0))
		}
	case glfw.KeyMinus:
		ModelView = ModelView.Mul4(mgl.Translate3D(0, 0, -0.1))
	case glfw.KeyEqual:
		ModelView = ModelView.Mul4(mgl.Translate3D(0, 0, 0.1))
	}
	go func(m mgl.Mat4f) { MVP <- m }(Projection.Mul4(ModelView))

}
