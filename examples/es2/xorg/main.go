package main

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/mortdeus/egles/egl"
	gl "github.com/mortdeus/egles/egl/gles/2"
	"github.com/mortdeus/mathgl"
	"log"
	"math"
	"time"
)

var (
	attr = []int32{
		egl.RED_SIZE, 1,
		egl.GREEN_SIZE, 1,
		egl.BLUE_SIZE, 1,
		egl.RENDERABLE_TYPE,
		egl.OPENGL_ES2_BIT,
		egl.NONE,
	}
	ctxAttr = []int32{egl.NONE}

	X            *xgbutil.XUtil
	display      egl.Display
	config       egl.Config
	context      egl.Context
	surface      egl.Surface
	nativeWindow egl.NativeWindowType
	numConfig    int32
	Done         = make(chan bool, 1)
	vid          int32
	redraw       = true
)

func main() {
	initialize()
	defer cleanup()
	for {
		select {
		case <-Done:
			return
		default:
			update()
			draw()

		}
	}
}
func initialize() {
	var err error
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	mousebind.Initialize(X)
	newWindow(X)
	go xevent.Main(X)
	println(egl.DefaultDisplay)
	display = egl.GetDisplay(egl.DefaultDisplay)
	if ok := egl.Initialize(display, nil, nil); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.ChooseConfig(display, attr, &config, 1, &numConfig); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(display, config, egl.NATIVE_VISUAL_ID, &vid); !ok {
		egl.LogError(egl.GetError())
	}
	egl.BindAPI(egl.OPENGL_ES_API)

	context = egl.CreateContext(display, config, egl.NoContext, nil)

	surface = egl.CreateWindowSurface(display, config,
		egl.NativeWindowType(uintptr(X.RootWin())), nil)

	var val int32
	if ok := egl.QuerySurface(display, &val, egl.WIDTH, surface); !ok {
		egl.LogError(egl.GetError())
	}

	if ok := egl.QuerySurface(display, &val, egl.HEIGHT, surface); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(display, config, egl.SURFACE_TYPE, &val); !ok {
		egl.LogError(egl.GetError())
	}
	gl.ClearColor(0.0, 0, 0, 1.0)

	p := Program(FragmentShader(fsh), VertexShader(vsh))
	gl.UseProgram(p)
	gl.BindAttribLocation(p, uint32(attrPos), "pos")
	gl.BindAttribLocation(p, uint32(attrColor), "color")
	uMatrix = int32(gl.GetUniformLocation(p, "modelviewProjection"))

}
func update() {
	time.Sleep(time.Millisecond * 10)
}

var (
	viewRotX float32
	viewRotY float32

	uMatrix   int32 = -1
	attrPos   uint32
	attrColor uint32 = 1
)

func draw() {
	if redraw {
		verts := [3][2]float32{{-1, -1}, {1, -1}, {0, -1}}
		colors := [3][3]float32{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
		var mat, rot, scale mathgl.Mat4

		makeZRotMatrix(float32(viewRotX), &rot)
		makeScaleMatrix(.5, .5, .5, &scale)
		rot.Multiply(&scale)
		mat = rot
		gl.UniformMatrix4fv(uMatrix, 1, false, (*float32)(&mat[0]))

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		{
			gl.VertexAttribPointer(attrPos, 2, gl.FLOAT, false, 0, gl.Void(&verts[0]))
			gl.VertexAttribPointer(attrColor, 3, gl.FLOAT, false, 0, gl.Void(&colors[0]))
			gl.EnableVertexAttribArray(attrPos)
			gl.EnableVertexAttribArray(attrColor)

			gl.DrawArrays(gl.TRIANGLES, 0, 3)

			gl.DisableVertexAttribArray(attrPos)
			gl.DisableVertexAttribArray(attrColor)
		}
		redraw = false
	}
}
func cleanup() {
	egl.DestroySurface(display, surface)
	egl.DestroyContext(display, context)
	egl.Terminate(display)
}

func reshape(width, height int) {
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
}

func makeZRotMatrix(angle float32, m *mathgl.Mat4) {
	c := float32(math.Cos(float64(angle) * math.Pi / 180.0))
	s := float32(math.Sin(float64(angle) * math.Pi / 180.0))
	m.Identity()
	m[0] = c
	m[1] = s
	m[4] = -s
	m[5] = c
}
func makeScaleMatrix(xs, ys, zs float32, m *mathgl.Mat4) {
	m[0] = xs
	m[5] = ys
	m[10] = zs
	m[15] = 1.0
}
