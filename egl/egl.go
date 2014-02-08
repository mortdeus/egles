package egl

/*
#cgo pkg-config: egl
#include <EGL/egl.h>
#include <EGL/eglplatform.h>
*/
import "C"
import (
	"unsafe"
)

func Initialize(d Display) bool {
	return goBool(C.eglInitialize(
		C.EGLDisplay(d),
		(*C.EGLint)(unsafe.Pointer(&Version.Maj)),
		(*C.EGLint)(unsafe.Pointer(&Version.Min))))
}
func Terminate(d Display) bool {
	return goBool(C.eglTerminate(C.EGLDisplay(d)))
}
func GetDisplay(d NativeDisplay) Display {
	return Display(C.eglGetDisplay(C.EGLNativeDisplayType(d)))
}
func QueryString(d Display, name int) string {
	return C.GoString(C.eglQueryString(C.EGLDisplay(d), C.EGLint(name)))
}
func DestroySurface(d Display, s Surface) bool {
	return goBool(C.eglDestroySurface(C.EGLDisplay(d), C.EGLSurface(s)))
}
func SwapInterval(d Display, inv int) bool {
	return goBool(C.eglSwapInterval(C.EGLDisplay(d), C.EGLint(inv)))
}
func DestroyContext(d Display, c Context) bool {
	return goBool(C.eglDestroyContext(C.EGLDisplay(d), C.EGLContext(c)))
}
func GetCurrentSurface(readdraw int) Surface {
	return Surface(C.eglGetCurrentSurface(C.EGLint(readdraw)))
}
func QuerySurface(d Display, val []int, attr int, s Surface) bool {
	return goBool(C.eglQuerySurface(
		C.EGLDisplay(d), C.EGLSurface(s), C.EGLint(attr),
		(*C.EGLint)(unsafe.Pointer(&val[0]))))
}
func GetConfigs(d Display, confs []Config, confSz int, nConf []int) bool {
	return goBool(C.eglGetConfigs(
		C.EGLDisplay(d), (*C.EGLConfig)(unsafe.Pointer(&confs[0])),
		C.EGLint(confSz), (*C.EGLint)(unsafe.Pointer(&nConf[0]))))
}

func GetConfigAttrib(d Display, conf Config, attr int, val []int) bool {
	return goBool(C.eglGetConfigAttrib(
		C.EGLDisplay(d), C.EGLConfig(conf), C.EGLint(attr),
		(*C.EGLint)(unsafe.Pointer(&val[0]))))
}
func ChooseConfig(d Display, atrribs []int, confs []Config, confSz int, nConf []int) bool {
	return goBool(C.eglChooseConfig(
		C.EGLDisplay(d), (*C.EGLint)(unsafe.Pointer(&atrribs[0])),
		(*C.EGLConfig)(&confs[0]), C.EGLint(confSz),
		(*C.EGLint)(unsafe.Pointer(&nConf[0]))))
}
func CreateContext(d Display, conf Config, shared Context, attribs []int) Context {
	return Context(C.eglCreateContext(
		C.EGLDisplay(d), C.EGLConfig(conf), C.EGLContext(shared),
		(*C.EGLint)(unsafe.Pointer(&attribs[0]))))
}

func CreateWindowSurface(d Display, conf Config, win NativeWindow, attribs []int) Surface {
	return Surface(C.eglCreateWindowSurface(
		C.EGLDisplay(d), C.EGLConfig(conf), C.EGLNativeWindowType(uintptr(win)),
		(*C.EGLint)(unsafe.Pointer(&attribs[0]))))
}
func CreatePbufferSurface(d Display, conf Config, attribs []int) Surface {
	return Surface(C.eglCreatePbufferSurface(
		C.EGLDisplay(d), C.EGLConfig(conf),
		(*C.EGLint)(unsafe.Pointer(&attribs[0]))))
}
func CreatePixmapSurface(d Display, conf Config, pixmap NativePixmap, attribs []int) Surface {
	return Surface(C.eglCreatePixmapSurface(
		C.EGLDisplay(d), C.EGLConfig(conf), C.EGLNativePixmapType(uintptr(pixmap)),
		(*C.EGLint)(unsafe.Pointer(&attribs[0]))))
}
func CreatePbufferFromClientBuffer(
	d Display, buftyp uint, conf Config, buf ClientBuffer, attribs []int) Surface {
	return Surface(C.eglCreatePbufferFromClientBuffer(
		C.EGLDisplay(d), C.EGLenum(buftyp),
		C.EGLClientBuffer(buf), C.EGLConfig(conf),
		(*C.EGLint)(unsafe.Pointer(&attribs[0]))))
}
func SurfaceAttrib(d Display, s Surface, attr int, val int) bool {
	return goBool(C.eglSurfaceAttrib(
		C.EGLDisplay(d), C.EGLSurface(s), C.EGLint(attr), C.EGLint(val)))
}
func BindTexImage(d Display, s Surface, buf int) bool {
	return goBool(C.eglBindTexImage(C.EGLDisplay(d), C.EGLSurface(s), C.EGLint(buf)))
}
func ReleaseTexImage(d Display, s Surface, buf int) bool {
	return goBool(C.eglReleaseTexImage(C.EGLDisplay(d), C.EGLSurface(s), C.EGLint(buf)))
}
func MakeCurrent(d Display, draw Surface, read Surface, c Context) bool {
	return goBool(C.eglMakeCurrent(
		C.EGLDisplay(d), C.EGLSurface(draw), C.EGLSurface(read), C.EGLContext(c)))
}
func QueryContext(d Display, c Context, attr int, val []int) bool {
	return goBool(C.eglQueryContext(
		C.EGLDisplay(d), C.EGLContext(c), C.EGLint(attr),
		(*C.EGLint)(unsafe.Pointer(&val[0]))))
}
func CopyBuffers(d Display, s Surface, target NativePixmap) bool {
	return goBool(C.eglCopyBuffers(
		C.EGLDisplay(d), C.EGLSurface(s),
		C.EGLNativePixmapType(uintptr(target))))
}
func SwapBuffers(d Display, s Surface) bool {
	return goBool(C.eglSwapBuffers(C.EGLDisplay(d), C.EGLSurface(s)))
}

func BindAPI(api uint) bool      { return goBool(C.eglBindAPI(C.EGLenum(api))) }
func WaitNative(engine int) bool { return goBool(C.eglWaitNative(C.EGLint(engine))) }
func QueryAPI() uint             { return uint(C.eglQueryAPI()) }
func WaitClient() bool           { return goBool(C.eglWaitClient()) }
func WaitGL() bool               { return goBool(C.eglWaitGL()) }
func ReleaseThread() bool        { return goBool(C.eglReleaseThread()) }
func GetCurrentDisplay() Display { return Display(C.eglGetCurrentDisplay()) }
func GetError() error            { return Error(C.eglGetError()) }
