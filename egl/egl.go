package egl

/*
#cgo pkg-config: egl
#include <stdlib.h>
#include <EGL/egl.h>

const EGLNativeDisplayType DefaultDisplay = EGL_DEFAULT_DISPLAY; 
const EGLSurface NoSurface = EGL_NO_SURFACE;
const EGLDisplay NoDisplay = EGL_NO_DISPLAY;
const EGLContext NoContext = EGL_NO_CONTEXT;

// This is a generic function pointer type, whose name indicates it must
// be cast to the proper type *and calling convention* before use.

	typedef void (*__eglMustCastToProperFunctionPointerType)(void);

// 	Now, define eglGetProcAddress using the generic function ptr. type 

	EGLAPI __eglMustCastToProperFunctionPointerType EGLAPIENTRY

 	eglGetProcAddress(const char *procname);


*/
import "C"

import (
	"unsafe"
)

type (
	Int               C.EGLint
	Boolean           C.EGLBoolean
	Enum              C.EGLenum
	Config            C.EGLConfig
	Context           C.EGLContext
	Display           C.EGLDisplay
	Surface           C.EGLSurface
	ClientBuffer      C.EGLClientBuffer
	NativeDisplayType C.NativeDisplayType
	NativeWindowType  C.NativeWindowType
	NativePixmapType  C.NativePixmapType
	ProcAdress        unsafe.Pointer
)

func GetProcAddress(procName string) unsafe.Pointer {
	s := C.CString(procName)
	defer C.free(unsafe.Pointer(s))
	return unsafe.Pointer(C.eglGetProcAddress(s))
}

var (
	DefaultDisplay = NativeDisplayType(C.DefaultDisplay)
	NoSurface      = Surface(C.NoSurface)
	NoDisplay      = Display(C.NoDisplay)
	NoContext      = Context(C.NoContext)
)

func GetError() Int {
	return (Int)(C.eglGetError())
}
func GetDisplay(displayID NativeDisplayType) Display {
	return Display(C.eglGetDisplay(displayID))
}
func Initialize(disp Display, major, minor *Int) Boolean {

	return Boolean((C.eglInitialize(
		C.EGLDisplay(disp),
		(*C.EGLint)(major),
		(*C.EGLint)(minor))))
}
func Terminate(disp Display) Boolean {
	return Boolean(C.eglTerminate(C.EGLDisplay(disp)))
}
func QueryString(disp Display, name Int) string {
	cstr := C.eglQueryString(C.EGLDisplay(disp), C.EGLint(name))
	s := string(C.GoString(cstr))
	C.free(unsafe.Pointer(cstr))
	return s
}
func GetConfigs(disp Display, configs *Config, configSize Int, numConfig *Int) Boolean {

	return Boolean(C.eglGetConfigs(
		C.EGLDisplay(disp),
		(*C.EGLConfig)(configs),
		C.EGLint(configSize),
		(*C.EGLint)(numConfig)))
}
func ChooseConfig(
	disp Display, atrribList *Int,
	configs *Config, configSize Int, numConfig *Int) Boolean {

	return Boolean(C.eglChooseConfig(
		C.EGLDisplay(disp),
		(*C.EGLint)(atrribList),
		(*C.EGLConfig)(configs),
		C.EGLint(configSize),
		(*C.EGLint)(numConfig)))
}
func GetConfigAttrib(disp Display, config Config, attribute Int, value *Int) Boolean {
	return Boolean(C.eglGetConfigAttrib(
		C.EGLDisplay(disp),
		C.EGLConfig(config),
		C.EGLint(attribute),
		(*C.EGLint)(value)))
}
func CreateWindowSurface(disp Display, config Config, win NativeWindowType, attribList *Int) Surface {
	return Surface(C.eglCreateWindowSurface(
		C.EGLDisplay(disp),
		C.EGLConfig(config),
		C.EGLNativeWindowType(win),
		(*C.EGLint)(attribList)))
}
func CreatePbufferSurface(disp Display, config Config, attribList *Int) Surface {
	return Surface(C.eglCreatePbufferSurface(
		C.EGLDisplay(disp),
		C.EGLConfig(config),
		(*C.EGLint)(attribList)))
}
func CreatePixmapSurface(disp Display, config Config, pixmap NativePixmapType, attribList *Int) Surface {
	return Surface(C.eglCreatePixmapSurface(
		C.EGLDisplay(disp),
		C.EGLConfig(config),
		C.EGLNativePixmapType(pixmap),
		(*C.EGLint)(attribList)))
}
func DestroySurface(disp Display, surface Surface) Boolean {
	return Boolean(C.eglDestroySurface(
		C.EGLDisplay(disp),
		C.EGLSurface(surface)))

}
func QuerySurface(disp Display, surface Surface, attribute Int, value *Int) Boolean {
	return Boolean(C.eglQuerySurface(
		C.EGLDisplay(disp),
		C.EGLSurface(surface),
		C.EGLint(attribute),
		(*C.EGLint)(value)))
}

func BindAPI(api Enum) Boolean {
	return Boolean(C.eglBindAPI(C.EGLenum(api)))
}
func QueryAPI() Enum {
	return Enum(C.eglQueryAPI())
}

func WaitClient() Boolean {
	return Boolean(C.eglWaitClient())
}

func ReleaseThread() Boolean {
	return Boolean(C.eglReleaseThread())
}

func CreatePbufferFromClientBuffer(
	disp Display, buftype Enum,
	buffer ClientBuffer, config Config, attribList *Int) Surface {

	return Surface(C.eglCreatePbufferFromClientBuffer(
		C.EGLDisplay(disp),
		C.EGLenum(buftype),
		C.EGLClientBuffer(buffer),
		C.EGLConfig(config),
		(*C.EGLint)(attribList)))
}

func SurfaceAttrib(disp Display, surface Surface, attribute Int, value Int) Boolean {
	return Boolean(C.eglSurfaceAttrib(
		C.EGLDisplay(disp),
		C.EGLSurface(surface),
		C.EGLint(attribute),
		C.EGLint(value)))
}
func BindTexImage(disp Display, surface Surface, buffer Int) Boolean {
	return Boolean(C.eglBindTexImage(
		C.EGLDisplay(disp),
		C.EGLSurface(surface),
		C.EGLint(buffer)))
}
func ReleaseTexImage(disp Display, surface Surface, buffer Int) Boolean {
	return Boolean(C.eglReleaseTexImage(
		C.EGLDisplay(disp),
		C.EGLSurface(surface),
		C.EGLint(buffer)))
}

func SwapInterval(disp Display, interval Int) Boolean {
	return Boolean(C.eglSwapInterval(
		C.EGLDisplay(disp),
		C.EGLint(interval)))
}

func CreateContext(disp Display, config Config, shareContext Context, attribList *Int) Context {
	return Context(C.eglCreateContext(
		C.EGLDisplay(disp),
		C.EGLConfig(config),
		C.EGLContext(shareContext),
		(*C.EGLint)(attribList)))
}
func DestroyContext(disp Display, ctx Context) Boolean {
	return Boolean(C.eglDestroyContext(
		C.EGLDisplay(disp),
		C.EGLContext(ctx)))
}
func MakeCurrent(disp Display, draw Surface, read Surface, ctx Context) Boolean {
	return Boolean(C.eglMakeCurrent(
		C.EGLDisplay(disp),
		C.EGLSurface(draw),
		C.EGLSurface(read),
		C.EGLContext(ctx)))
}

func GetCurrentContext() Context {
	return Context(C.eglGetCurrentContext())
}
func GetCurrentSurface(readdraw Int) Surface {
	return Surface(C.eglGetCurrentSurface(C.EGLint(readdraw)))
}
func GetCurrentDisplay() Display {
	return Display(C.eglGetCurrentDisplay())
}
func QueryContext(disp Display, ctx Context, attribute Int, value *Int) Boolean {
	return Boolean(C.eglQueryContext(
		C.EGLDisplay(disp),
		C.EGLContext(ctx),
		C.EGLint(attribute),
		(*C.EGLint)(value)))
}

func WaitGL() Boolean {
	return Boolean(C.eglWaitGL())
}
func WaitNative(engine Int) Boolean {
	return Boolean(C.eglWaitNative(C.EGLint(engine)))
}
func SwapBuffers(disp Display, surface Surface) Boolean {
	return Boolean(C.eglSwapBuffers(
		C.EGLDisplay(disp),
		C.EGLSurface(surface)))
}
func CopyBuffers(disp Display, surface Surface, target NativePixmapType) Boolean {
	return Boolean(C.eglCopyBuffers(
		C.EGLDisplay(disp),
		C.EGLSurface(surface),
		C.EGLNativePixmapType(target)))
}
