package main

//#include<stdlib.h>
import "C"

import (
	"fmt"
	"github.com/mortdeus/egles/egl"
	//"github.com/mortdeus/egles/gl"
	"time"
	"unsafe"
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
	disp := egl.GetDisplay(egl.DefaultDisplay)
	defer egl.Terminate(disp)

	if ok := egl.Initialize(disp, &max, &min); ok == egl.FALSE {
		println("Initialize() failed")
		return
	}

	fmt.Printf("EGL Version: %v, %v\n", max, min)

	str := egl.QueryString(disp, egl.VENDOR)
	fmt.Printf("EGL Vendor:  %s\n", str)

	if b = egl.GetConfigs(disp, nil, 0, &numConfig); b == egl.FALSE {
		println("GetConfigs() failed")
	}

	configs := make([]egl.Config, int(numConfig))

	if b = egl.GetConfigs(disp, &configs[0], numConfig, &numConfig); b == egl.FALSE {
		println("GetConfigs() failed")
		return
	}

	egl.BindAPI(egl.OPENGL_API)

	ctx = egl.CreateContext(disp, configs[0], egl.NoContext, nil)
	if ctx == egl.NoContext {
		println("CreateContext() failed")
		return
	}

	pbuf = egl.CreatePbufferSurface(disp, configs[0], &attr[0])

	C.free(unsafe.Pointer(&configs[0]))
	configs = nil
	b = egl.MakeCurrent(disp, pbuf, pbuf, ctx)
	if b == egl.FALSE {
		println("MakeCurrent() failed")
		return

	}
	if b = egl.MakeCurrent(disp, egl.NoSurface, egl.NoSurface, egl.NoContext); b == egl.FALSE {
		println("MakeCurrent() failed")
		return
	}

	time.Sleep(7 * time.Second)
	return
}

func main() {
	run()
}
