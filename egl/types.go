package egl

/*
#cgo pkg-config: egl
#include <stdlib.h>
#include <EGL/egl.h>
#include <EGL/eglplatform.h>
#include <EGL/eglext.h>

const NativeDisplayType EGLDefaultDisplay = EGL_DEFAULT_DISPLAY;
const EGLDisplay EGLNoDisplay = EGL_NO_DISPLAY;
*/
import "C"

type (
	Enum              uint32
	Config            uintptr
	Context           uintptr
	Display           uintptr
	Surface           uintptr
	ClientBuffer      uintptr
	NativeDisplayType C.NativeDisplayType
	NativeWindowType  C.NativeWindowType
	NativePixmapType  C.NativePixmapType
)

var DefaultDisplay = C.EGLDefaultDisplay
var NoDisplay = C.EGLNoDisplay

func goBoolean(n C.EGLBoolean) bool {
	return n == 1
}
func eglBoolean(n bool) C.EGLBoolean {
	var b int
	if n == true {
		b = 1
	}
	return C.EGLBoolean(b)
}

/*
func ProcAdress(proc string) uintptr {

}
*/
