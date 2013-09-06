package main

import (
	"flag"
	"fmt"
	glfw "github.com/go-gl/glfw3"
	gl "github.com/mortdeus/egles/es2"
	glm "github.com/mortdeus/mathgl"
	"runtime"
	"sync/atomic"
	"time"
)

var (
	Title    = "egles"
	Width    = 640
	Height   = 480
	vertices = [12]float32{
		-0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, 0.0, 1.0,
		0.0, 0.5, 0.0, 1.0,
	}
	colors = [12]float32{
		1.0, 0.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
		0.0, 0.0, 1.0, 1.0,
	}
	verticesArrayBuffer, colorsArrayBuffer uint32
	attrPos, attrColor                     int32
	View, Model, Projection, MVP           glm.Mat4f
	uMVP                                   int32
	fps                                    = new(uint64)
)
var debug = flag.Bool("d", false, "Prints FPS to stdout")

func init() {
	runtime.LockOSThread()
	runtime.GOMAXPROCS(2)
}

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
	gl.Viewport(0, 0, gl.Sizei(Width), gl.Sizei(Height))
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
	Projection = glm.Perspective(45.0, float32(Width)/float32(Height), 0.1, 100.0)
	View = glm.LookAt(0.0, 0.0, 5.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)
	Model = glm.Ident4f()

	program := Program(FragmentShader(fsh), VertexShader(vsh))
	gl.UseProgram(program)
	uMVP = gl.GetUniformLocation(program, "uMVP")
	attrPos = gl.GetAttribLocation(program, "pos")
	attrColor = gl.GetAttribLocation(program, "color")

	gl.GenBuffers(1, &verticesArrayBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, verticesArrayBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.SizeiPtr(len(vertices))*4,
		gl.Void(&vertices[0]), gl.STATIC_DRAW)

	gl.GenBuffers(1, &colorsArrayBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorsArrayBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.SizeiPtr(len(colors))*4,
		gl.Void(&colors[0]), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(uint32(attrPos))
	gl.EnableVertexAttribArray(uint32(attrColor))
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

}
func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	MVP = Projection.Mul4(View).Mul4(Model)
	gl.UniformMatrix4fv(int32(uMVP), 1, false, &MVP[0])
	gl.BindBuffer(gl.ARRAY_BUFFER, verticesArrayBuffer)
	gl.VertexAttribPointer(uint32(attrPos), 4, gl.FLOAT, false, 0, 0)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorsArrayBuffer)
	gl.VertexAttribPointer(uint32(attrColor), 4, gl.FLOAT, false, 0, 0)

	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.Flush()
	gl.Finish()

}
func reshape(w *glfw.Window, width, height int) {
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
	Projection = glm.Perspective(45.0, float32(width)/float32(height), 0.1, 100.0)

}

func keyEvent(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	switch key {
	case glfw.KeyEscape:
		w.SetShouldClose(true)
	case glfw.KeyLeft:
		if mods == glfw.ModShift {
			Model = Model.Mul4(glm.HomogRotate3DY(5))
		} else {
			Model = Model.Mul4(glm.Translate3D(-0.1, 0, 0))
		}
	case glfw.KeyRight:
		if mods == glfw.ModShift {
			Model = Model.Mul4(glm.HomogRotate3DY(-5))
		} else {
			Model = Model.Mul4(glm.Translate3D(0.1, 0, 0))
		}
	case glfw.KeyUp:
		if mods == glfw.ModShift {
			Model = Model.Mul4(glm.HomogRotate3DX(5))
		} else {
			Model = Model.Mul4(glm.Translate3D(0, 0.1, 0))
		}
	case glfw.KeyDown:
		if mods == glfw.ModShift {
			Model = Model.Mul4(glm.HomogRotate3DX(-5))
		} else {
			Model = Model.Mul4(glm.Translate3D(0, -0.1, 0))
		}
	case glfw.KeyMinus:
		Model = Model.Mul4(glm.Translate3D(0, 0, -0.1))
	case glfw.KeyEqual:
		Model = Model.Mul4(glm.Translate3D(0, 0, 0.1))
	}

}
