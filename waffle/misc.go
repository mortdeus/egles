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

func IsExtensionInString(ext string, extName string) bool {
	s1 := C.CString(ext)
	defer C.free(unsafe.Pointer(s1))

	s2 := C.CString(extName)
	defer C.free(unsafe.Pointer(s2))
	return bool(C.waffle_is_extension_in_string(s1, s2))
}
func MakeCurrent(d *Display, w *Window, c *Context) bool {
	return bool(C.waffle_make_current(
		(*C.struct_waffle_display)(&(*d)),
		(*C.struct_waffle_window)(&(*w)),
		(*C.struct_waffle_context)(&(*c))))
}
