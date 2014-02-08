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
	return Display(C.eglGetDisplay(
		C.EGLNativeDisplayType(d)))
}
func QueryString(d Display, name int) string {
	return C.GoString(C.eglQueryString(
		C.EGLDisplay(d),
		C.EGLint(name)))
}
func DestroySurface(d Display, s Surface) bool {
	return goBool(C.eglDestroySurface(
		C.EGLDisplay(d),
		C.EGLSurface(s)))
}
func SwapInterval(d Display, inv int) bool {
	return goBool(C.eglSwapInterval(
		C.EGLDisplay(d),
		C.EGLint(inv)))
}
func DestroyContext(d Display, c Context) bool {
	return goBool(C.eglDestroyContext(
		C.EGLDisplay(d),
		C.EGLContext(c)))
}
func GetCurrentSurface(readdraw int) Surface {
	return Surface(C.eglGetCurrentSurface(C.EGLint(readdraw)))
}
func QuerySurface(d Display, value []int, attribute int, surface Surface) bool {
	return goBool(C.eglQuerySurface(C.EGLDisplay(d),
		C.EGLSurface(surface),
		C.EGLint(attribute),
		(*C.EGLint)(unsafe.Pointer(&value[0]))))
}
func GetConfigs(d Display, configs []Config, configSize int, numConfig []int) bool {
	return goBool(C.eglGetConfigs(C.EGLDisplay(d),
		(*C.EGLConfig)(unsafe.Pointer(&configs[0])),
		C.EGLint(configSize),
		(*C.EGLint)(unsafe.Pointer(&numConfig[0]))))
}

func GetConfigAttrib(d Display, config Config, attribute int, value []int) bool {
	return goBool(C.eglGetConfigAttrib(C.EGLDisplay(d),
		C.EGLConfig(config),
		C.EGLint(attribute),
		(*C.EGLint)(unsafe.Pointer(&value[0]))))
}
func ChooseConfig(d Display, atrribList []int, configs []Config,
	configSize int, numConfig []int) bool {

	return goBool(C.eglChooseConfig(C.EGLDisplay(d),
		(*C.EGLint)(unsafe.Pointer(&atrribList[0])),
		(*C.EGLConfig)(&configs[0]),
		C.EGLint(configSize),
		(*C.EGLint)(unsafe.Pointer(&numConfig[0]))))
}
func CreateContext(d Display, config Config,
	shareContext Context, attribList []int) Context {

	return Context(C.eglCreateContext(C.EGLDisplay(d),
		C.EGLConfig(config),
		C.EGLContext(shareContext),
		(*C.EGLint)(unsafe.Pointer(&attribList[0]))))
}

func CreateWindowSurface(d Display, config Config,
	win NativeWindow, attribList []int) Surface {
	return Surface(C.eglCreateWindowSurface(C.EGLDisplay(d),
		C.EGLConfig(config),
		C.EGLNativeWindowType(uintptr(win)),
		(*C.EGLint)(unsafe.Pointer(&attribList[0]))))
}
func CreatePbufferSurface(d Display, config Config, attribList []int) Surface {

	return Surface(C.eglCreatePbufferSurface(C.EGLDisplay(d),
		C.EGLConfig(config),
		(*C.EGLint)(unsafe.Pointer(&attribList[0]))))
}
func CreatePixmapSurface(d Display, config Config,
	pixmap NativePixmap, attribList []int) Surface {

	return Surface(C.eglCreatePixmapSurface(C.EGLDisplay(d),
		C.EGLConfig(config),
		C.EGLNativePixmapType(uintptr(pixmap)),
		(*C.EGLint)(unsafe.Pointer(&attribList[0]))))
}

func CreatePbufferFromClientBuffer(d Display, buftype uint, config Config,
	buffer ClientBuffer, attribList []int) Surface {

	return Surface(C.eglCreatePbufferFromClientBuffer(C.EGLDisplay(d),
		C.EGLenum(buftype),
		C.EGLClientBuffer(buffer),
		C.EGLConfig(config),
		(*C.EGLint)(unsafe.Pointer(&attribList[0]))))
}
func SurfaceAttrib(d Display, surface Surface,
	attribute int, value int) bool {

	return goBool(C.eglSurfaceAttrib(C.EGLDisplay(d),
		C.EGLSurface(surface),
		C.EGLint(attribute),
		C.EGLint(value)))
}
func BindTexImage(d Display, surface Surface, buffer int) bool {
	return goBool(C.eglBindTexImage(C.EGLDisplay(d),
		C.EGLSurface(surface),
		C.EGLint(buffer)))
}
func ReleaseTexImage(d Display, surface Surface, buffer int) bool {
	return goBool(C.eglReleaseTexImage(C.EGLDisplay(d),
		C.EGLSurface(surface),
		C.EGLint(buffer)))
}
func MakeCurrent(d Display, draw Surface, read Surface, ctx Context) bool {
	return goBool(C.eglMakeCurrent(C.EGLDisplay(d),
		C.EGLSurface(draw),
		C.EGLSurface(read),
		C.EGLContext(ctx)))
}
func QueryContext(d Display, ctx Context, attribute int, value []int) bool {
	return goBool(C.eglQueryContext(C.EGLDisplay(d),
		C.EGLContext(ctx),
		C.EGLint(attribute),
		(*C.EGLint)(unsafe.Pointer(&value[0]))))
}
func CopyBuffers(d Display, surface Surface, target NativePixmap) bool {
	return goBool(C.eglCopyBuffers(C.EGLDisplay(d),
		C.EGLSurface(surface),
		C.EGLNativePixmapType(uintptr(target))))
}
func SwapBuffers(d Display, surface Surface) bool {
	return goBool(C.eglSwapBuffers(C.EGLDisplay(d),
		C.EGLSurface(surface)))
}
func BindAPI(api uint) bool {
	return goBool(C.eglBindAPI(C.EGLenum(api)))
}
func WaitNative(engine int) bool {
	return goBool(C.eglWaitNative(C.EGLint(engine)))
}
func QueryAPI() uint {
	return uint(C.eglQueryAPI())
}
func WaitClient() bool {
	return goBool(C.eglWaitClient())
}
func WaitGL() bool {
	return goBool(C.eglWaitGL())
}
func ReleaseThread() bool {
	return goBool(C.eglReleaseThread())
}
func GetCurrentDisplay() Display {
	return Display(C.eglGetCurrentDisplay())
}
func GetError() error {
	return Error(C.eglGetError())
}
