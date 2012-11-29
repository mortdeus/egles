package main

//#include<stdlib.h>
import "C"

import (
	"fmt"
	"github.com/mortdeus/egles/egl"
	"github.com/mortdeus/egles/gles/2"
)

var (
	attr = [...]egl.Int{
		egl.WIDTH, 500,
		egl.HEIGHT, 500,
		egl.NONE,
	}
	numConfig egl.Int
	max       egl.Int
	min       egl.Int
	ctx       egl.Context
	pbuf      egl.Surface
	configs   egl.Config
	b         egl.Boolean
)

func run() {
	disp := egl.GetDisplay(egl.DEFAULT_DISPLAY)
	defer egl.Terminate(disp)

	if ok := egl.Initialize(disp, &max, &min); !ok {
		panic("Initialize() failed")
		return
	}

	fmt.Printf("EGL Version: %v, %v\n", max, min)

	if ok := egl.GetConfigs(disp, nil, 0, &numConfig); !ok {
		panic("GetConfigs() failed")
	}
	configs := make([]egl.Config, int(numConfig))

	if ok := egl.GetConfigs(disp, &configs[0], numConfig, &numConfig); !ok {
		panic("GetConfigs() failed")
	}

	egl.BindAPI(egl.OPENGL_ES_API)
	ctx = egl.CreateContext(disp, configs[0], egl.NO_CONTEXT, nil)
	if ctx == egl.NO_CONTEXT {
		panic("CreateContext() failed")
		return
	}

	pbuf = egl.CreatePbufferSurface(disp, configs[0], &attr[0])
	configs = nil
	if ok := egl.MakeCurrent(disp, pbuf, pbuf, ctx); !ok {
		panic("MakeCurrent() failed")
		return
	}
	if ok := egl.MakeCurrent(disp, egl.NO_SURFACE, egl.NO_SURFACE, ctx); !ok {
		panic("MakeCurrent() failed")
		return
	}
	_ = egl.DestroySurface(disp, pbuf)
	_ = egl.DestroyContext(disp, ctx)

	return
}

func main() {
	run()
	println("Done")
}
