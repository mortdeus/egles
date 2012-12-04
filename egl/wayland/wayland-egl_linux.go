package wayland

import (
	"github.com/mortdeus/gowl"
	"unsafe"
)

type Window struct {
	Surface *gowl.Surface
	width, height,
	dx, dy int32
	AttachedWidth,
	AttachedHeight int32
	private        unsafe.Pointer
	ResizeCallback func(*Window, unsafe.Pointer)
}

func (this *Window) Resize(width, height, dx, dy int32) {
	this.width = width
	this.height = height
	this.dx = dx
	this.dy = dy

	if this.ResizeCallback != nil {
		this.ResizeCallback(this, this.private)
	}
}
func CreateWindow(surface *gowl.Surface, width, height int32) *Window {
	window := new(Window)
	window.Surface = surface
	window.Resize(width, height, 0, 0)
	return window
}
func (window *Window) Destroy() {
	window = nil
}
func (window *Window) GetAttachedSize() (int32, int32) {
	return window.AttachedWidth, window.AttachedHeight
}
