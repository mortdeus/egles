package x

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"
	"log"
	"os"
)

var (
	x   *xgbutil.XUtil
	err error
)

type Window struct {
	*xwindow.Window
}

func (this Window) Init() *Window {

	if x == nil {
		if x, err = xgbutil.NewConn(); err != nil {
			log.Fatal(err)
		}
		mousebind.Initialize(x)
	}

	this.Window, err = xwindow.Create(x, 0)
	if err != nil {
		log.Fatal(err)
	}
	this.WMGracefulClose(func(w *xwindow.Window) {
		xevent.Detach(w.X, w.Id)
		mousebind.Detach(w.X, w.Id)
		w.Destroy()

		
	})
	this.Map()

	return &this
}
