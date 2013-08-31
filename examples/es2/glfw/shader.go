package main

import gl "github.com/mortdeus/egles/es2"
import "log"

const (
	fsh = `
	varying mediump vec4 v_color;

	void main() {
		gl_FragColor = v_color;
	}
`
	vsh = `
        attribute vec4 pos;
        attribute vec4 color;
        varying mediump vec4 v_color;

        void main() {
          	gl_Position = pos;
          	v_color = color;
        }
`
)

func FragmentShader(s string) uint32 {
	shader := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(shader, 1, &s, nil)
	gl.CompileShader(shader)
	var stat int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &stat)
	if stat == 0 {
		log.Printf("Fragment Shader compilation failed.\n")
	}
	return shader
}

func VertexShader(s string) uint32 {
	shader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(shader, 1, &s, nil)
	gl.CompileShader(shader)
	var stat int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &stat)
	if stat == 0 {
		log.Printf("Vertex Shader compilation failed. \n")
	}
	return shader
}
func Program(fsh, vsh uint32) uint32 {
	p := gl.CreateProgram()
	gl.AttachShader(p, fsh)
	gl.AttachShader(p, vsh)
	gl.LinkProgram(p)
	var stat int32
	gl.GetProgramiv(p, gl.LINK_STATUS, &stat)
	if stat == 0 {
		var s = make([]byte, 1000)
		var length gl.Sizei
		_log := string(s)
		gl.GetProgramInfoLog(p, 1000, &length, &_log)
		log.Printf("Error: linking:\n%s\n", _log)
	}
	return p
}
