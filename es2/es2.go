package es2

//#cgo linux pkg-config: glesv2
//#include <stdlib.h>
//#include <GLES2/gl2.h>
//#include <GLES2/gl2ext.h>
//#include <GLES2/gl2platform.h>
import "C"
import "unsafe"

func ActiveTexture(texture int) {
	C.glActiveTexture(C.GLenum(texture))
}
func AttachShader(program, shader int) {
	C.glAttachShader(C.GLuint(program), C.GLuint(shader))
}
func BindAttribLocation(program, index int, name string) {
	C.glBindAttribLocation(C.GLuint(program), C.GLuint(index), CString(name))
}
func BindBuffer(target, buffer int) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(buffer))
}
func BindFramebuffer(target, framebuffer int) {
	C.glBindFramebuffer(C.GLenum(target), C.GLuint(framebuffer))
}
func BindRenderbuffer(target, renderbuffer int) {
	C.glBindRenderbuffer(C.GLenum(target), C.GLuint(renderbuffer))
}
func BindTexture(target, texture int) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}
func BlendColor(red, green, blue, alpha float32) {
	C.glBlendColor(
		C.GLclampf(red), C.GLclampf(green),
		C.GLclampf(blue), C.GLclampf(alpha))
}
func BlendEquation(mode int) { C.glBlendEquation(C.GLenum(mode)) }

func BlendEquationSeparate(modeRGB, modeAlpha int) {
	C.glBlendEquationSeparate(C.GLenum(modeRGB), C.GLenum(modeAlpha))
}
func BlendFunc(sfactor, dfactor int) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
}
func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha int) {
	C.glBlendFuncSeparate(
		C.GLenum(srcRGB), C.GLenum(dstRGB),
		C.GLenum(srcAlpha), C.GLenum(dstAlpha))
}
func BufferData(target int, size int, data Void, usage int) {
	C.glBufferData(
		C.GLenum(target), C.GLsizeiptr(size),
		unsafe.Pointer(data), C.GLenum(usage))
}
func BufferSubData(target int, offset int, size int, data Void) {
	C.glBufferSubData(
		C.GLenum(target), C.GLintptr(offset),
		C.GLsizeiptr(size), unsafe.Pointer(data))
}
func Clear(mask int) { C.glClear(C.GLbitfield(mask)) }

func ClearColor(red, green, blue, alpha float32) {
	C.glClearColor(
		C.GLclampf(red), C.GLclampf(green),
		C.GLclampf(blue), C.GLclampf(alpha))
}

func ClearDepthf(depth float32) { C.glClearDepthf(C.GLclampf(depth)) }

func ClearStencil(s int) { C.glClearStencil(C.GLint(s)) }

func ColorMask(red, green, blue, alpha bool) {
	C.glColorMask(
		glBoolean(red), glBoolean(green),
		glBoolean(blue), glBoolean(alpha))
}

func CompileShader(shader int) { C.glCompileShader(C.GLuint(shader)) }

func CompressedTexImage2D(target int, level int, internalformat,
	width, height, border, imageSize int, data Void) {
	C.glCompressedTexImage2D(
		C.GLenum(target), C.GLint(level), C.GLenum(internalformat),
		C.GLsizei(width), C.GLsizei(height),
		C.GLint(border), C.GLsizei(imageSize),
		unsafe.Pointer(data))
}
func CompressedTexSubImage2D(target int, level, xoffset, yoffset int,
	width, height, format, imageSize int, data Void) {
	C.glCompressedTexSubImage2D(
		C.GLenum(target), C.GLint(level),
		C.GLint(xoffset), C.GLint(yoffset),
		C.GLsizei(width), C.GLsizei(height),
		C.GLenum(format), C.GLsizei(imageSize),
		unsafe.Pointer(data))
}
func CopyTexImage2D(target int, level int, internalformat int,
	x, y int, width, height int, border int) {
	C.glCopyTexImage2D(
		C.GLenum(target), C.GLint(level), C.GLenum(internalformat),
		C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height),
		C.GLint(border))
}
func CopyTexSubImage2D(target int, level, xoffset, yoffset,
	x, y int, width, height int) {
	C.glCopyTexSubImage2D(
		C.GLenum(target), C.GLint(level),
		C.GLint(xoffset), C.GLint(yoffset),
		C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
func CullFace(mode int) { C.glCullFace(C.GLenum(mode)) }

func DeleteBuffers(n int, buffers *uint32) {
	C.glDeleteBuffers(C.GLsizei(n), (*C.GLuint)(buffers))
}
func DeleteFramebuffers(n int, framebuffers *uint32) {
	C.glDeleteFramebuffers(C.GLsizei(n), (*C.GLuint)(framebuffers))
}
func DeleteProgram(program int) {
	C.glDeleteProgram(C.GLuint(program))
}
func DeleteRenderbuffers(n int, renderbuffers *uint32) {
	C.glDeleteRenderbuffers(C.GLsizei(n), (*C.GLuint)(renderbuffers))
}
func DeleteShader(shader int) {
	C.glDeleteShader(C.GLuint(shader))
}
func DeleteTextures(n int, textures *uint32) {
	C.glDeleteTextures(C.GLsizei(n), (*C.GLuint)(textures))
}
func DepthFunc(f int) { C.glDepthFunc(C.GLenum(f)) }

func DepthMask(flag bool) { C.glDepthMask(glBoolean(flag)) }

func DepthRangef(zNear, zFar float32) {
	C.glDepthRangef(C.GLclampf(zNear), C.GLclampf(zFar))
}
func DetachShader(program, shader int) {
	C.glDetachShader(C.GLuint(program), C.GLuint(shader))
}
func Disable(cap int) { C.glDisable(C.GLenum(cap)) }

func DisableVertexAttribArray(index int) {
	C.glDisableVertexAttribArray(C.GLuint(index))
}
func DrawArrays(mode int, first, count int) {
	C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}
func DrawElements(mode, count, type_ int, indices Void) {
	C.glDrawElements(
		C.GLenum(mode), C.GLsizei(count),
		C.GLenum(type_), unsafe.Pointer(indices))
}

func EnableVertexAttribArray(index int) {
	C.glEnableVertexAttribArray(C.GLuint(index))
}
func Enable(cap int) { C.glEnable(C.GLenum(cap)) }
func Finish()        { C.glFinish() }
func Flush()         { C.glFlush() }

func FramebufferRenderbuffer(target, attachment, renderbuffertarget, renderbuffer int) {
	C.glFramebufferRenderbuffer(
		C.GLenum(target), C.GLenum(attachment),
		C.GLenum(renderbuffertarget), C.GLuint(renderbuffer))
}
func FramebufferTexture2D(target, attachment, textarget, texture, level int) {
	C.glFramebufferTexture2D(
		C.GLenum(target), C.GLenum(attachment),
		C.GLenum(textarget), C.GLuint(texture), C.GLint(level))
}
func FrontFace(mode int) {
	C.glFrontFace(C.GLenum(mode))
}
func GenBuffers(n int, buffers Void) {
	C.glGenBuffers(C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}
func GenerateMipmap(target int) {
	C.glGenerateMipmap(C.GLenum(target))
}
func GenFramebuffers(n int, framebuffers *uint32) {
	C.glGenFramebuffers(C.GLsizei(n), (*C.GLuint)(framebuffers))
}
func GenRenderbuffers(n int, renderbuffers *uint32) {
	C.glGenRenderbuffers(C.GLsizei(n), (*C.GLuint)(renderbuffers))
}
func GenTextures(n int, textures Void) {
	C.glGenTextures(C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)))
}
func GetActiveAttrib(program, index, bufsize int) (size int32, type_ uint32, name string) {
	cs := CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.glGetActiveAttrib(C.GLuint(program), C.GLuint(index),
		C.GLsizei(bufsize), nil, (*C.GLint)(&size), (*C.GLenum)(&type_), cs)
	name = GoString(cs)
	return

}
func GetActiveUniform(program, index, bufsize int) (size int32, type_ uint32, name string) {
	cs := CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.glGetActiveUniform(C.GLuint(program), C.GLuint(index),
		C.GLsizei(bufsize), nil, (*C.GLint)(&size), (*C.GLenum)(&type_), cs)
	name = GoString(cs)
	return

}
func GetAttachedShaders(program int, shaders []uint32) []uint32 {
	C.glGetAttachedShaders(
		C.GLuint(program), C.GLsizei(cap(shaders)), nil, (*C.GLuint)(&shaders[0]))
	return shaders
}
func GetBooleanv(pname int, params []bool) []bool {
	C.glGetBooleanv(C.GLenum(pname), (*C.GLboolean)(unsafe.Pointer(&params[0])))
	return params
}
func GetBufferParameteriv(target, pname int, params []int32) []int32 {
	C.glGetBufferParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetFloatv(pname int, params []float32) []float32 {
	C.glGetFloatv(C.GLenum(pname), (*C.GLfloat)(&params[0]))
	return params
}
func GetFramebufferAttachmentParameteriv(
	target, attachment, pname int, params []int32) []int32 {
	C.glGetFramebufferAttachmentParameteriv(
		C.GLenum(target), C.GLenum(attachment),
		C.GLenum(pname), (*C.GLint)(&params[0]))
	return params

}
func GetIntegerv(pname int, params []int32) []int32 {
	C.glGetIntegerv(C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetProgramiv(program, pname int, params []int32) []int32 {
	C.glGetProgramiv(C.GLuint(program), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetProgramInfoLog(program, bufsize int) string {
	cs := CString("")
	defer C.free(unsafe.Pointer(cs))
	C.glGetProgramInfoLog(C.GLuint(program), C.GLsizei(bufsize), nil, cs)
	return GoString(cs)

}
func GetRenderbufferParameteriv(target, pname int, params []int32) []int32 {
	C.glGetRenderbufferParameteriv(
		C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetShaderiv(shader int, pname int, params []int32) []int32 {
	C.glGetShaderiv(C.GLuint(shader), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetShaderInfoLog(shader int, bufsize int) string {
	cs := CString("")
	defer C.free(unsafe.Pointer(cs))
	C.glGetShaderInfoLog(C.GLuint(shader), C.GLsizei(bufsize), nil, cs)
	return GoString(cs)
}
func GetShaderPrecisionFormat(shadertype int, precisiontype int) (
	range_ int32, precision int32) {
	C.glGetShaderPrecisionFormat(
		C.GLenum(shadertype), C.GLenum(precisiontype),
		(*C.GLint)(&range_), (*C.GLint)(&precision))
	return
}
func GetShaderSource(shader int, bufsize int) string {
	cs := CString("")
	defer C.free(unsafe.Pointer(cs))
	C.glGetShaderSource(C.GLuint(shader), C.GLsizei(bufsize), nil, cs)
	return GoString(cs)
}
func GetTexParameterfv(target int, pname int, params []float32) []float32 {
	C.glGetTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
	return params
}
func GetTexParameteriv(target int, pname int, params []int32) []int32 {
	C.glGetTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetUniformfv(program int, location int, params []float32) []float32 {
	C.glGetUniformfv(C.GLuint(program), C.GLint(location), (*C.GLfloat)(&params[0]))
	return params
}
func GetUniformiv(program int, location int, params []int32) []int32 {
	C.glGetUniformiv(C.GLuint(program), C.GLint(location), (*C.GLint)(&params[0]))
	return params
}
func GetVertexAttribfv(index int, pname int, params []float32) []float32 {
	C.glGetVertexAttribfv(C.GLuint(index), C.GLenum(pname), (*C.GLfloat)(&params[0]))
	return params
}
func GetVertexAttribiv(index int, pname int, params []int32) []int32 {
	C.glGetVertexAttribiv(C.GLuint(index), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetVertexAttribPointerv(index int, pname int, pointer []Void) []Void {
	C.glGetVertexAttribPointerv(C.GLuint(index), C.GLenum(pname), (*unsafe.Pointer)(&pointer[0]))
	return pointer
}
func Hint(target int, mode int) {
	C.glHint(C.GLenum(target), C.GLenum(mode))
}
func LineWidth(width float32) {
	C.glLineWidth(C.GLfloat(width))
}
func LinkProgram(program int) {
	C.glLinkProgram(C.GLuint(program))
}
func PixelStorei(pname, param int) {
	C.glPixelStorei(C.GLenum(pname), C.GLint(param))
}
func PolygonOffset(factor float32, units float32) {
	C.glPolygonOffset(C.GLfloat(factor), C.GLfloat(units))
}
func ReadPixels(x, y, width, height, format, type_ int, pixels Void) {
	C.glReadPixels(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height),
		C.GLenum(format), C.GLenum(type_), unsafe.Pointer(pixels))
}
func ReleaseShaderCompiler() {
	C.glReleaseShaderCompiler()
}
func RenderbufferStorage(target, internalformat, width, height int) {
	C.glRenderbufferStorage(C.GLenum(target), C.GLenum(internalformat),
		C.GLsizei(width), C.GLsizei(height))
}
func SampleCoverage(value float32, invert bool) {
	C.glSampleCoverage(C.GLclampf(value), glBoolean(invert))
}
func Scissor(x, y, width, height int) {
	C.glScissor(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
func ShaderBinary(n int, shaders Void, binaryformat int, binary Void, length int) {
	C.glShaderBinary(C.GLsizei(n), (*C.GLuint)(shaders),
		C.GLenum(binaryformat), unsafe.Pointer(binary), C.GLsizei(length))
}
func ShaderSource(shader int, s string) {
	cs := CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.glShaderSource(C.GLuint(shader), C.GLsizei(1), &cs, nil)
}
func StencilFunc(func_ int, ref int, mask int) {
	C.glStencilFunc(C.GLenum(func_), C.GLint(ref), C.GLuint(mask))
}
func StencilFuncSeparate(face int, func_ int, ref int, mask int) {
	C.glStencilFuncSeparate(
		C.GLenum(face), C.GLenum(func_), C.GLint(ref), C.GLuint(mask))
}
func StencilMask(mask int) {
	C.glStencilMask(C.GLuint(mask))
}
func StencilMaskSeparate(face int, mask int) {
	C.glStencilMaskSeparate(C.GLenum(face), C.GLuint(mask))
}
func StencilOp(fail, zfail, zpass int) {
	C.glStencilOp(C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}
func StencilOpSeparate(face, fail, zfail, zpass int) {
	C.glStencilOpSeparate(
		C.GLenum(face), C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}
func TexImage2D(target, level, internalformat, width, height,
	border, format, type_ int, pixels Void) {
	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GLint(internalformat),
		C.GLsizei(width), C.GLsizei(height), C.GLint(border),
		C.GLenum(format), C.GLenum(type_), unsafe.Pointer(pixels))
}
func TexParameterf(target, pname int, param float32) {
	C.glTexParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}
func TexParameterfv(target, pname int, params []float32) []float32 {
	C.glTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
	return params
}
func TexParameteri(target, pname, param int) {
	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}
func TexParameteriv(target int, pname int, params []int32) []int32 {
	C.glTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func TexSubImage2D(target, level, xoffset, yoffset, width, height,
	format, type_ int, pixels Void) {

	C.glTexSubImage2D(C.GLenum(target), C.GLint(level),
		C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height),
		C.GLenum(format), C.GLenum(type_),
		unsafe.Pointer(pixels))
}
func Uniform1f(location int, x float32) {
	C.glUniform1f(C.GLint(location), C.GLfloat(x))
}
func Uniform1fv(location, count int, value []float32) {
	C.glUniform1fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&value[0]))
}
func Uniform1i(location, x int) {
	C.glUniform1i(C.GLint(location), C.GLint(x))
}
func Uniform1iv(location, count int, value []int32) {
	C.glUniform1iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(&value[0]))
}
func Uniform2f(location int, x, y float32) {
	C.glUniform2f(C.GLint(location), C.GLfloat(x), C.GLfloat(y))
}
func Uniform2fv(location, count int, value []float32) {
	C.glUniform2fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&value[0]))
}
func Uniform2i(location, x, y int) {
	C.glUniform2i(C.GLint(location), C.GLint(x), C.GLint(y))
}
func Uniform2iv(location, count int, value []int32) {
	C.glUniform2iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(&value[0]))
}
func Uniform3f(location int, x, y, z float32) {
	C.glUniform3f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}
func Uniform3fv(location, count int, value []float32) {
	C.glUniform3fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&value[0]))
}
func Uniform3i(location, x, y, z int) {
	C.glUniform3i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z))
}
func Uniform3iv(location, count int, value []int32) {
	C.glUniform3iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(&value[0]))
}
func Uniform4f(location int, x, y, z, w float32) {
	C.glUniform4f(
		C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}
func Uniform4fv(location, count int, value []float32) {
	C.glUniform4fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&value[0]))
}
func Uniform4i(location, x int, y int, z int, w int) {
	C.glUniform4i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}
func Uniform4iv(location, count int, value []int32) {
	C.glUniform4iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(&value[0]))
}
func UniformMatrix2fv(location, count int, transpose bool, value []float32) {
	C.glUniformMatrix2fv(C.GLint(location), C.GLsizei(count), glBoolean(transpose),
		(*C.GLfloat)(&value[0]))
}
func UniformMatrix3fv(location, count int, transpose bool, value []float32) {
	C.glUniformMatrix3fv(C.GLint(location), C.GLsizei(count), glBoolean(transpose),
		(*C.GLfloat)(&value[0]))
}
func UniformMatrix4fv(location, count int, transpose bool, value []float32) {
	C.glUniformMatrix4fv(C.GLint(location), C.GLsizei(count),
		glBoolean(transpose), (*C.GLfloat)(&value[0]))
}
func UseProgram(program int) {
	C.glUseProgram(C.GLuint(program))
}
func ValidateProgram(program int) {
	C.glValidateProgram(C.GLuint(program))
}
func VertexAttrib1f(indx int, x float32) {
	C.glVertexAttrib1f(C.GLuint(indx), C.GLfloat(x))
}
func VertexAttrib1fv(indx int, values []float32) {
	C.glVertexAttrib1fv(C.GLuint(indx), (*C.GLfloat)(&values[0]))
}
func VertexAttrib2f(indx int, x, y float32) {
	C.glVertexAttrib2f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y))
}
func VertexAttrib2fv(indx int, values []float32) {
	C.glVertexAttrib2fv(C.GLuint(indx), (*C.GLfloat)(&values[0]))
}
func VertexAttrib3f(indx int, x, y, z float32) {
	C.glVertexAttrib3f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}
func VertexAttrib3fv(indx int, values []float32) {
	C.glVertexAttrib3fv(C.GLuint(indx), (*C.GLfloat)(&values[0]))
}
func VertexAttrib4f(indx int, x, y, z, w float32) {
	C.glVertexAttrib4f(C.GLuint(indx),
		C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}
func VertexAttrib4fv(indx int, values []float32) {
	C.glVertexAttrib4fv(C.GLuint(indx), (*C.GLfloat)(&values[0]))
}
func VertexAttribPointer(indx, size, type_ int, normalized bool, stride int, ptr Void) {
	C.glVertexAttribPointer(C.GLuint(indx), C.GLint(size), C.GLenum(type_),
		glBoolean(normalized), C.GLsizei(stride), unsafe.Pointer(ptr))
}
func Viewport(x int, y int, width int, height int) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
func IsBuffer(buffer int) bool {
	return *goBoolean(C.glIsBuffer(C.GLuint(buffer)))
}
func IsEnabled(cap int) bool {
	return *goBoolean(C.glIsEnabled(C.GLenum(cap)))
}
func IsFramebuffer(framebuffer int) bool {
	return *goBoolean(C.glIsFramebuffer(C.GLuint(framebuffer)))
}
func IsProgram(program int) bool {
	return *goBoolean(C.glIsProgram(C.GLuint(program)))
}
func IsRenderbuffer(renderbuffer int) bool {
	return *goBoolean(C.glIsRenderbuffer(C.GLuint(renderbuffer)))
}
func IsShader(shader int) bool {
	return *goBoolean(C.glIsShader(C.GLuint(shader)))
}
func IsTexture(texture int) bool {
	return *goBoolean(C.glIsTexture(C.GLuint(texture)))
}
func GetError() int {
	return int(C.glGetError())
}
func CheckFramebufferStatus(target int) int {
	return int(C.glCheckFramebufferStatus(C.GLenum(target)))
}
func CreateProgram() int {
	return int(C.glCreateProgram())
}
func CreateShader(type_ int) int {
	return int(C.glCreateShader(C.GLenum(type_)))
}
func GetAttribLocation(program int, name string) int {
	return int(C.glGetAttribLocation(C.GLuint(program), CString(name)))
}
func GetUniformLocation(program int, name string) int {
	return int(C.glGetUniformLocation(C.GLuint(program), CString(name)))
}
func GetString(name int) string {
	s := (*C.GLchar)(unsafe.Pointer(C.glGetString(C.GLenum(name))))
	defer C.free(unsafe.Pointer(s))
	return GoString(s)

}
