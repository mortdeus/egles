package es2

//#cgo linux pkg-config: glesv2
//#include <stdlib.h>
//#include <GLES2/gl2.h>
//#include <GLES2/gl2ext.h>
//#include <GLES2/gl2platform.h>
import "C"
import "unsafe"

func ActiveTexture(texture uint) {
	C.glActiveTexture(C.GLenum(texture))
}
func AttachShader(program, shader uint) {
	C.glAttachShader(C.GLuint(program), C.GLuint(shader))
}
func BindAttribLocation(program, index uint, name string) {
	C.glBindAttribLocation(C.GLuint(program), C.GLuint(index), CString(name))
}
func BindBuffer(target, buffer uint) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(buffer))
}
func BindFramebuffer(target, framebuffer uint) {
	C.glBindFramebuffer(C.GLenum(target), C.GLuint(framebuffer))
}
func BindRenderbuffer(target, renderbuffer uint) {
	C.glBindRenderbuffer(C.GLenum(target), C.GLuint(renderbuffer))
}
func BindTexture(target, texture uint) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}
func BlendColor(red, green, blue, alpha float32) {
	C.glBlendColor(
		C.GLclampf(red), C.GLclampf(green),
		C.GLclampf(blue), C.GLclampf(alpha))
}
func BlendEquation(mode uint) { C.glBlendEquation(C.GLenum(mode)) }

func BlendEquationSeparate(modeRGB, modeAlpha uint) {
	C.glBlendEquationSeparate(C.GLenum(modeRGB), C.GLenum(modeAlpha))
}
func BlendFunc(sfactor, dfactor uint) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
}
func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha uint) {
	C.glBlendFuncSeparate(
		C.GLenum(srcRGB), C.GLenum(dstRGB),
		C.GLenum(srcAlpha), C.GLenum(dstAlpha))
}
func BufferData(target uint, size int, data Void, usage uint) {
	C.glBufferData(
		C.GLenum(target), C.GLsizeiptr(size),
		unsafe.Pointer(data), C.GLenum(usage))
}
func BufferSubData(target uint, offset int, size int, data Void) {
	C.glBufferSubData(
		C.GLenum(target), C.GLintptr(offset),
		C.GLsizeiptr(size), unsafe.Pointer(data))
}
func Clear(mask uint) { C.glClear(C.GLbitfield(mask)) }

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

func CompileShader(shader uint) { C.glCompileShader(C.GLuint(shader)) }

func CompressedTexImage2D(target uint, level int, internalformat uint,
	width, height, border, imageSize int, data Void) {
	C.glCompressedTexImage2D(
		C.GLenum(target), C.GLint(level), C.GLenum(internalformat),
		C.GLsizei(width), C.GLsizei(height),
		C.GLint(border), C.GLsizei(imageSize),
		unsafe.Pointer(data))
}
func CompressedTexSubImage2D(target uint, level, xoffset, yoffset int,
	width, height int, format uint, imageSize int, data Void) {
	C.glCompressedTexSubImage2D(
		C.GLenum(target), C.GLint(level),
		C.GLint(xoffset), C.GLint(yoffset),
		C.GLsizei(width), C.GLsizei(height),
		C.GLenum(format), C.GLsizei(imageSize),
		unsafe.Pointer(data))
}
func CopyTexImage2D(target uint, level int, internalformat uint,
	x, y, width, height, border int) {
	C.glCopyTexImage2D(
		C.GLenum(target), C.GLint(level), C.GLenum(internalformat),
		C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height),
		C.GLint(border))
}
func CopyTexSubImage2D(target uint, level, xoffset, yoffset,
	x, y, width, height int) {
	C.glCopyTexSubImage2D(
		C.GLenum(target), C.GLint(level),
		C.GLint(xoffset), C.GLint(yoffset),
		C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
func CullFace(mode uint) { C.glCullFace(C.GLenum(mode)) }

func DeleteBuffers(n int, buffers *uint32) {
	C.glDeleteBuffers(C.GLsizei(n), (*C.GLuint)(buffers))
}
func DeleteFramebuffers(n int, framebuffers *uint32) {
	C.glDeleteFramebuffers(C.GLsizei(n), (*C.GLuint)(framebuffers))
}
func DeleteProgram(program uint) {
	C.glDeleteProgram(C.GLuint(program))
}
func DeleteRenderbuffers(n int, renderbuffers *uint32) {
	C.glDeleteRenderbuffers(C.GLsizei(n), (*C.GLuint)(renderbuffers))
}
func DeleteShader(shader uint) {
	C.glDeleteShader(C.GLuint(shader))
}
func DeleteTextures(n int, textures *uint32) {
	C.glDeleteTextures(C.GLsizei(n), (*C.GLuint)(textures))
}
func DepthFunc(f uint) { C.glDepthFunc(C.GLenum(f)) }

func DepthMask(flag bool) { C.glDepthMask(glBoolean(flag)) }

func DepthRangef(zNear, zFar float32) {
	C.glDepthRangef(C.GLclampf(zNear), C.GLclampf(zFar))
}
func DetachShader(program, shader uint) {
	C.glDetachShader(C.GLuint(program), C.GLuint(shader))
}
func Disable(cap uint) { C.glDisable(C.GLenum(cap)) }

func DisableVertexAttribArray(index uint) {
	C.glDisableVertexAttribArray(C.GLuint(index))
}
func DrawArrays(mode uint, first, count int) {
	C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}
func DrawElements(mode uint, count int, type_ uint, indices Void) {
	C.glDrawElements(
		C.GLenum(mode), C.GLsizei(count),
		C.GLenum(type_), unsafe.Pointer(indices))
}

func EnableVertexAttribArray(index uint) {
	C.glEnableVertexAttribArray(C.GLuint(index))
}
func Enable(cap int) { C.glEnable(C.GLenum(cap)) }
func Finish()        { C.glFinish() }
func Flush()         { C.glFlush() }

func FramebufferRenderbuffer(target, attachment, renderbuffertarget, renderbuffer uint) {
	C.glFramebufferRenderbuffer(
		C.GLenum(target), C.GLenum(attachment),
		C.GLenum(renderbuffertarget), C.GLuint(renderbuffer))
}
func FramebufferTexture2D(target, attachment, textarget, texture uint, level int) {
	C.glFramebufferTexture2D(
		C.GLenum(target), C.GLenum(attachment),
		C.GLenum(textarget), C.GLuint(texture), C.GLint(level))
}
func FrontFace(mode uint) {
	C.glFrontFace(C.GLenum(mode))
}
func GenBuffers(n int, buffers Void) {
	C.glGenBuffers(C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}
func GenerateMipmap(target uint) {
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
func GetActiveAttrib(program, index uint, bufsize int) (size int32, type_ uint32, name string) {
	cs := CString(name)
	defer C.free(unsafe.Pointer(cs))
	C.glGetActiveAttrib(C.GLuint(program), C.GLuint(index),
		C.GLsizei(bufsize), nil, (*C.GLint)(&size), (*C.GLenum)(&type_), cs)
	name = GoString(cs)
	return

}
func GetActiveUniform(program, index uint, bufsize int) (size int32, type_ uint32, name string) {
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
func GetBooleanv(pname uint, params []bool) []bool {
	C.glGetBooleanv(C.GLenum(pname), (*C.GLboolean)(unsafe.Pointer(&params[0])))
	return params
}
func GetBufferParameteriv(target, pname uint, params []int32) []int32 {
	C.glGetBufferParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetFloatv(pname uint, params []float32) []float32 {
	C.glGetFloatv(C.GLenum(pname), (*C.GLfloat)(&params[0]))
	return params
}
func GetFramebufferAttachmentParameteriv(
	target, attachment, pname uint, params []int32) []int32 {
	C.glGetFramebufferAttachmentParameteriv(
		C.GLenum(target), C.GLenum(attachment),
		C.GLenum(pname), (*C.GLint)(&params[0]))
	return params

}
func GetIntegerv(pname uint, params []int32) []int32 {
	C.glGetIntegerv(C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetProgramiv(program, pname uint, params []int32) []int32 {
	C.glGetProgramiv(C.GLuint(program), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetProgramInfoLog(program uint, bufsize int) string {
	cs := CString("")
	defer C.free(unsafe.Pointer(cs))
	C.glGetProgramInfoLog(C.GLuint(program), C.GLsizei(bufsize), nil, cs)
	return GoString(cs)

}
func GetRenderbufferParameteriv(target, pname uint, params []int32) []int32 {
	C.glGetRenderbufferParameteriv(
		C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetShaderiv(shader, pname uint, params []int32) []int32 {
	C.glGetShaderiv(C.GLuint(shader), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetShaderInfoLog(shader uint, bufsize int) string {
	cs := CString("")
	defer C.free(unsafe.Pointer(cs))
	C.glGetShaderInfoLog(C.GLuint(shader), C.GLsizei(bufsize), nil, cs)
	return GoString(cs)
}
func GetShaderPrecisionFormat(shadertype, precisiontype uint) (
	range_ int32, precision int32) {
	C.glGetShaderPrecisionFormat(
		C.GLenum(shadertype), C.GLenum(precisiontype),
		(*C.GLint)(&range_), (*C.GLint)(&precision))
	return
}
func GetShaderSource(shader uint, bufsize int) string {
	cs := CString("")
	defer C.free(unsafe.Pointer(cs))
	C.glGetShaderSource(C.GLuint(shader), C.GLsizei(bufsize), nil, cs)
	return GoString(cs)
}
func GetTexParameterfv(target, pname uint, params []float32) []float32 {
	C.glGetTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
	return params
}
func GetTexParameteriv(target, pname uint, params []int32) []int32 {
	C.glGetTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetUniformfv(program uint, location int, params []float32) []float32 {
	C.glGetUniformfv(C.GLuint(program), C.GLint(location), (*C.GLfloat)(&params[0]))
	return params
}
func GetUniformiv(program uint, location int, params []int32) []int32 {
	C.glGetUniformiv(C.GLuint(program), C.GLint(location), (*C.GLint)(&params[0]))
	return params
}
func GetVertexAttribfv(index, pname uint, params []float32) []float32 {
	C.glGetVertexAttribfv(C.GLuint(index), C.GLenum(pname), (*C.GLfloat)(&params[0]))
	return params
}
func GetVertexAttribiv(index, pname uint, params []int32) []int32 {
	C.glGetVertexAttribiv(C.GLuint(index), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func GetVertexAttribPointerv(index, pname uint, pointer []Void) []Void {
	C.glGetVertexAttribPointerv(C.GLuint(index), C.GLenum(pname), (*unsafe.Pointer)(&pointer[0]))
	return pointer
}
func Hint(target, mode uint) {
	C.glHint(C.GLenum(target), C.GLenum(mode))
}
func LineWidth(width float32) {
	C.glLineWidth(C.GLfloat(width))
}
func LinkProgram(program uint) {
	C.glLinkProgram(C.GLuint(program))
}
func PixelStorei(pname, param uint) {
	C.glPixelStorei(C.GLenum(pname), C.GLint(param))
}
func PolygonOffset(factor float32, units float32) {
	C.glPolygonOffset(C.GLfloat(factor), C.GLfloat(units))
}
func ReadPixels(x, y, width, height int, format, type_ uint, pixels Void) {
	C.glReadPixels(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height),
		C.GLenum(format), C.GLenum(type_), unsafe.Pointer(pixels))
}
func ReleaseShaderCompiler() {
	C.glReleaseShaderCompiler()
}
func RenderbufferStorage(target, internalformat uint, width, height int) {
	C.glRenderbufferStorage(C.GLenum(target), C.GLenum(internalformat),
		C.GLsizei(width), C.GLsizei(height))
}
func SampleCoverage(value float32, invert bool) {
	C.glSampleCoverage(C.GLclampf(value), glBoolean(invert))
}
func Scissor(x, y, width, height int) {
	C.glScissor(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
func ShaderBinary(n int, shaders Void, binaryformat uint, binary Void, length int) {
	C.glShaderBinary(C.GLsizei(n), (*C.GLuint)(shaders),
		C.GLenum(binaryformat), unsafe.Pointer(binary), C.GLsizei(length))
}
func ShaderSource(shader uint, s string) {
	cs := CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.glShaderSource(C.GLuint(shader), C.GLsizei(1), &cs, nil)
}
func StencilFunc(func_ uint, ref int, mask uint) {
	C.glStencilFunc(C.GLenum(func_), C.GLint(ref), C.GLuint(mask))
}
func StencilFuncSeparate(face, func_ uint, ref int, mask uint) {
	C.glStencilFuncSeparate(
		C.GLenum(face), C.GLenum(func_), C.GLint(ref), C.GLuint(mask))
}
func StencilMask(mask uint) {
	C.glStencilMask(C.GLuint(mask))
}
func StencilMaskSeparate(face, mask uint) {
	C.glStencilMaskSeparate(C.GLenum(face), C.GLuint(mask))
}
func StencilOp(fail, zfail, zpass uint) {
	C.glStencilOp(C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}
func StencilOpSeparate(face, fail, zfail, zpass uint) {
	C.glStencilOpSeparate(
		C.GLenum(face), C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}
func TexImage2D(target uint, level, internalformat, width, height,
	border int, format, type_ uint, pixels Void) {
	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GLint(internalformat),
		C.GLsizei(width), C.GLsizei(height), C.GLint(border),
		C.GLenum(format), C.GLenum(type_), unsafe.Pointer(pixels))
}
func TexParameterf(target, pname uint, param float32) {
	C.glTexParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}
func TexParameterfv(target, pname uint, params []float32) []float32 {
	C.glTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
	return params
}
func TexParameteri(target, pname uint, param int) {
	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}
func TexParameteriv(target, pname uint, params []int32) []int32 {
	C.glTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
	return params
}
func TexSubImage2D(target uint, level, xoffset, yoffset, width, height int,
	format, type_ uint, pixels Void) {

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
func UseProgram(program uint) {
	C.glUseProgram(C.GLuint(program))
}
func ValidateProgram(program uint) {
	C.glValidateProgram(C.GLuint(program))
}
func VertexAttrib1f(indx uint, x float32) {
	C.glVertexAttrib1f(C.GLuint(indx), C.GLfloat(x))
}
func VertexAttrib1fv(indx uint, values []float32) {
	C.glVertexAttrib1fv(C.GLuint(indx), (*C.GLfloat)(&values[0]))
}
func VertexAttrib2f(indx uint, x, y float32) {
	C.glVertexAttrib2f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y))
}
func VertexAttrib2fv(indx uint, values []float32) {
	C.glVertexAttrib2fv(C.GLuint(indx), (*C.GLfloat)(&values[0]))
}
func VertexAttrib3f(indx uint, x, y, z float32) {
	C.glVertexAttrib3f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}
func VertexAttrib3fv(indx uint, values []float32) {
	C.glVertexAttrib3fv(C.GLuint(indx), (*C.GLfloat)(&values[0]))
}
func VertexAttrib4f(indx uint, x, y, z, w float32) {
	C.glVertexAttrib4f(C.GLuint(indx),
		C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}
func VertexAttrib4fv(indx uint, values []float32) {
	C.glVertexAttrib4fv(C.GLuint(indx), (*C.GLfloat)(&values[0]))
}
func VertexAttribPointer(indx uint, size int, type_ uint, normalized bool, stride int, ptr Void) {
	C.glVertexAttribPointer(C.GLuint(indx), C.GLint(size), C.GLenum(type_),
		glBoolean(normalized), C.GLsizei(stride), unsafe.Pointer(ptr))
}
func Viewport(x int, y int, width int, height int) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
func IsBuffer(buffer uint) bool {
	return *goBoolean(C.glIsBuffer(C.GLuint(buffer)))
}
func IsEnabled(cap uint) bool {
	return *goBoolean(C.glIsEnabled(C.GLenum(cap)))
}
func IsFramebuffer(framebuffer uint) bool {
	return *goBoolean(C.glIsFramebuffer(C.GLuint(framebuffer)))
}
func IsProgram(program uint) bool {
	return *goBoolean(C.glIsProgram(C.GLuint(program)))
}
func IsRenderbuffer(renderbuffer uint) bool {
	return *goBoolean(C.glIsRenderbuffer(C.GLuint(renderbuffer)))
}
func IsShader(shader uint) bool {
	return *goBoolean(C.glIsShader(C.GLuint(shader)))
}
func IsTexture(texture uint) bool {
	return *goBoolean(C.glIsTexture(C.GLuint(texture)))
}
func GetError() int {
	return int(C.glGetError())
}
func CheckFramebufferStatus(target uint) uint {
	return uint(C.glCheckFramebufferStatus(C.GLenum(target)))
}
func CreateProgram() uint {
	return uint(C.glCreateProgram())
}
func CreateShader(type_ uint) uint {
	return uint(C.glCreateShader(C.GLenum(type_)))
}
func GetAttribLocation(program uint, name string) int {
	s := CString(name)
	defer C.free(unsafe.Pointer(s))
	return int(C.glGetAttribLocation(C.GLuint(program), s))
}
func GetUniformLocation(program uint, name string) int {
	s := CString(name)
	defer C.free(unsafe.Pointer(s))
	return int(C.glGetUniformLocation(C.GLuint(program), s))
}
func GetString(name uint) string {
	s := (*C.GLchar)(unsafe.Pointer(C.glGetString(C.GLenum(name))))
	defer C.free(unsafe.Pointer(s))
	return GoString(s)

}
