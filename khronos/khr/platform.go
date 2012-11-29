package khr

/*
#cgo pkg-config: egl
#include <KHR/khrplatform.h>
*/
import "C"

import (
	"github.com/mortdeus/egles/misc"
)

var (
	Int64Support bool
	FloatSupport bool
)

func init() {
	glog.Init(0)

	if C.KHRONOS_SUPPORT_INT64 == 1 {
		Int64Support = true

	}
	if C.KHRONOS_SUPPORT_FLOAT == 1 {
		FloatSupport = true
	}
	glog.Printf(
		"Int64 Support: %v\nFloat Support: %v\n",
		Int64Support, FloatSupport)
}

func Int8(n int8) C.khronos_int8_t {
	return C.khronos_int8_t(n)
}
func Uint8(n uint8) C.khronos_uint8_t {
	return C.khronos_uint8_t(n)
}
func Int16(n int16) C.khronos_int16_t {
	return C.khronos_int16_t(n)
}
func Uint16(n uint16) C.khronos_uint16_t {
	return C.khronos_uint16_t(n)
}
func Int32(n int32) C.khronos_int32_t {
	return C.khronos_int32_t(n)
}
func Uint32(n uint32) C.khronos_uint32_t {
	return C.khronos_uint32_t(n)
}
func Int64(n int64) C.khronos_int64_t {
	if !Int64Support {
		glog.Panicln("Int64's are not supported on this architecture, aborting.")
	}
	return C.khronos_int64_t(n)
}
func Uint64(n uint64) C.khronos_uint64_t {
	if !Int64Support {
		glog.Panicln("Int64's are not supported on this architecture, aborting.")
	}
	return C.khronos_uint64_t(n)
}
func Intptr(n uintptr) C.khronos_intptr_t {
	return C.khronos_intptr_t(n)
}
func Uintptr(n uintptr) C.khronos_uintptr_t {
	return C.khronos_uintptr_t(n)
}
func Ssize(n uintptr) C.khronos_ssize_t {
	return C.khronos_ssize_t(n)
}
func Usize(n uintptr) C.khronos_usize_t {
	return C.khronos_usize_t(n)
}
func Float(n float32) C.khronos_float_t {
	if !FloatSupport {
		glog.Panicln("Floats are not supported on this architecture, aborting.")
	}
	return C.khronos_float_t(n)
}

func Utime(nano uint64) C.khronos_utime_nanoseconds_t {
	if !Int64Support {
		glog.Panicln("Int64's are not supported on this architecture, aborting.")
	}
	return C.khronos_utime_nanoseconds_t(nano)
}
func Stime(nano int64) C.khronos_stime_nanoseconds_t {
	if !Int64Support {
		glog.Panicln("Int64's are not supported on this architecture, aborting.")
	}
	return C.khronos_stime_nanoseconds_t(nano)
}
func Boolean(n bool) C.khronos_boolean_enum_t {
	if !n {
		return C.KHRONOS_FALSE
	}
	return C.KHRONOS_TRUE
}
