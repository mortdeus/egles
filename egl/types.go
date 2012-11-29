package egl

/*
#cgo pkg-config: egl
#include <stdlib.h>
#include <EGL/egl.h>
#include <EGL/eglplatform.h>
*/
import "C"

type (
	Int               int32
	Boolean           bool
	Enum              uint
	Config            uintptr
	Context           uintptr
	Display           uintptr
	Surface           uintptr
	ClientBuffer      uintptr
	NativeDisplayType uintptr
	NativeWindowType  uintptr
	NativePixmapType  uintptr
)

func goBoolean(n C.EGLBoolean) (b Boolean) {
	b = Boolean(TRUE == n)
	return
}
func eglBoolean(n Boolean) C.EGLBoolean {
	var b int
	if n == true {
		b = TRUE
	}
	return C.EGLBoolean(b)
}

/*
func ProcAdress(proc string) uintptr {

}
*/
