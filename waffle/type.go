package waffle

/*
#cgo pkg-config: waffle-1
#include <stdint.h>
#include <stdbool.h>
#include "waffle.h"
*/
import "C"

type (
	Config  C.struct_waffle_config
	Display C.struct_waffle_display
	Context C.struct_waffle_context

	GbmConfig     C.struct_waffle_gbm_config
	GlxConfig     C.struct_glx_config
	XeglConfig    C.struct__x11_egl_config
	WaylandConfig C.struct_wayland_config
	NativeConfig  C.union_waffle_native_config

	GbmContext     C.struct_waffle_gbm_Context
	GlxContext     C.struct_glx_Context
	XeglContext    C.struct__x11_egl_Context
	WaylandContext C.struct_wayland_Context
	NativeContext  C.union_waffle_native_Context
)
