package main

//#include<stdlib.h>
import "C"

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/mortdeus/egles/egl"
	gl "github.com/mortdeus/egles/egl/gles/2"
	"log"
	"time"
)

var (
	attr = []egl.Int{
		egl.RED_SIZE, 1,
		egl.GREEN_SIZE, 1,
		egl.BLUE_SIZE, 1,
		egl.NONE,
	}
	X            *xgbutil.XUtil
	display      egl.Display
	config       egl.Config
	context      egl.Context
	surface      egl.Surface
	nativeWindow egl.NativeWindowType
	numConfig    egl.Int
	Done         = make(chan bool, 1)
	i            = 0
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

	display = egl.GetDisplay(egl.DEFAULT_DISPLAY)
	if ok := egl.Initialize(display, nil, nil); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.ChooseConfig(display, attr, &config, 1, &numConfig); !ok {
		egl.LogError(egl.GetError())
	}
	context = egl.CreateContext(display, config, egl.NO_CONTEXT, nil)
	surface = egl.CreateWindowSurface(display, config,
		egl.NativeWindowType(uintptr(X.RootWin())), nil)

	if ok := egl.MakeCurrent(display, surface, surface, context); !ok {
		egl.LogError(egl.GetError())
	}
	gl.ClearColor(0.0, 0, 0, 1.0)
	gl.UseProgram(program())
}
func update() {
	time.Sleep(time.Millisecond * 10)
}
func draw() {

}
func cleanup() {
	egl.Terminate(display)
}
