package gles2

//#cgo linux LDFLAGS: -lGLESv2  -lEGL  
//#include <stdlib.h>
//#include <GLES2/gl2.h>
//#include <GLES2/gl2ext.h>
//#include <GLES2/gl2platform.h>
import "C"
import "unsafe"

type (
	Void     unsafe.Pointer
	Char     byte
	Enum     uint32
	Boolean  bool
	Bitfield uint32
	Byte     int8
	Short    int16
	Int      int32
	Sizei    int32
	UByte    byte
	UShort   uint16
	Uint     uint32
	Float    float32
	Clampf   float32
	Fixed    uintptr
	IntPtr   uintptr
	SizeiPtr uintptr
)

func glString(s string) *C.GLchar {
	return (*C.GLchar)(C.CString(s))
}
func goString(s *C.GLchar) *string {
	gs := C.GoString((*C.char)(s))
	return &gs
}
func goBoolean(n C.GLboolean) *Boolean {
	b := Boolean(TRUE == n)
	return &b
}
func glBoolean(n Boolean) C.GLboolean {
	var b int
	if n == true {
		b = TRUE
	}
	return C.GLboolean(b)
}
