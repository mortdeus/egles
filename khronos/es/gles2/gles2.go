package gles2

//#cgo linux LDFLAGS: -lGLESv2  -lEGL  
//#include <stdlib.h>
//#include <GLES2/gl2.h>
//#include <GLES2/gl2ext.h>
//#include <GLES2/gl2platform.h>
import "C"
import "unsafe"

func ActiveTexture(
	texture Enum) {
	C.glActiveTexture(
		C.GLenum(texture))
}
func AttachShader(
	program Uint, shader Uint) {
	C.glAttachShader(
		C.GLuint(program),
		C.GLuint(shader))
}
func BindAttribLocation(
	program Uint, index Uint, name *string) {
	s := glString(*name)
	C.glBindAttribLocation(
		C.GLuint(program),
		C.GLuint(index),
		s)
	name = goString(s)
}
func BindBuffer(
	target Enum, buffer Uint) {
	C.glBindBuffer(
		C.GLenum(target),
		C.GLuint(buffer))
}
func BindFramebuffer(
	target Enum, framebuffer Uint) {
	C.glBindFramebuffer(
		C.GLenum(target),
		C.GLuint(framebuffer))
}
func BindRenderbuffer(
	target Enum, renderbuffer Uint) {
	C.glBindRenderbuffer(
		C.GLenum(target),
		C.GLuint(renderbuffer))

}
func BindTexture(
	target Enum, texture Uint) {
	C.glBindTexture(
		C.GLenum(target),
		C.GLuint(texture))

}
func BlendColor(
	red Clampf, green Clampf,
	blue Clampf, alpha Clampf) {
	C.glBlendColor(
		C.GLclampf(red),
		C.GLclampf(green),
		C.GLclampf(blue),
		C.GLclampf(alpha))

}
func BlendEquation(
	mode Enum) {
	C.glBlendEquation(
		C.GLenum(mode))

}
func BlendEquationSeparate(
	modeRGB Enum, modeAlpha Enum) {
	C.glBlendEquationSeparate(
		C.GLenum(modeRGB),
		C.GLenum(modeAlpha))

}
func BlendFunc(
	sfactor Enum, dfactor Enum) {
	C.glBlendFunc(
		C.GLenum(sfactor),
		C.GLenum(dfactor))

}
func BlendFuncSeparate(
	srcRGB Enum, dstRGB Enum,
	srcAlpha Enum, dstAlpha Enum) {
	C.glBlendFuncSeparate(
		C.GLenum(srcRGB),
		C.GLenum(dstRGB),
		C.GLenum(srcAlpha),
		C.GLenum(dstAlpha))

}
func BufferData(
	target Enum, size SizeiPtr,
	data Void, usage Enum) {
	C.glBufferData(
		C.GLenum(target),
		C.GLsizeiptr(size),
		unsafe.Pointer(data),
		C.GLenum(usage))

}
func BufferSubData(
	target Enum, offset IntPtr,
	size SizeiPtr, data Void) {
	C.glBufferSubData(
		C.GLenum(target),
		C.GLintptr(offset),
		C.GLsizeiptr(size),
		unsafe.Pointer(data))

}
func Clear(
	mask Bitfield) {
	C.glClear(
		C.GLbitfield(mask))

}
func ClearColor(
	red Clampf, green Clampf,
	blue Clampf, alpha Clampf) {
	C.glClearColor(
		C.GLclampf(red),
		C.GLclampf(green),
		C.GLclampf(blue),
		C.GLclampf(alpha))

}
func ClearDepthf(
	depth Clampf) {
	C.glClearDepthf(
		C.GLclampf(depth))

}
func ClearStencil(
	s Int) {
	C.glClearStencil(
		C.GLint(s))

}
func ColorMask(
	red Boolean, green Boolean,
	blue Boolean, alpha Boolean) {
	C.glColorMask(
		glBoolean(red),
		glBoolean(green),
		glBoolean(blue),
		glBoolean(alpha))

}
func CompileShader(
	shader Uint) {
	C.glCompileShader(
		C.GLuint(shader))

}
func CompressedTexImage2D(
	target Enum, level Int, internalformat Enum,
	width Sizei, height Sizei, border Int,
	imageSize Sizei, data Void) {
	C.glCompressedTexImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLenum(internalformat),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLint(border),
		C.GLsizei(imageSize),
		unsafe.Pointer(data))

}
func CompressedTexSubImage2D(
	target Enum, level Int,
	xoffset Int, yoffset Int, width Sizei, height Sizei,
	format Enum, imageSize Sizei, data Void) {
	C.glCompressedTexSubImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(xoffset),
		C.GLint(yoffset),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLsizei(imageSize),
		unsafe.Pointer(data))

}
func CopyTexImage2D(
	target Enum, level Int, internalformat Enum,
	x Int, y Int, width Sizei, height Sizei, border Int) {
	C.glCopyTexImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLenum(internalformat),
		C.GLint(x),
		C.GLint(y),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLint(border))

}
func CopyTexSubImage2D(
	target Enum, level Int, xoffset Int,
	yoffset Int, x Int, y Int, width Sizei, height Sizei) {
	C.glCopyTexSubImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(xoffset),
		C.GLint(yoffset),
		C.GLint(x),
		C.GLint(y),
		C.GLsizei(width),
		C.GLsizei(height))

}
func CullFace(
	mode Enum) {
	C.glCullFace(
		C.GLenum(mode))

}
func DeleteBuffers(
	n Sizei, buffers *Uint) {
	C.glDeleteBuffers(
		C.GLsizei(n),
		(*C.GLuint)(buffers))

}
func DeleteFramebuffers(
	n Sizei, framebuffers *Uint) {
	C.glDeleteFramebuffers(
		C.GLsizei(n),
		(*C.GLuint)(framebuffers))

}
func DeleteProgram(
	program Uint) {
	C.glDeleteProgram(
		C.GLuint(program))

}
func DeleteRenderbuffers(
	n Sizei, renderbuffers *Uint) {
	C.glDeleteRenderbuffers(
		C.GLsizei(n),
		(*C.GLuint)(renderbuffers))

}
func DeleteShader(
	shader Uint) {
	C.glDeleteShader(
		C.GLuint(shader))

}
func DeleteTextures(
	n Sizei, textures *Uint) {
	C.glDeleteTextures(
		C.GLsizei(n),
		(*C.GLuint)(textures))

}
func DepthFunc(
	func_ Enum) {
	C.glDepthFunc(
		C.GLenum(func_))

}
func DepthMask(
	flag Boolean) {
	C.glDepthMask(
		glBoolean(flag))

}
func DepthRangef(
	zNear Clampf, zFar Clampf) {
	C.glDepthRangef(
		C.GLclampf(zNear),
		C.GLclampf(zFar))

}
func DetachShader(
	program Uint, shader Uint) {
	C.glDetachShader(
		C.GLuint(program),
		C.GLuint(shader))

}
func Disable(
	cap Enum) {
	C.glDisable(
		C.GLenum(cap))

}
func DisableVertexAttribArray(
	index Uint) {
	C.glDisableVertexAttribArray(
		C.GLuint(index))

}
func DrawArrays(
	mode Enum, first Int, count Sizei) {
	C.glDrawArrays(
		C.GLenum(mode),
		C.GLint(first),
		C.GLsizei(count))

}
func DrawElements(
	mode Enum, count Sizei,
	type_ Enum, indices Void) {
	C.glDrawElements(
		C.GLenum(mode),
		C.GLsizei(count),
		C.GLenum(type_),
		unsafe.Pointer(indices))
}
func Enable(
	cap Enum) {
	C.glEnable(C.GLenum(cap))

}
func EnableVertexAttribArray(
	index Uint) {
	C.glEnableVertexAttribArray(
		C.GLuint(index))
}
func Finish() {
	C.glFinish()
}
func Flush() {
	C.glFlush()
}
func FramebufferRenderbuffer(
	target Enum, attachment Enum,
	renderbuffertarget Enum, renderbuffer Uint) {
	C.glFramebufferRenderbuffer(
		C.GLenum(target),
		C.GLenum(attachment),
		C.GLenum(renderbuffertarget),
		C.GLuint(renderbuffer))
}
func FramebufferTexture2D(
	target Enum, attachment Enum,
	textarget Enum, texture Uint, level Int) {
	C.glFramebufferTexture2D(
		C.GLenum(target),
		C.GLenum(attachment),
		C.GLenum(textarget),
		C.GLuint(texture),
		C.GLint(level))
}
func FrontFace(
	mode Enum) {
	C.glFrontFace(
		C.GLenum(mode))
}
func GenBuffers(
	n Sizei, buffers *Uint) {
	C.glGenBuffers(
		C.GLsizei(n),
		(*C.GLuint)(buffers))
}
func GenerateMipmap(
	target Enum) {

}
func GenFramebuffers(
	n Sizei, framebuffers *Uint) {
	C.glGenFramebuffers(
		C.GLsizei(n),
		(*C.GLuint)(framebuffers))
}
func GenRenderbuffers(
	n Sizei, renderbuffers *Uint) {
	C.glGenRenderbuffers(
		C.GLsizei(n),
		(*C.GLuint)(renderbuffers))
}
func GenTextures(
	n Sizei, textures *Uint) {
	C.glGenTextures(
		C.GLsizei(n),
		(*C.GLuint)(textures))
}
func GetActiveAttrib(
	program Uint, index Uint, bufsize Sizei,
	length *Sizei, size *Int, type_ *Enum, name *string) {
	s := glString(*name)
	C.glGetActiveAttrib(
		C.GLuint(program),
		C.GLuint(index),
		C.GLsizei(bufsize),
		(*C.GLsizei)(length),
		(*C.GLint)(size),
		(*C.GLenum)(type_),
		s)
	name = goString(s)
}
func GetActiveUniform(
	program Uint, index Uint, bufsize Sizei,
	length *Sizei, size *Int, type_ *Enum, name *string) {
	s := glString(*name)
	C.glGetActiveUniform(
		C.GLuint(program),
		C.GLuint(index),
		C.GLsizei(bufsize),
		(*C.GLsizei)(length),
		(*C.GLint)(size),
		(*C.GLenum)(type_),
		s)
	name = goString(s)
}
func GetAttachedShaders(
	program Uint, maxcount Sizei, count *Sizei, shaders *Uint) {
	C.glGetAttachedShaders(
		C.GLuint(program),
		C.GLsizei(maxcount),
		(*C.GLsizei)(count),
		(*C.GLuint)(shaders))
}
func GetBooleanv(
	pname Enum, params *Boolean) {
	p := glBoolean(*params)
	C.glGetBooleanv(
		C.GLenum(pname),
		&p)
	params = goBoolean(p)

}
func GetBufferParameteriv(
	target Enum, pname Enum, params *Int) {
	C.glGetBufferParameteriv(
		C.GLenum(target),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetFloatv(
	pname Enum, params *Float) {
	C.glGetFloatv(
		C.GLenum(pname),
		(*C.GLfloat)(params))
}
func GetFramebufferAttachmentParameteriv(
	target Enum, attachment Enum, pname Enum, params *Int) {
	C.glGetFramebufferAttachmentParameteriv(
		C.GLenum(target),
		C.GLenum(attachment),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetIntegerv(
	pname Enum, params *Int) {
	C.glGetIntegerv(
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetProgramiv(
	program Uint, pname Enum, params *Int) {
	C.glGetProgramiv(
		C.GLuint(program),
		C.GLenum(pname),
		(*C.GLint)(params))

}
func GetProgramInfoLog(
	program Uint, bufsize Sizei,
	length *Sizei, infolog *string) {
	s := glString(*infolog)
	C.glGetProgramInfoLog(
		C.GLuint(program),
		C.GLsizei(bufsize),
		(*C.GLsizei)(length),
		s)
	infolog = goString(s)
}
func GetRenderbufferParameteriv(
	target Enum, pname Enum, params *Int) {
	C.glGetRenderbufferParameteriv(
		C.GLenum(target),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetShaderiv(
	shader Uint, pname Enum, params *Int) {
	C.glGetShaderiv(
		C.GLuint(shader),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetShaderInfoLog(
	shader Uint, bufsize Sizei,
	length *Sizei, infolog *string) {
	s := glString(*infolog)
	C.glGetShaderInfoLog(
		C.GLuint(shader),
		C.GLsizei(bufsize),
		(*C.GLsizei)(length),
		s)
	infolog = goString(s)
}
func GetShaderPrecisionFormat(
	shadertype Enum, precisiontype Enum,
	range_ *Int, precision *Int) {
	C.glGetShaderPrecisionFormat(
		C.GLenum(shadertype),
		C.GLenum(precisiontype),
		(*C.GLint)(range_),
		(*C.GLint)(precision))
}
func GetShaderSource(
	shader Uint, bufsize Sizei,
	length *Sizei, source *string) {
	s := glString(*source)
	C.glGetShaderSource(
		C.GLuint(shader),
		C.GLsizei(bufsize),
		(*C.GLsizei)(length),
		s)
	source = goString(s)
}
func GetTexParameterfv(
	target Enum, pname Enum, params *Float) {
	C.glGetTexParameterfv(
		C.GLenum(target),
		C.GLenum(pname),
		(*C.GLfloat)(params))
}
func GetTexParameteriv(
	target Enum, pname Enum, params *Int) {
	C.glGetTexParameteriv(
		C.GLenum(target),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetUniformfv(
	program Uint, location Int, params *Float) {
	C.glGetUniformfv(
		C.GLuint(program),
		C.GLint(location),
		(*C.GLfloat)(params))
}
func GetUniformiv(
	program Uint, location Int, params *Int) {
	C.glGetUniformiv(
		C.GLuint(program),
		C.GLint(location),
		(*C.GLint)(params))
}
func GetVertexAttribfv(
	index Uint, pname Enum, params *Float) {
	C.glGetVertexAttribfv(
		C.GLuint(index),
		C.GLenum(pname),
		(*C.GLfloat)(params))
}
func GetVertexAttribiv(
	index Uint, pname Enum, params *Int) {
	C.glGetVertexAttribiv(
		C.GLuint(index),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetVertexAttribPointerv(
	index Uint, pname Enum, pointer *Void) {
	C.glGetVertexAttribPointerv(
		C.GLuint(index),
		C.GLenum(pname),
		(*unsafe.Pointer)(*pointer))
}
func Hint(
	target Enum, mode Enum) {
	C.glHint(
		C.GLenum(target),
		C.GLenum(mode))
}
func LineWidth(
	width Float) {
	C.glLineWidth(
		C.GLfloat(width))
}
func LinkProgram(
	program Uint) {
	C.glLinkProgram(
		C.GLuint(program))
}
func PixelStorei(
	pname Enum, param Int) {
	C.glPixelStorei(
		C.GLenum(pname),
		C.GLint(param))
}
func PolygonOffset(
	factor Float, units Float) {
	C.glPolygonOffset(
		C.GLfloat(factor),
		C.GLfloat(units))
}
func ReadPixels(
	x Int, y Int, width Sizei, height Sizei,
	format Enum, type_ Enum, pixels Void) {
	C.glReadPixels(
		C.GLint(x),
		C.GLint(y),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(type_),
		unsafe.Pointer(pixels))
}
func ReleaseShaderCompiler() {
	C.glReleaseShaderCompiler()
}
func RenderbufferStorage(
	target Enum, internalformat Enum,
	width Sizei, height Sizei) {
	C.glRenderbufferStorage(
		C.GLenum(target),
		C.GLenum(internalformat),
		C.GLsizei(width),
		C.GLsizei(height))
}
func SampleCoverage(
	value Clampf, invert Boolean) {
	C.glSampleCoverage(
		C.GLclampf(value),
		glBoolean(invert))
}
func Scissor(
	x Int, y Int, width Sizei, height Sizei) {
	C.glScissor(
		C.GLint(x),
		C.GLint(y),
		C.GLsizei(width),
		C.GLsizei(height))
}
func ShaderBinary(
	n Sizei, shaders *Uint,
	binaryformat Enum, binary Void, length Sizei) {
	C.glShaderBinary(
		C.GLsizei(n),
		(*C.GLuint)(shaders),
		C.GLenum(binaryformat),
		unsafe.Pointer(binary),
		C.GLsizei(length))
}
func ShaderSource(
	shader Uint, count Sizei,
	string_ *string, length *Int) {
	s := glString(*string_)
	C.glShaderSource(
		C.GLuint(shader),
		C.GLsizei(count),
		&s,
		(*C.GLint)(length))
	string_ = goString(s)
}
func StencilFunc(
	func_ Enum, ref Int, mask Uint) {
	C.glStencilFunc(
		C.GLenum(func_),
		C.GLint(ref),
		C.GLuint(mask))
}
func StencilFuncSeparate(
	face Enum, func_ Enum, ref Int, mask Uint) {
	C.glStencilFuncSeparate(
		C.GLenum(face),
		C.GLenum(func_),
		C.GLint(ref),
		C.GLuint(mask))
}
func StencilMask(
	mask Uint) {
	C.glStencilMask(
		C.GLuint(mask))
}
func StencilMaskSeparate(
	face Enum, mask Uint) {
	C.glStencilMaskSeparate(
		C.GLenum(face),
		C.GLuint(mask))
}
func StencilOp(
	fail Enum, zfail Enum, zpass Enum) {
	C.glStencilOp(
		C.GLenum(fail),
		C.GLenum(zfail),
		C.GLenum(zpass))
}
func StencilOpSeparate(
	face Enum, fail Enum,
	zfail Enum, zpass Enum) {
	C.glStencilOpSeparate(
		C.GLenum(face),
		C.GLenum(fail),
		C.GLenum(zfail),
		C.GLenum(zpass))
}
func TexImage2D(
	target Enum, level Int, internalformat Int,
	width Sizei, height Sizei, border Int, format Enum,
	type_ Enum, pixels Void) {
	C.glTexImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(internalformat),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLint(border),
		C.GLenum(format),
		C.GLenum(type_),
		unsafe.Pointer(pixels))
}
func TexParameterf(
	target Enum, pname Enum, param Float) {
	C.glTexParameterf(
		C.GLenum(target),
		C.GLenum(pname),
		C.GLfloat(param))
}
func TexParameterfv(
	target Enum, pname Enum, params *Float) {
	C.glTexParameterfv(
		C.GLenum(target),
		C.GLenum(pname),
		(*C.GLfloat)(params))
}
func TexParameteri(
	target Enum, pname Enum, param Int) {
	C.glTexParameteri(
		C.GLenum(target),
		C.GLenum(pname),
		C.GLint(param))
}
func TexParameteriv(
	target Enum, pname Enum, params *Int) {
	C.glTexParameteriv(
		C.GLenum(target),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func TexSubImage2D(
	target Enum, level Int, xoffset Int, yoffset Int,
	width Sizei, height Sizei, format Enum, type_ Enum, pixels Void) {
	C.glTexSubImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(xoffset),
		C.GLint(yoffset),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(type_),
		unsafe.Pointer(pixels))
}
func Uniform1f(
	location Int, x Float) {
	C.glUniform1f(
		C.GLint(location),
		C.GLfloat(x))
}
func Uniform1fv(
	location Int, count Sizei, v *Float) {
	C.glUniform1fv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLfloat)(v))
}
func Uniform1i(
	location Int, x Int) {
	C.glUniform1i(
		C.GLint(location),
		C.GLint(x))
}
func Uniform1iv(
	location Int, count Sizei, v *Int) {
	C.glUniform1iv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLint)(v))
}
func Uniform2f(
	location Int, x Float, y Float) {
	C.glUniform2f(
		C.GLint(location),
		C.GLfloat(x),
		C.GLfloat(y))
}
func Uniform2fv(
	location Int, count Sizei, v *Float) {
	C.glUniform2fv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLfloat)(v))
}
func Uniform2i(
	location Int, x Int, y Int) {
	C.glUniform2i(
		C.GLint(location),
		C.GLint(x),
		C.GLint(y))
}
func Uniform2iv(
	location Int, count Sizei, v *Int) {
	C.glUniform2iv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLint)(v))
}
func Uniform3f(
	location Int, x Float, y Float, z Float) {
	C.glUniform3f(
		C.GLint(location),
		C.GLfloat(x),
		C.GLfloat(y),
		C.GLfloat(z))
}
func Uniform3fv(
	location Int, count Sizei, v *Float) {
	C.glUniform3fv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLfloat)(v))
}
func Uniform3i(
	location Int, x Int, y Int, z Int) {
	C.glUniform3i(
		C.GLint(location),
		C.GLint(x),
		C.GLint(y),
		C.GLint(z))
}
func Uniform3iv(
	location Int, count Sizei, v *Int) {
	C.glUniform3iv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLint)(v))
}
func Uniform4f(
	location Int, x Float, y Float, z Float, w Float) {
	C.glUniform4f(
		C.GLint(location),
		C.GLfloat(x),
		C.GLfloat(y),
		C.GLfloat(z),
		C.GLfloat(w))
}
func Uniform4fv(
	location Int, count Sizei, v *Float) {
	C.glUniform4fv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLfloat)(v))
}
func Uniform4i(
	location Int, x Int, y Int, z Int, w Int) {
	C.glUniform4i(
		C.GLint(location),
		C.GLint(x),
		C.GLint(y),
		C.GLint(z),
		C.GLint(w))
}
func Uniform4iv(
	location Int, count Sizei, v *Int) {
	C.glUniform4iv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLint)(v))
}
func UniformMatrix2fv(
	location Int, count Sizei,
	transpose Boolean, value *Float) {
	C.glUniformMatrix2fv(
		C.GLint(location),
		C.GLsizei(count),
		glBoolean(transpose),
		(*C.GLfloat)(value))
}
func UniformMatrix3fv(
	location Int, count Sizei,
	transpose Boolean, value *Float) {
	C.glUniformMatrix3fv(
		C.GLint(location),
		C.GLsizei(count),
		glBoolean(transpose),
		(*C.GLfloat)(value))
}
func UniformMatrix4fv(
	location Int, count Sizei,
	transpose Boolean, value *Float) {
	C.glUniformMatrix4fv(
		C.GLint(location),
		C.GLsizei(count),
		glBoolean(transpose),
		(*C.GLfloat)(value))
}
func UseProgram(
	program Uint) {
	C.glUseProgram(
		C.GLuint(program))
}
func ValidateProgram(
	program Uint) {
	C.glValidateProgram(
		C.GLuint(program))
}
func VertexAttrib1f(
	indx Uint, x Float) {
	C.glVertexAttrib1f(
		C.GLuint(indx),
		C.GLfloat(x))
}
func VertexAttrib1fv(
	indx Uint, values *Float) {
	C.glVertexAttrib1fv(
		C.GLuint(indx),
		(*C.GLfloat)(values))
}
func VertexAttrib2f(
	indx Uint, x Float, y Float) {
	C.glVertexAttrib2f(
		C.GLuint(indx),
		C.GLfloat(x),
		C.GLfloat(y))
}
func VertexAttrib2fv(
	indx Uint, values *Float) {
	C.glVertexAttrib2fv(
		C.GLuint(indx),
		(*C.GLfloat)(values))
}
func VertexAttrib3f(
	indx Uint, x Float, y Float, z Float) {
	C.glVertexAttrib3f(
		C.GLuint(indx),
		C.GLfloat(x),
		C.GLfloat(y),
		C.GLfloat(z))
}
func VertexAttrib3fv(
	indx Uint, values *Float) {
	C.glVertexAttrib3fv(
		C.GLuint(indx),
		(*C.GLfloat)(values))
}
func VertexAttrib4f(
	indx Uint, x Float, y Float, z Float, w Float) {
	C.glVertexAttrib4f(
		C.GLuint(indx),
		C.GLfloat(x),
		C.GLfloat(y),
		C.GLfloat(z),
		C.GLfloat(w))
}
func VertexAttrib4fv(
	indx Uint, values *Float) {
	C.glVertexAttrib4fv(
		C.GLuint(indx),
		(*C.GLfloat)(values))
}
func VertexAttribPointer(
	indx Uint, size Int, type_ Enum,
	normalized Boolean, stride Sizei, ptr Void) {
	C.glVertexAttribPointer(
		C.GLuint(indx),
		C.GLint(size),
		C.GLenum(type_),
		glBoolean(normalized),
		C.GLsizei(stride),
		unsafe.Pointer(ptr))
}
func Viewport(
	x Int, y Int, width Sizei, height Sizei) {
	C.glViewport(
		C.GLint(x),
		C.GLint(y),
		C.GLsizei(width),
		C.GLsizei(height))
}
func IsBuffer(buffer Uint) Boolean {
	return *goBoolean(C.glIsBuffer(
		C.GLuint(buffer)))
}
func IsEnabled(cap Enum) Boolean {
	return *goBoolean(C.glIsEnabled(
		C.GLenum(cap)))
}
func IsFramebuffer(framebuffer Uint) Boolean {
	return *goBoolean(C.glIsFramebuffer(
		C.GLuint(framebuffer)))
}
func IsProgram(program Uint) Boolean {
	return *goBoolean(C.glIsProgram(
		C.GLuint(program)))
}
func IsRenderbuffer(renderbuffer Uint) Boolean {
	return *goBoolean(C.glIsRenderbuffer(
		C.GLuint(renderbuffer)))
}
func IsShader(shader Uint) Boolean {
	return *goBoolean(C.glIsShader(
		C.GLuint(shader)))
}
func IsTexture(texture Uint) Boolean {
	return *goBoolean(C.glIsTexture(
		C.GLuint(texture)))
}
func GetError() Enum {
	return Enum(C.glGetError())
}
func CheckFramebufferStatus(target Enum) Enum {
	return Enum(C.glCheckFramebufferStatus(
		C.GLenum(target)))
}
func CreateProgram() Uint {
	return Uint(C.glCreateProgram())
}
func CreateShader(type_ Enum) Uint {
	return Uint(C.glCreateShader(
		C.GLenum(type_)))
}
func GetAttribLocation(program Uint, name *string) uintptr {
	s := glString(*name)	
	return uintptr(C.glGetAttribLocation(
		C.GLuint(program),
		s))
}
func GetUniformLocation(program Uint, name *string) uintptr {
	s := glString(*name)
	return uintptr(C.glGetUniformLocation(
		C.GLuint(program),
		s))
}
func GetString(name Enum) string {
	return string(byte(*C.glGetString(
		C.GLenum(name))))
}
