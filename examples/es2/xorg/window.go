package main

import (
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"
	"log"
)

func newWindow(X *xgbutil.XUtil) {
	win, err := xwindow.Generate(X)
	if err != nil {
		log.Fatal(err)
	}
	win.Create(X.RootWin(), 0, 0, 200, 200,
		xproto.CwBackPixel|xproto.CwEventMask,
		0, xproto.EventMaskButtonRelease)
	win.WMGracefulClose(
		func(w *xwindow.Window) {
			xevent.Detach(w.X, w.Id)
			mousebind.Detach(w.X, w.Id)
			w.Destroy()
			Done <- true
		})

	win.Map()

	err = mousebind.ButtonReleaseFun(
		func(X *xgbutil.XUtil, ev xevent.ButtonReleaseEvent) {
			newWindow(X)
		}).Connect(X, win.Id, "1", false, false)

	if err != nil {
		log.Fatal(err)
	}
}
