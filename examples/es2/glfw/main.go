package main

import (
	"fmt"
	glfw "github.com/go-gl/glfw3"
	gl "github.com/mortdeus/egles/es2"
)

const (
	Title  = "egles"
	Width  = 640
	Height = 480
)

var (
	vertices = [12]float32{
		-0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, 0.0, 1.0,
		0.0, 0.5, 0.0, 1.0,
	}
	colors = [12]float32{
		1.0, 0.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
		0.0, 0.0, 1.0, 1.0,
	}
	verticesArrayBuffer, colorsArrayBuffer uint32
	attrPos, attrColor                     uint32
)

func errorCallback(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}
func main() {
	glfw.SetErrorCallback(errorCallback)
	if !glfw.Init() {
		println("glfw init failure")
	}
	defer glfw.Terminate()
	glfw.WindowHint(glfw.ClientApi, glfw.OpenglEsApi)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)
	window, err := glfw.CreateWindow(Width, Height, Title, nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	initScene()
	reshape(window.GetSize())
	for !window.ShouldClose() {
		drawScene()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
func initScene() {
	program := Program(FragmentShader(fsh), VertexShader(vsh))
	gl.UseProgram(program)
	attrPos = gl.GetAttribLocation(program, "pos")
	attrColor = gl.GetAttribLocation(program, "color")
	gl.GenBuffers(1, &verticesArrayBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, verticesArrayBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.SizeiPtr(len(vertices))*4, gl.Void(&vertices[0]), gl.STATIC_DRAW)
	gl.GenBuffers(1, &colorsArrayBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorsArrayBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.SizeiPtr(len(colors))*4, gl.Void(&colors[0]), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(attrPos)
	gl.EnableVertexAttribArray(attrColor)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

}
func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.BindBuffer(gl.ARRAY_BUFFER, verticesArrayBuffer)
	gl.VertexAttribPointer(attrPos, 4, gl.FLOAT, false, 0, 0)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorsArrayBuffer)
	gl.VertexAttribPointer(attrColor, 4, gl.FLOAT, false, 0, 0)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.Flush()
	gl.Finish()

}
func reshape(width, height int) {
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))

}
