package main

//#include<stdlib.h>
import "C"

import (
	"github.com/mortdeus/egles"
	"github.com/mortdeus/egles/egl"
	gl "github.com/mortdeus/egles/gles/2"
)

var (
	attr = []egl.Int{
		egl.RED_SIZE, 1,
		egl.GREEN_SIZE, 1,
		egl.BLUE_SIZE, 1,
		egl.NONE,
	}
	window       *egles.Window
	display      egl.Display
	config       egl.Config
	context      egl.Context
	surface      egl.Surface
	nativeWindow egl.NativeWindowType
	numConfig    egl.Int
)

func main() {
	initialize()
	draw()
	println("Done")
}
func initialize() {
	egles.Init()

	display = egl.GetDisplay(egl.DEFAULT_DISPLAY)
	defer egl.Terminate(display)
	
	_ = egl.Initialize(display, nil, nil)
	_ = egl.ChooseConfig(display, attr, &config, 1, &numConfig)
	context = egl.CreateContext(display, config, egl.NO_CONTEXT, nil)
	surface = egl.CreateWindowSurface(display, config, nativeWindow, nil)
	_ = egl.MakeCurrent(display, surface, surface, context)
	gl.ClearColor(1.0, 0, 0, 1.0)
	gl.UseProgram(program())
}
func draw() {

}
