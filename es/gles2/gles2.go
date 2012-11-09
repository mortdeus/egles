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
	Enum     uint
	Boolean  bool
	Bitfield uint
	Byte     int8
	Short    int16
	Int      int64
	Sizei    int64
	UByte    byte
	UShort   uint16
	Uint     uint
	Float    float32
	Clampf   float32
	Fixed    int
	IntPtr   int
	SizeiPtr int

)

func GLStringCast(s string)) *C.GLchar {
	return (*C.GLchar)(C.CString(s))
}
func ActiveTexture(texture Enum) {
	C.glActiveTexture(C.GLenum(texture))
}
func AttachShader(program Uint, shader Uint) {
	C.glAttachShader(C.GLuint(program), C.GLuint(shader))
}
func BindAttribLocation(program Uint, index Uint, name Char) {
	C.glBindAttribLocation(C.GLuint(program), C.GLuint(index), GLStringCast(name))
}
func BindBuffer(target Enum, buffer Uint) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(buffer))
}
func BindFramebuffer(target Enum, framebuffer Uint) {
	C.glBindFramebuffer(C.GLenum(target), C.GLuint(framebuffer))
}
func BindRenderbuffer(target Enum, renderbuffer Uint) {
	C.glBindRenderbuffer(C.GLenum(target), C.GLuint(renderbuffer))

}
func BindTexture(target Enum, texture Uint) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))

}
func BlendColor(red Clampf, green Clampf, blue Clampf, alpha Clampf) {
	C.glBlendColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))

}
func BlendEquation(mode Enum) {
	C.glBlendEquation(C.GLenum(mode))

}
func BlendEquationSeparate(modeRGB Enum, modeAlpha Enum) {
	C.glBlendEquationSeparate(C.GLenum(modeRGB), C.GLenum(modeAlpha))

}
func BlendFunc(sfactor Enum, dfactor Enum) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))

}
func BlendFuncSeparate(srcRGB Enum, dstRGB Enum, srcAlpha Enum, dstAlpha Enum) {
	C.glBlendFuncSeparate(C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))

}
func BufferData(target Enum, size SizeiPtr, data Void, usage Enum) {
	C.glBufferData(C.GLenum(target), C.GLsizeiptr(size), (*C.GLvoid)(data), C.GLenum(usage))

}
func BufferSubData(target Enum, offset IntPtr, size SizeiPtr, data Void) {
	C.glBufferSubData(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size), (*C.GLvoid)(data))

}
func Clear(mask Bitfield) {
	C.glClear(C.GLbitfield(mask))

}
func ClearColor(red Clampf, green Clampf, blue Clampf, alpha Clampf) {
	C.glClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))

}
func ClearDepthf(depth Clampf) {
	C.glClearDepthf(C.GLclampf(depth))

}
func ClearStencil(s Int) {
	C.glClearStencil(C.GLint(s))

}
func ColorMask(red Boolean, green Boolean, blue Boolean, alpha Boolean) {
	C.glColorMask(C.GLboolean(red), C.GLboolean(green), C.GLboolean(blue), C.GLboolean(alpha))

}
func CompileShader(shader Uint) {
	C.glCompileShader(C.GLuint(shader))

}
func CompressedTexImage2D(target Enum, level Int, internalformat Enum, width Sizei, height Sizei, border Int, imageSize Sizei, data Void) {
	C.glCompressedTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(imageSize), (*C.GLvoid)(data))

}
func CompressedTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, imageSize Sizei, data Void) {
	C.glCompressedTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(imageSize), (*C.GLvoid)(data))

}
func CopyTexImage2D(target Enum, level Int, internalformat Enum, x Int, y Int, width Sizei, height Sizei, border Int) {
	C.glCopyTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))

}
func CopyTexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, x Int, y Int, width Sizei, height Sizei) {
	C.glCopyTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))

}
func CullFace(mode Enum) {
	C.glCullFace(C.GLenum(mode))

}
func DeleteBuffers(n Sizei, buffers *Uint) {
	C.glDeleteBuffers(C.GLsizei(n), (*C.GLuint)(buffers))

}
func DeleteFramebuffers(n Sizei, framebuffers *Uint) {
	C.glDeleteFramebuffers(C.GLsizei(n), (*C.GLuint)(framebuffers))

}
func DeleteProgram(program Uint) {
	C.glDeleteProgram(C.GLuint(program))

}
func DeleteRenderbuffers(n Sizei, renderbuffers *Uint) {
	C.glDeleteRenderbuffers(C.GLsizei(n), (*C.GLuint)(renderbuffers))

}
func DeleteShader(shader Uint) {
	C.glDeleteShader(C.GLuint(shader))

}
func DeleteTextures(n Sizei, textures *Uint) {
	C.glDeleteTextures(C.GLsizei(n), (*C.GLuint)(textures))

}
func DepthFunc(_func Enum) {
	C.glDepthFunc(C.GLenum(_func))

}
func DepthMask(flag Boolean) {
	C.glDepthMask(C.GLboolean(flag))

}
func DepthRangef(zNear Clampf, zFar Clampf) {
	C.glDepthRangef(C.GLclampf(zNear), C.GLclampf(zFar))

}
func DetachShader(program Uint, shader Uint) {
	C.glDetachShader(C.GLuint(program), C.GLuint(shader))

}
func Disable(cap Enum) {
	C.glDisable(C.GLenum(cap))

}
func DisableVertexAttribArray(index Uint) {
	C.glDisableVertexAttribArray(C.GLuint(index))

}
func DrawArrays(mode Enum, first Int, count Sizei) {
	C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))

}
func DrawElements(mode Enum, count Sizei, _type Enum, indices Void) {
	C.glDrawElements(C.GLenum(mode), C.GLsizei(count), C.GLenum(_type), (*C.GLvoid)(indices))
}
func Enable(cap Enum) {
	C.glEnable(C.GLenum(cap))

}
func EnableVertexAttribArray(index Uint) {
	C.glEnableVertexAttribArray(C.GLuint(index))

}
func Finish() {
	C.glEnableVertexAttribArray(C.GLuint(index))

}
func Flush() {
	C.glEnableVertexAttribArray(C.GLuint(index))

}
func FramebufferRenderbuffer(target Enum, attachment Enum, renderbuffertarget Enum, renderbuffer Uint) {
	C.glFramebufferRenderbuffer(C.GLenum(target), C.GLenum(attachment), C.GLenum(renderbuffertarget), C.GLuint(renderbuffe))

}
func FramebufferTexture2D(target Enum, attachment Enum, textarget Enum, texture Uint, level Int) {
	C.glFramebufferRenderbuffer(C.GLenum(target), C.GLenum(attachment), C.GLenum(renderbuffertarget), C.GLuint(renderbuffe))

}
func FrontFace(mode Enum) {
	C.glFrontFace(C.GLenum(mode))

}
func GenBuffers(n Sizei, buffers *Uint) {
	C.glGenBuffers(C.GLsizei(n), (*C.GLuint)(buffers))

}
func GenerateMipmap(target Enum) {

}
func GenFramebuffers(n Sizei, framebuffers *Uint) {
	C.glGenFramebuffers(C.GLsizei(n), (*C.GLuint)(framebuffers))

}
func GenRenderbuffers(n Sizei, renderbuffers *Uint) {

	C.glGenRenderbuffers(C.GLsizei(n), (*C.GLuint)(renderbuffers))
}
func GenTextures(n Sizei, textures *Uint) {

	C.glGenTextures(C.GLsizei(n), (*C.GLuint)(textures))
}
func GetActiveAttrib(program Uint, index Uint, bufsize Sizei, length *Sizei, size Int, _type Enum, name *Char) {

	C.glGetActiveAttrib(C.GLuint(program), C.GLuint(index), C.GLsizei(bufsize), (*C.GLsizei)(length), (*C.GLint)(size), C.GLenum(_type), (*C.GLchar)(name))
}
func GetActiveUniform(program Uint, index Uint, bufsize Sizei, length *Sizei, size *Int, _type Enum, name *Char) {

	C.glGetActiveUniform(C.GLuint(program), C.GLuint(index), C.GLsizei(bufsize), (*C.GLsizei)(length), (*C.GLint)(size), C.GLenum(_type), (*C.GLchar)(name))
}
func GetAttachedShaders(program Uint, maxcount Sizei, count *Sizei, shaders *Uint) {

	C.glGetAttachedShaders(C.GLuint(program), C.GLsizei(maxcount), (*C.GLsizei)(count), (*C.GLuint)(shaders))
}
func GetBooleanv(pname Enum, params *Boolean) {

	C.glGetBooleanv(C.GLenum(pname), C.GLboolean*(params))
}
func GetBufferParameteriv(target Enum, pname Enum, params *Int) {

	C.glGetBufferParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(params))
}
func GetFloatv(pname Enum, params *Float) {

	C.glGetFloatv(C.GLenum(pname), (*C.GLfloat)(params))
}
func GetFramebufferAttachmentParameteriv(target Enum, attachment Enum, pname Enum, params *Int) {

	C.glGetFramebufferAttachmentParameteriv(C.GLenum(target), C.GLenum(attachment), C.GLenum(pname), (*C.GLint)(params))
}
func GetIntegerv(pname Enum, params *Int) {

	C.glGetIntegerv(C.GLenum(pname), (*C.GLint)(params))
}
func GetProgramiv(program Uint, pname Enum, params *Int) {

	C.glGetProgramiv(C.GLuint(program), C.GLenum(pname), (*C.GLint)(params))

}
func GetProgramInfoLog(program Uint, bufsize Sizei, length *Sizei, infolog *Char) {

	C.glGetProgramInfoLog(C.GLuint(program), C.GLsizei(bufsize), (*C.GLsizei)(length), (*C.GLchar)(infolog))

}
func GetRenderbufferParameteriv(target Enum, pname Enum, params *Int) {

	C.glGetRenderbufferParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(params))
}
func GetShaderiv(shader Uint, pname Enum, params *Int) {

	C.glGetShaderiv(C.GLuint(shader), C.GLenum(pname), (*C.GLint)(params))
}
func GetShaderInfoLog(shader Uint, bufsize Sizei, length *Sizei, infolog *Char) {

	C.glGetShaderInfoLog(C.GLuint(shader), C.GLsizei(bufsize), (*C.GLsizei)(length), (*C.GLchar)(infolog))
}
func GetShaderPrecisionFormat(shadertype Enum, precisiontype Enum, _range *Int, precision *Int) {

	C.glGetShaderPrecisionFormat(C.GLenum(shadertype), C.GLenum(precisiontype), (*C.GLint)(_range), (*C.GLint)(precision))
}
func GetShaderSource(shader Uint, bufsize Sizei, length *Sizei, source *Char) {

	C.glGetShaderSource(C.GLuint(shader), C.GLsizei(bufsize), (*C.GLsizei)(length), (*C.GLchar)(source))
}
func GetTexParameterfv(target Enum, pname Enum, params *Float) {

	C.glGetTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(params))
}
func GetTexParameteriv(target Enum, pname Enum, params *Int) {

	C.glGetTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(params))
}
func GetUniformfv(program Uint, location Int, params *Float) {

	C.glGetUniformfv(C.GLuint(program), C.GLint(location), (*C.GLfloat)(params))
}
func GetUniformiv(program Uint, location Int, params *Int) {

	C.glGetUniformiv(C.GLuint(program), C.GLint(location), (*C.GLint)(params))
}
func GetVertexAttribfv(index Uint, pname Enum, params *Float) {

	C.glGetVertexAttribfv(C.GLuint(index), C.GLenum(pname), (*C.GLfloat)(params))
}
func GetVertexAttribiv(index Uint, pname Enum, params *Int) {

	C.glGetVertexAttribiv(C.GLuint(index), C.GLenum(pname), (*C.GLint)(params))
}
func GetVertexAttribPointerv(index Uint, pname Enum, pointer Void) {

	C.glGetVertexAttribPointerv(C.GLuint(index), C.GLenum(pname), (*C.GLvoid)(pointer))
}
func Hint(target Enum, mode Enum) {

	C.glHint(C.GLenum(target), C.GLenum(mode))
}
func LineWidth(width Float) {

	C.glLineWidth(C.GLfloat(width))
}
func LinkProgram(program Uint) {

	C.glLinkProgram(C.GLuint(program))
}
func PixelStorei(pname Enum, param Int) {

	C.glPixelStorei(C.GLenum(pname), C.GLint(param))
}
func PolygonOffset(factor Float, units Float) {

	C.glPolygonOffset(C.GLfloat(factor), C.GLfloat(units))
}
func ReadPixels(x Int, y Int, width Sizei, height Sizei, format Enum, _type Enum, pixels Void) {

	C.glReadPixels(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(_type), (*C.GLvoid)(pixels))
}
func ReleaseShaderCompiler() {

	C.glReleaseShaderCompiler()
}
func RenderbufferStorage(target Enum, internalformat Enum, width Sizei, height Sizei) {

	C.glRenderbufferStorage(C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height))
}
func SampleCoverage(value Clampf, invert Boolean) {

	C.glSampleCoverage(C.GLclampf(value), C.GLboolean(invert))
}
func Scissor(x Int, y Int, width Sizei, height Sizei) {

	C.glScissor(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(heigh))
}
func ShaderBinary(n Sizei, shaders *Uint, binaryformat Enum, binary Void, length Sizei) {

	C.glShaderBinary(C.GLsizei(n), (*C.GLuint)(shaders), C.GLenum(binaryformat), (*C.GLvoid)(binary), C.GLsizei(length))
}
func ShaderSource(shader Uint, count Sizei, _string *Char, length *Int) {

	C.glShaderSource(C.GLuint(shader), C.GLsizei(count), (*C.GLchar)(*string), (*C.GLint)(length))
}
func StencilFunc(_func Enum, ref Int, mask Uint) {

	C.glStencilFunc(C.GLenum(_func), C.GLint(ref), C.GLuint(mask))
}
func StencilFuncSeparate(face Enum, _func Enum, ref Int, mask Uint) {

	C.glStencilFuncSeparate(C.GLenum(face), C.GLenum(_func), C.GLint(ref), C.GLuint(mask))
}
func StencilMask(mask Uint) {

	C.glStencilMask(C.GLuint(mask))
}
func StencilMaskSeparate(face Enum, mask Uint) {

	C.glStencilMaskSeparate(C.GLenum(face), C.GLuint(mask))
}
func StencilOp(fail Enum, zfail Enum, zpass Enum) {

	C.glStencilOp(C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}
func StencilOpSeparate(face Enum, fail Enum, zfail Enum, zpass Enum) {

	C.glStencilOpSeparate(C.GLenum(face), C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}
func TexImage2D(target Enum, level Int, internalformat Int, width Sizei, height Sizei, border Int, format Enum, _type Enum, pixels Void) {

	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(_type), (*C.GLvoid)(pixels))
}
func TexParameterf(target Enum, pname Enum, param Float) {

	C.glTexParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}
func TexParameterfv(target Enum, pname Enum, params *Float) {

	C.glTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(params))
}
func TexParameteri(target Enum, pname Enum, param Int) {

	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}
func TexParameteriv(target Enum, pname Enum, params *Int) {

	C.glTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(params))
}
func TexSubImage2D(target Enum, level Int, xoffset Int, yoffset Int, width Sizei, height Sizei, format Enum, _type Enum, pixels Void) {

	C.glTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(_type), (*C.GLvoid)(pixels))
}
func Uniform1f(location Int, x Float) {

	C.glUniform1f(C.GLint(location), C.GLfloat(x))
}
func Uniform1fv(location Int, count Sizei, v *Float) {

	C.glUniform1fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(v))
}
func Uniform1i(location Int, x Int) {

	C.glUniform1i(C.GLint(location), C.GLint(x))
}
func Uniform1iv(location Int, count Sizei, v *Int) {

	C.glUniform1iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(v))
}
func Uniform2f(location Int, x Float, y Float) {

	C.glUniform2f(C.GLint(location), C.GLfloat(x), C.GLfloat(y))
}
func Uniform2fv(location Int, count Sizei, v *Float) {

	C.glUniform2fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(v))
}
func Uniform2i(location Int, x Int, y Int) {

	C.glUniform2i(C.GLint(location), C.GLint(x), C.GLint(y))
}
func Uniform2iv(location Int, count Sizei, v *Int) {

	C.glUniform2iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(v))
}
func Uniform3f(location Int, x Float, y Float, z Float) {

	C.glUniform3f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}
func Uniform3fv(location Int, count Sizei, v *Float) {

	C.glUniform3fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(v))
}
func Uniform3i(location Int, x Int, y Int, z Int) {

	C.glUniform3i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z))
}
func Uniform3iv(location Int, count Sizei, v *Int) {

	C.glUniform3iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(v))
}
func Uniform4f(location Int, x Float, y Float, z Float, w Float) {

	C.glUniform4f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}
func Uniform4fv(location Int, count Sizei, v *Float) {

	C.glUniform4fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(v))
}
func Uniform4i(location Int, x Int, y Int, z Int, w Int) {

	C.glUniform4i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}
func Uniform4iv(location Int, count Sizei, v *Int) {

	C.glUniform4iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(v))
}
func UniformMatrix2fv(location Int, count Sizei, transpose Boolean, value *Float) {

	C.glUniformMatrix2fv(C.GLint(location), C.GLsizei(count), C.GLboolean(transpose), (*C.GLfloat)(value))
}
func UniformMatrix3fv(location Int, count Sizei, transpose Boolean, value *Float) {

	C.glUniformMatrix3fv(C.GLint(location), C.GLsizei(count), C.GLboolean(transpose), (*C.GLfloat)(value))
}
func UniformMatrix4fv(location Int, count Sizei, transpose Boolean, value *Float) {

	C.glUniformMatrix4fv(C.GLint(location), C.GLsizei(count), C.GLboolean(transpose), (*C.GLfloat)(value))
}
func UseProgram(program Uint) {

	C.glUseProgram(C.GLuint(program))
}
func ValidateProgram(program Uint) {

	C.glValidateProgram(C.GLuint(program))
}
func VertexAttrib1f(indx Uint, x Float) {

	C.glVertexAttrib1f(C.GLuint(indx), C.GLfloat(x))
}
func VertexAttrib1fv(indx Uint, values *Float) {

	C.glVertexAttrib1fv(C.GLuint(indx), (*C.GLfloat)(values))
}
func VertexAttrib2f(indx Uint, x Float, y Float) {

	C.glVertexAttrib2f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y))
}
func VertexAttrib2fv(indx Uint, values *Float) {

	C.glVertexAttrib2fv(C.GLuint(indx), (*C.GLfloat)(values))
}
func VertexAttrib3f(indx Uint, x Float, y Float, z Float) {

	C.glVertexAttrib3f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}
func VertexAttrib3fv(indx Uint, values *Float) {

	C.glVertexAttrib3fv(C.GLuint(indx), (*C.GLfloat)(values))
}
func VertexAttrib4f(indx Uint, x Float, y Float, z Float, w Float) {

	C.glVertexAttrib4f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}
func VertexAttrib4fv(indx Uint, values *Float) {

	C.glVertexAttrib4fv(C.GLuint(indx), (*C.GLfloat)(values))
}
func VertexAttribPointer(indx Uint, size Int, _type Enum, normalized Boolean, stride Sizei, ptr Void) {

	C.glVertexAttribPointer(C.GLuint(indx), C.GLint(size), C.GLenum(_type), C.GLboolean(normalized), C.GLsizei(stride), (*C.GLvoid)(ptr))
}
func Viewport(x Int, y Int, width Sizei, height Sizei) {

	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
func IsBuffer(buffer Uint) Boolean {

	return C.glIsBuffer(C.GLuint(buffer))
}
func IsEnabled(cap Enum) Boolean {

	return C.glIsEnabled(C.GLenum(cap))
}
func IsFramebuffer(framebuffer Uint) Boolean {

	return C.glIsFramebuffer(C.GLuint(framebuffer))
}
func IsProgram(program Uint) Boolean {

	return C.glIsProgram(C.GLuint(program))
}
func IsRenderbuffer(renderbuffer Uint) Boolean {

	return C.glIsRenderbuffer(C.GLuint(renderbuffer))
}
func IsShader(shader Uint) Boolean {

	return C.glIsShader(C.GLuint(shader))
}
func IsTexture(texture Uint) Boolean {

	return C.glIsTexture(C.GLuint(texture))
}
func GetError() Enum {

	return C.glGetError()
}
func CheckFramebufferStatus(target Enum) Enum {

	return C.glCheckFramebufferStatus(C.GLenum(target))
}
func CreateProgram() Uint {

	return C.glCreateProgram()
}
func CreateShader(_type Enum) Uint {

	return C.glCreateShader(C.GLenum(_type))
}
func GetAttribLocation(program Uint, name *Char) int {

	return C.glGetAttribLocation(C.GLuint(program), (*C.GLchar)(name))
}
func GetUniformLocation(program Uint, name Char) int {

	return C.glGetUniformLocation(C.GLuint(program), (*C.GLchar)(name))
}
func GetString(name Enum) *UByte {

	return C.glGetString(C.GLenum(name))
}
