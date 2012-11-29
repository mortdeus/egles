package egl

/*
#cgo pkg-config: egl
#include <stdlib.h>
#include <EGL/egl.h>


NativeDisplayType DefaultDisplay(){
	return EGL_DEFAULT_DISPLAY;
}
EGLContext NoContext(){
	return EGL_NO_CONTEXT;
}
EGLDisplay *NoDisplay(){
	return EGL_NO_DISPLAY;
}
EGLSurface *NoSurface(){
	return EGL_NO_SURFACE;
}
EGLint DontCare(){
	return EGL_DONT_CARE;
} 

*/
import "C"
import "unsafe"

type (
	Int               int32
	Boolean           bool
	Enum              uint
	Config            uintptr
	Context           uintptr
	Display           uintptr
	Surface           uintptr
	ClientBuffer      uintptr
	NativeDisplayType C.EGLNativeDisplayType
	NativeWindowType  C.EGLNativeWindowType
	NativePixmapType  C.EGLNativePixmapType
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
func DefaultDisplay() *NativeDisplayType {
	disp := NativeDisplayType(C.DefaultDisplay())
	return &disp
}

func NoSurface() Surface {
	return Surface(unsafe.Pointer(C.NoSurface()))
}

func NoDisplay() Display {
	return Display(unsafe.Pointer(C.NoDisplay()))
}

func NoContext() Context {
	return Context(unsafe.Pointer(C.NoContext()))
}
func DontCare() Int {
	return Int(C.DontCare())
}
