package main

import (
	"fmt"
	gl "github.com/mortdeus/egles/gles2"
	"github.com/tapir/glfw3-go"
)

const (
	Title  = "egles"
	Width  = 640
	Height = 480
)

var (
	redraw = true
)

func errorCallback(err int, desc string) {
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

	gl.ClearColor(.4, .4, .4, 0)
	p := Program(FragmentShader(fsh), VertexShader(vsh))
	gl.UseProgram(p)
}
func drawScene() {
	if redraw {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		redraw = false
	}
}
func reshape(width, height int) {
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
}
