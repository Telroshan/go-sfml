# go-sfml

Up-to-date Go bindings for [SFML](http://www.sfml-dev.org), the Simple and Fast Multimedia Library.

These bindings *are entirely generated* using [SWIG](http://www.swig.org/) and the official [C bindings of SFML](https://www.sfml-dev.org/download/csfml/).  
Hence they should be fairly easy to maintain & keep in sync with future releases from upstream.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

**Table of Contents**
- [go-sfml](#go-sfml)
  - [Portability](#portability)
  - [Requirements](#requirements)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Modules](#modules)
    - [Example](#example)
  - [License](#license-)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

---

## Portability

I have only tested these on darwin/amd64, using the latest macOS Sierra.  
Feel free to open issues and/or PRs if you're running into problems on other platforms.

## Requirements

- [SFML 2.4.0](https://www.sfml-dev.org/download/sfml/2.4.0)
- [CSFML 2.4](https://www.sfml-dev.org/download/csfml/)

## Installation

Assuming both libraries and their header files are correctly installed and available through the standard paths of your system:
```
$ go get -u gopkg.in/teh-cmc/go-sfml.v24/...
```

That's it!

## Usage

The generated APIs very closely follow those of [SFML's C bindings](https://www.sfml-dev.org/download/csfml/): the [tutorials](http://www.sfml-dev.org/tutorials/) & [documentation](http://www.sfml-dev.org/documentation/) available for the official C++ implementation, as well as an editor with a well configured Go-autocompletion, will get you a long way.

### Modules

The original C & C++ implementations of SFML come with 5 modules: [*Audio*](https://www.sfml-dev.org/documentation/2.4.0/group__audio.php), [*Graphics*](https://www.sfml-dev.org/documentation/2.4.0/group__graphics.php), [*Network*](https://www.sfml-dev.org/documentation/2.4.0/group__network.php), [*System*](https://www.sfml-dev.org/documentation/2.4.0/group__system.php) and [*Window*](https://www.sfml-dev.org/documentation/2.4.0/group__window.php).

Of these 5 modules:
- ***Audio***, ***Graphics*** & ***Window*** come with complete Go packages counterparts
- ***System*** also has a dedicated Go package, but only contains the `sfVector2` & `sfVector3` classes; everything else is already available in one form or another in Go's standard library
- ***Network*** has been entirely discard, use Go's standard library instead

### Example

Here's a straightforward example of creating a window and handling events:
```Go
package main

import (
	"runtime"

	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
	"gopkg.in/teh-cmc/go-sfml.v24/window"
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
```

For comparison's sake, the exact same thing in C:
```C
#include <SFML/Window.h>
#include <SFML/Graphics.h>

int main() {
    sfVideoMode mode = {800, 600, 32};
    sfRenderWindow* window;
    sfEvent event;

    /* Create the main window */
    window = sfRenderWindow_create(mode, "SFML window", sfResize | sfClose, NULL);
    if (!window)
        return 1;

    /* Start the game loop */
    while (sfRenderWindow_isOpen(window)) {
        /* Process events */
        while (sfRenderWindow_pollEvent(window, &event)) {
            /* Close window : exit */
            if (event.type == sfEvtClosed)
                sfRenderWindow_close(window);
        }
        /* Clear the screen */
        sfRenderWindow_clear(window, sfRed);
        sfRenderWindow_display(window);
    }

    /* Cleanup resources */
    sfRenderWindow_destroy(window);
}
```

## License ![License](https://img.shields.io/badge/license-Zlib-blue.svg?style=plastic)

The Zlib/libpng - see [LICENSE](./LICENSE) for more details.

Copyright (c) 2017	Clement 'cmc' Rey	<cr.rey.clement@gmail.com> [@teh_cmc](https://twitter.com/teh_cmc)
