package egl

/*
#include <stdlib.h>
#include <EGL/egl.h>
#include <EGL/eglplatform.h>
#include <EGL/eglext.h>
*/
import "C"
import "unsafe"

type (
	Config        unsafe.Pointer
	Context       unsafe.Pointer
	Display       unsafe.Pointer
	Surface       unsafe.Pointer
	ClientBuffer  unsafe.Pointer
	NativeDisplay unsafe.Pointer
	NativeWindow  unsafe.Pointer
	NativePixmap  unsafe.Pointer
)

var (
	DEFAULT_DISPLAY NativeDisplay
	NO_CONTEXT      Context
	NO_DISPLAY      Display
	NO_SURFACE      Surface
)

var Version struct{ Maj, Min int }

func goBool(n C.EGLBoolean) bool {
	return n == 1
}
func eglBool(n bool) C.EGLBoolean {
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
