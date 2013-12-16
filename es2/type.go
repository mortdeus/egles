package es2

//#cgo linux LDFLAGS: -lGLESv2  -lEGL
//#include <stdlib.h>
//#include <GLES2/gl2.h>
//#include <GLES2/gl2ext.h>
//#include <GLES2/gl2platform.h>
import "C"

func CString(s string) *C.GLchar {
	return (*C.GLchar)(C.CString(s))
}
func GoString(s *C.GLchar) string {
	return C.GoString((*C.char)(s))
}
func goBoolean(n C.GLboolean) *bool {
	b := n == 1
	return &b
}
func glBoolean(n bool) C.GLboolean {
	var b int
	if n == true {
		b = 1
	}
	return C.GLboolean(b)
}
