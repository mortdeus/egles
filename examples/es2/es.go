package main

import gl "github.com/mortdeus/egles/egl/gles/2"
import "github.com/mortdeus/mathgl"
import "math"

func reshape(width, height int) {
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
}

func makeZRotMatrix(angle float32, m *mathgl.Mat4) {
	c := float32(math.Cos(float64(angle) * math.Pi / 180.0))
	s := float32(math.Sin(float64(angle) * math.Pi / 180.0))
	m.Identity()
	m[0] = c
	m[1] = s
	m[4] = -s
	m[5] = c
}
func makeScaleMatrix(xs, ys, zs float32, m *mathgl.Mat4) {
	m[0] = xs
	m[5] = ys
	m[10] = zs
	m[15] = 1.0
}
