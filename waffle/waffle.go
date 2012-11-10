package waffle

/*
#cgo pkg-config: waffle-1
#include <stdint.h>
#include <stdbool.h>
#include "waffle.h"
*/
import "C"

func Init(attrbList *int32) {
	C.waffle_init((*C.int32_t)(attrbList))
}

//waffle_config
/*************************/
func ChooseConfig(d *Display, attrbList []int32) *Config {
	c := Config(*C.waffle_config_choose(
		(*C.struct_waffle_display)(&(*d)),
		(*C.int32_t)(&attrbList[0])))
	return &c
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

//waffle_context
/*******************************/
func CreateContext(config *Config, sharedCtx *Context) *Context {
	c := Context(*C.waffle_context_create(
		(*C.struct_waffle_config)(&(*config)),
		(*C.struct_waffle_context)(&(sharedCtx))))
	return &c
}
