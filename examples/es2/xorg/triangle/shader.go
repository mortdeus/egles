package main

import gl "github.com/mortdeus/egles/es2"
import "fmt"

var fsh = "" +
	"varying mediump vec4 v_color;" +
	"void main(){" +
	"gl_FragColor = v_color;" +
	"}"

var vsh = "" +
	"attribute vec4 pos;" +
	"attribute vec4 color;" +
	"varying mediump vec4 v_color;" +
	"void main() {" +
	"gl_Position = pos;" +
	"v_color = color;" +
	"}"

func FragmentShader(s string) uint {
	shader := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(shader, s)
	gl.CompileShader(shader)

	if stat := gl.GetShaderiv(shader, gl.COMPILE_STATUS, make([]int32, 1)); stat[0] == 0 {
		fmt.Printf("FSH:\n%s\n", gl.GetShaderInfoLog(shader, 1000))

	}
	return shader
}
func VertexShader(s string) uint {
	shader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(shader, s)
	gl.CompileShader(shader)

	if stat := gl.GetShaderiv(shader, gl.COMPILE_STATUS, make([]int32, 1)); stat[0] == 0 {
		fmt.Printf("VSH:\n%s\n", gl.GetShaderInfoLog(shader, 1000))
	}
	return shader
}
func Program(fsh, vsh uint) uint {
	p := gl.CreateProgram()
	gl.AttachShader(p, fsh)
	gl.AttachShader(p, vsh)
	gl.LinkProgram(p)
	if stat := gl.GetProgramiv(p, gl.LINK_STATUS, make([]int32, 1)); stat[0] == 0 {
		fmt.Printf("PROG:\n%s\n", gl.GetProgramInfoLog(p, 1000))
	}
	return p
}
