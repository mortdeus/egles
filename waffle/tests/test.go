package main

import (
	"github.com/mortdeus/egles/waffle"
	"time"
)

type GLClamp float32
type GLBitfield uint32

const GL_COLOR_BUFFER_BIT = 0x00004000

type (
	glClearColor func(r, g, b, a float32)
	glClear      func(mask uint32)
)

func main() {
	disp := new(waffle.Display)
	conf := new(waffle.Config)
	win := new(waffle.Window)
	ctx := new(waffle.Context)

	initAttrib := (waffle.AttributeList)([]waffle.Attribute{
		waffle.PLATFORM, waffle.PLATFORM_X11_EGL, 0})
	confAttrib := (waffle.AttributeList)([]waffle.Attribute{
		waffle.CONTEXT_API, waffle.CONTEXT_OPENGL_ES2,

		waffle.RED_SIZE, 8,
		waffle.BLUE_SIZE, 8,
		waffle.RED_SIZE, 8,
		0,
	})
	width, height := int32(320), int32(240)

	waffle.Init((*int32)(&initAttrib[0]))

	disp.Connect("")

	// Exit if OpenGL ES2 is unsupported.
	if ok := disp.SupportContextAPI(waffle.CONTEXT_OPENGL_ES2); !(ok || waffle.CanOpenDL(waffle.DL_OPENGL_ES2)) {

		panic(waffle.ErrorToString(waffle.GetErrorCode()))
	}

	// Get GL functions.
	glClearColor = waffle.DLSym(waffle.DL_OPENGL_ES2, "glClearColor")
	glClear = waffle.DLSym(waffle.DL_OPENGL_ES2, "glClear")

	conf.Choose(disp, confAttrib)
	win.Create(conf, width, height)
	ctx.Create(nil)

	_ = win.Show()
	_ = waffle.MakeCurrent(disp, win, ctx)

	// Draw red.
	glClearColor(1.0, 0.0, 0.0, 1.0)
	glClear(GL_COLOR_BUFFER_BIT)
	win.SwapBuffers()
	time.Sleep(1 * time.Second)

	// Draw green.
	glClearColor(1.0, 0.0, 0.0, 1.0)
	glClearColor(0.0, 1.0, 0.0, 1.0)
	glClear(GL_COLOR_BUFFER_BIT)
	win.SwapBuffers()
	time.Sleep(1 * time.Second)

	// Draw blue.
	glClearColor(0.0, 0.0, 1.0, 1.0)
	glClear(GL_COLOR_BUFFER_BIT)
	win.SwapBuffers()
	time.Sleep(1 * time.Second)

	// Clean up.
	_ = waffle.MakeCurrent(disp, nil, nil)
	_ = win.Destroy()
	_ = ctx.Destroy()
	_ = conf.Destroy()
	_ = disp.Disconnect()
}
