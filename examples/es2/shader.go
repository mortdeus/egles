package main

import gl "github.com/mortdeus/egles/gles/2"

var fragShader = func() gl.Uint {
	s := `
	varying vec4 v_color;
	
	void main() {
		gl_FragColor = v_color;
	}
`
	shader := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(shader, 1, &s, nil)
	gl.CompileShader(shader)
	var stat *gl.Int
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, stat)
	if stat == nil {

		panic("Fragment Shader did not compile")
	}
	return shader

}
var vertShader = func() gl.Uint {
	s := `
    uniform mat4 modelviewProjection;
    attribute vec4 pos;
    attribute vec4 color;
    varying vec4 v_color;
    
    void main() {
       gl_Position = modelviewProjection * pos;
       v_color = color;
    }
`
	shader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(shader, 1, &s, nil)
	gl.CompileShader(shader)
	var stat *gl.Int
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, stat)
	if stat == nil {
		panic("Vertex shader did not compile")
	}
	return shader
}

var program = func() gl.Uint {
	p := gl.CreateProgram()
	gl.AttachShader(p, fragShader())
	gl.AttachShader(p, vertShader())
	gl.LinkProgram(p)
	return p
}
