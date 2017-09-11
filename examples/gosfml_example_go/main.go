package main

import (
	"runtime"

	"github.com/teh-cmc/go-sfml/graphics"
	"github.com/teh-cmc/go-sfml/window"
)

func init() { runtime.LockOSThread() }

func main() {
	vm := window.NewSfVideoMode()
	defer window.DeleteSfVideoMode(vm)
	vm.SetWidth(800)
	vm.SetHeight(600)
	vm.SetBitsPerPixel(32)

	/* Create the main window */
	cs := window.NewSfContextSettings()
	defer window.DeleteSfContextSettings(cs)
	w := graphics.SfRenderWindow_create(vm, "SFML window", uint(window.SfResize|window.SfClose), cs)
	defer window.SfWindow_destroy(w)

	ev := window.NewSfEvent()
	defer window.DeleteSfEvent(ev)

	/* Start the game loop */
	for window.SfWindow_isOpen(w) > 0 {
		/* Process events */
		for window.SfWindow_pollEvent(w, ev) > 0 {
			/* Close window: exit */
			if ev.GetXtype() == window.SfEventType(window.SfEvtClosed) {
				return
			}
		}
		graphics.SfRenderWindow_clear(w, graphics.GetSfRed())
		graphics.SfRenderWindow_display(w)
	}
}
