// +build raspberry

package main

import (
	"github.com/mortdeus/egles/egl"
	"github.com/mortdeus/egles/egl/platform/raspberry"
)

const (
	INITIAL_WINDOW_WIDTH  = 1920
	INITIAL_WINDOW_HEIGHT = 1080
)

func initialize() {
	egl.BCMHostInit()
	raspberry.Initialize(raspberry.DefaultConfigAttributes, raspberry.DefaultContextAttributes)
}
