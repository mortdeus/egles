package waffle

/*
#cgo pkg-config: waffle-1
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#include <stddef.h>
#include "waffle.h"
*/
import "C"
import "unsafe"

func Init(attrbList *int32) {
	C.waffle_init((*C.int32_t)(attrbList))
}

//waffle_config.h

type (
	Config        C.struct_waffle_config
	GbmConfig     C.struct_waffle_gbm_config
	GlxConfig     C.struct_glx_config
	XeglConfig    C.struct__x11_egl_config
	WaylandConfig C.struct_wayland_config
	NativeConfig  C.union_waffle_native_config
)

func (c *Config) Choose(d *Display, attrbList []int32) {
	nc := Config(*C.waffle_config_choose(
		(*C.struct_waffle_display)(&(*d)),
		(*C.int32_t)(&attrbList[0])))
	c = &nc
}
func (c *Config) Destroy() bool {
	b := C.waffle_config_destroy(
		(*C.struct_waffle_config)(&(*c)))
	return bool(b)
}
func (c *Config) GetNative() *NativeConfig {
	nc := NativeConfig(*C.waffle_config_get_native(
		(*C.struct_waffle_config)(&(*c))))
	return &nc
}

//waffle_context.h

type (
	Context        C.struct_waffle_context
	GbmContext     C.struct_waffle_gbm_context
	GlxContext     C.struct_glx_context
	XeglContext    C.struct__x11_egl_context
	WaylandContext C.struct_wayland_context
	NativeContext  C.union_waffle_native_context
)

func (c *Context) Create(sharedCtx *Context) {
	nc := Context(*C.waffle_context_create(
		(*C.struct_waffle_config)(&(*c)),
		(*C.struct_waffle_context)(&(*sharedCtx))))
	c = &nc
}
func (c *Context) Destroy() bool {
	return bool(C.waffle_context_destroy(
		(*C.struct_waffle_context)(&(*c))))

}
func (c *Context) GetNative() *NativeContext {
	nc := NativeContext((*C.waffle_context_get_native(
		(*C.struct_waffle_context)(&(*c)))))
	return &nc
}

//waffle_display.h

type (
	Display        C.struct_waffle_display
	GbmDisplay     C.struct_waffle_gbm_display
	GlxDisplay     C.struct_glx_display
	XeglDisplay    C.struct__x11_egl_display
	WaylandDisplay C.struct_wayland_display
	NativeDisplay  C.union_waffle_native_display
)

func (d *Display) Connect(name string) {
	s := C.CString(name)
	defer C.free(unsafe.Pointer(s))
	nd := Display(*C.waffle_display_connect(
		C.CString(name)))
	d = &nd
}
func (d *Display) Disconnect() bool {
	return bool(C.waffle_display_disconnect(
		(*C.struct_waffle_display)(&(*d))))

}
func (d *Display) SupportContextAPI(contextAPI int32) bool {
	return bool(C.waffle_display_supports_context_api(
		(*C.struct_waffle_display)(&(*d)),
		C.int32_t(contextAPI)))
}
func (d *Display) GetNative() *NativeDisplay {
	nd := NativeDisplay((*C.waffle_display_get_native(
		(*C.struct_waffle_display)(&(*d)))))
	return &nd
}

//Waffle_dl.h
func CanOpenDL(dl int32) bool {
	return bool(C.waffle_dl_can_open(C.int32_t(dl)))
}
func DLSym(dl int32, name string) unsafe.Pointer {
	s := C.CString(name)
	defer C.free(unsafe.Pointer(s))
	return C.waffle_dl_sym(C.int32_t(dl), s)
}

//waffle_error.h
type Error C.struct_waffle_error_info

const (
	NO_ERROR = 0x00 << iota
	ERROR_FATAL
	ERROR_UNKNOWN
	ERROR_INTERNAL
	ERROR_BAD_ALLOC
	ERROR_NOT_INITIALIZED
	ERROR_ALREADY_INITIALIZED
	_
	ERROR_BAD_ATTRIBUTE
	_
	ERROR_BAD_PARAMETER
	ERROR_BAD_DISPLAY_MATCH
	ERROR_UNSUPPORTED_ON_PLATFORM
	ERROR_BUILT_WITHOUT_SUPPORT
)

func GetErrorCode() int32 {
	return int32(C.waffle_error_get_code())
}
func (e *Error) GetInfo() {
	ne := Error(*C.waffle_error_get_info())
	e = &ne
}
func ErrorToString(e int32) string {
	return C.GoString(C.waffle_error_to_string(C.int32_t(e)))
}
