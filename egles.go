package egles

import "github.com/mortdeus/egles/xorg"

type Window struct {
}
func (this Window) Init() window {
	return x.Window.Init()
}
