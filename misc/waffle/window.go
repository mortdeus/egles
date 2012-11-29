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

//import "unsafe"

//waffle_window.h

type (
	Window        C.struct_waffle_Window
	GbmWindow     C.struct_waffle_gbm_Window
	GlxWindow     C.struct_waffle_glx_Window
	XeglWindow    C.struct_waffle_x11_egl_Window
	WaylandWindow C.struct_waffle_wayland_Window
	NativeWindow  C.union_waffle_native_Window
)

func (w *Window) Create(c *Config, height, width int32) {
	nw := Window(*C.waffle_window_create(
		(*C.struct_waffle_window)(&(*c)),
		C.int32_t(height),
		C.int32_t(width)))
	w = &nw
}
func (w *Window) Destroy() bool {
	return bool(C.waffle_window_destroy(
		(*C.struct_waffle_window)(&(*w))))
}
func (w *Window) GetNative() *NativeConfig {
	nw := NativeConfig(*C.waffle_window_get_native(
		(*C.struct_waffle_window)(&(*w))))
	return &nw
}
func (w *Window) Show() bool {
	return bool(C.waffle_window_show((*C.struct_waffle_window)(&(*w))))
}

func (w *Window) SwapBuffers() bool {
	return bool(C.waffle_window_swap_buffers((*C.struct_waffle_window)(&(*w))))
}

//waffle_wayland.h

type (
	WLCompositor   C.struct_wl_compositor
	WLDisplay      C.struct_wl_display
	WLEglWindow    C.struct_wl_egl_window
	WLShell        C.struct_wl_shell
	WLShellSurface C.struct_wl_shell_surface
	WLSurface      C.struct_wl_surface
)

//waffle_gbm.h

type (
	GbmDevice  C.struct_gbm_surface
	GbmSurface C.struct_gbm_device
)
