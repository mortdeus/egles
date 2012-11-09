package egl

/*
#cgo pkg-config: egl
#include <stdlib.h>
#include <EGL/egl.h>


void *DefaultDisplay(){
	return EGL_DEFAULT_DISPLAY;
}
void *NoContext(){
	return EGL_NO_CONTEXT;
}
void *NoDisplay(){
	return EGL_NO_DISPLAY;
}
void *NoSurface(){
	return EGL_NO_SURFACE;
}
EGLint DontCare(){
	return EGL_DONT_CARE;
} 

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
func DefaultDisplay() NativeDisplayType {
	return NativeDisplayType(C.DefaultDisplay())
}

func NoSurface() Surface {
	return Surface(C.NoSurface())
}

func NoDisplay() Display {
	return Display(C.NoDisplay())
}

func NoContext() Context {
	return Context(C.NoContext())
}
func DontCare() Int {
	return Int(C.DontCare())
}
