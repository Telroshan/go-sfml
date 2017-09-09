package main

import (
	"runtime"

	"github.com/teh-cmc/go-sfml/graphics"
	"github.com/teh-cmc/go-sfml/window"
)

func init() { runtime.LockOSThread() }

func main() {
	vm := window.NewSfVideoMode()
	vm.SetWidth(800)
	vm.SetHeight(600)
	vm.SetBitsPerPixel(32)
	cs := window.NewSfContextSettings()
	cs.SetDepthBits(32)
	w := window.SfWindow_create(
		vm, "My window",
		uint(window.SfResize|window.SfClose),
		window.NewSfContextSettings())

openLoop:
	for window.SfWindow_isOpen(w) > 0 {
		e := window.NewSfEvent()
		for window.SfWindow_pollEvent(w, e) > 0 {
			if e.GetXtype() == window.SfEventType(window.SfEvtClosed) {
				break openLoop
			}
		}
		graphics.SfRenderWindow_clear(w, graphics.GetSfBlack())
		graphics.SfRenderWindow_display(w)
	}

	window.SfWindow_destroy(w)
}
