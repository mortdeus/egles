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

//waffle_gl_misc.h

type ProcAddress unsafe.Pointer

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
func GetProcAddress(name string) ProcAddress {
	s := C.CString(name)
	defer C.free(unsafe.Pointer(s))
	return ProcAddress(C.waffle_get_proc_address(s))
}

//waffle_attrib_list.h

type Attribute C.int32_t
type AttributeList []Attribute

func (al *AttributeList) Length() int32 {
	return int32(C.waffle_attrib_list_length((*C.int32_t)(&(*al)[0])))
}

func (al *AttributeList) Get(key int32, value *int32) bool {
	return bool(C.waffle_attrib_list_get(
		(*C.int32_t)(&(*al)[0]),
		C.int32_t(key),
		(*C.int32_t)(value)))
}

func (al *AttributeList) GetWithDefault(key, defaultVal int32, value *int32) bool {
	return bool(C.waffle_attrib_list_get_with_default(
		(*C.int32_t)(&(*al)[0]),
		C.int32_t(key),
		(*C.int32_t)(value),
		C.int32_t(defaultVal)))
}
func (al *AttributeList) Update(key int32, value int32) bool {
	list := C.int32_t((*al)[0])
	return bool(C.waffle_attrib_list_update(
		(*C.int32_t)(&list),
		C.int32_t(key),
		C.int32_t(value)))
}
