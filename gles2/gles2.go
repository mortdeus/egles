package gles2

//#cgo linux LDFLAGS: -lGLESv2  -lEGL  
//#include <stdlib.h>
//#include <GLES2/gl2.h>
//#include <GLES2/gl2ext.h>
//#include <GLES2/glplatform.h>

import "C"
import "unsafe"

type (
	Void     C.GLvoid
	Char     C.GLchar
	Enum     C.GLenum
	Boolean  C.GLboolean
	Bitfield C.GLbitfield
	Byte     C.GLbyte
	Short    C.GLshort
	Int      C.GLint
	Sizei    C.GLsizei
	UByte    C.GLubyte
	UShort   C.GLushort
	UInt     C.GLuint
	Float    C.GLfloat
	Clampf   C.GLclampf
	Fixed    C.GLfixed
	IntPtr   C.GLintptr
	SizeiPtr C.GLsizeiptr
)

func GLStringCast(s string) *Char {
	str := C.CString(s)
	defer C.free(unsafe.Pointer(str))
	return (*Char)(&str[0])
}

func ActiveTexture(texture Enum) {
	C.glActiveTexture(texture)
}

func AttachShader(program, shader UInt) {
	C.glAttachShader(program, shader)
}

func BindAttribLocation(program, index UInt, name string) {
	C.glBindAttribLocation(program, index, GLStringCast(s))
}

func BindBuffer(target Enum, buffer UInt) {
	C.glBindBuffer(target, buffer)
}

func BindFrameBuffer(target Enum, framebuffer UInt) {
	C.glBindFrameBuffer(target, framebuffer)
}

func BindRenderBuffer() {}

func BindTexture() {}

func BlendColor() {}

func BlendEquation() {}

func BlendEquationSeparate() {}

func BlendFunc() {}

func BlendFuncSeparate() {}

func BufferData() {}

func BufferSubData() {}

func CheckFrameBufferStatus() {}

func Clear() {}

func ClearColor() {}

func ClearDepthf() {}

func ClearStencil() {}

func ColorMask() {}

func CompileShader() {}

func CompressedTexImage2D() {}

func CompressedTexSubImage2D() {}

func CopyTexImage2D() {}

func CopyTexSubImage2D() {}

func CreateProgram() {}

func CreateShader() {}

func cullface() {}

func DeleteBuffers() {}

func DeleteFrameBuffers() {}

func DeleteProgram() {}

func DeleteRenderBuffers() {}

func DeleteShader() {}

func DeleteTextures() {}

func DepthFunc() {}

func DepthMask() {}

func DepthRangef() {}

func DetachShader() {}

func Disable() {}

func DisableVertexAttribArray() {}

func DrawArrays() {}

func DrawElements() {}

func Enable() {}

func EnableVertexAttribArray() {}

func Finish() {}

func Flush() {}

func FrameBufferRenderBuffer() {}

func FrameBufferTexture2D() {}

func FrontFace() {}

func GenBuffers() {}

func GenerateMipMap() {}

func GenFrameBuffers() {}

func GenRenderBuffers() {}

func GenTextures() {}

func GetActiveAttrib() {}

func GetActiveUniform() {}

func GetAttachedshaders() {}

func GetAttribLocation() {}

func GetBooleanv() {}

func GetBufferParameterIv() {}

func GetError() {}

func GetFloatv() {}

func GetFrameBufferAttachmentParameterIv() {}

func GetIntegerv() {}

func GetProgramIv() {}

func GetProgramInfoLog() {}

func GetRenderBufferParameterIv() {}

func GetShaderIv() {}

func GetShaderInfoLog() {}

func GetShaderPrecisionFormat() {}

func GetShaderSource() {}

func GetString() {}

func GetTexParameterFv() {}

func GetTexParameterIv() {}

func GetUniformFv() {}

func GetUniformIv() {}

func GetUniformLocation() {}

func GetVertexAttribFv() {}

func getvertexattribiv() {}

func getvertexattribpointerv() {}

func hint() {}

func IsBuffer() {}

func IsEnabled() {}

func IsFrameBuffer() {}

func IsProgram() {}

func IsRenderBuffer() {}

func IsShader() {}

func IsTexture() {}

func LineWidth() {}

func LinkProgram() {}

func PixelStorei() {}

func PolygonOffset() {}

func ReadPixels() {}

func ReleaseShaderCompiler() {}

func RenderBufferStorage() {}

func SampleCoverage() {}

func Scissor() {}

func ShaderBinary() {}

func ShaderSource() {}

func StencilFunc() {}

func StencilFuncSeparate() {}

func StencilMask() {}

func StencilMaskSeparate() {}

func stencilop() {}

func StencilOpSeparate() {}

func TexImage2D() {}

func TexParameterf() {}

func TexParameterFv() {}

func TexParameteri() {}

func TexParameterIv() {}

func TexSubImageTarget() {}

func Uniform1f() {}

func Uniform1fv() {}

func Uniform1i() {}

func Uniform1iv() {}

func Uniform2f() {}

func Uniform2fv() {}

func Uniform2i() {}

func Uniform2iv() {}

func Uniform3f() {}

func Uniform3fv() {}

func Uniform3i() {}

func Uniform3iv() {}

func Uniform4f() {}

func Uniform4fv() {}

func Uniform4i() {}

func Uniform4iv() {}

func UniformMatrix2fv() {}

func UniformMatrix3fv() {}

func UniformMatrix4fv() {}

func UseProgram() {}

func ValidateProgram() {}

func VertexAttrib1f() {}

func VertexAttrib1fv() {}

func VertexAttrib2f() {}

func VertexAttrib2fv() {}

func VertexAttrib3f() {}

func VertexAttrib3fv() {}

func VertexAttrib4f() {}

func VertexAttrib4fv() {}

func VertexAttribPointer() {}

func Viewport() {}
