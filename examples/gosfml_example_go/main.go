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

	/* Create the main window */
	var ev = window.NewSfEvent()
	w := window.SfWindow_create(
		vm, "SFML window",
		uint(window.SfResize|window.SfClose),
		window.NewSfContextSettings())
	if w == nil {
		panic("failed to create window")
	}

	/* Cleanup resources (on exit) */
	defer window.SfWindow_destroy(w)
	defer window.Swig_free(ev.Swigcptr())

	/* Start the game loop */
	for window.SfWindow_isOpen(w) > 0 {
		/* Process events */
		for window.SfWindow_pollEvent(w, ev) > 0 {
			/* Close window: exit */
			if ev.GetXtype() == window.SfEventType(window.SfEvtClosed) {
				return
			}
		}
		graphics.SfRenderWindow_clear(w, graphics.GetSfBlack())
		graphics.SfRenderWindow_display(w)
	}
}
