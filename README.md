# go-sfml

Go bindings for [SFML](http://www.sfml-dev.org), the Simple and Fast Multimedia Library, version [2.5.1](https://www.sfml-dev.org/changelog.php#sfml-2.5.1)

Originally made by [teh-cmc](https://github.com/teh-cmc)

These bindings *are entirely generated* using [SWIG](http://www.swig.org/) and the official [C bindings of SFML](https://www.sfml-dev.org/download/csfml/).  
Hence they should be fairly easy to maintain & keep in sync with future releases from upstream

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Portability](#portability)
- [Usage](#usage)
  - [Installation](#installation)
    - [First setup](#first-setup)
    - [Get the module](#get-the-module)
    - [Troubleshooting](#troubleshooting)
      - [Error: `SFML/Window.h: No such file or directory`](#error-sfmlwindowh-no-such-file-or-directory)
      - [Error: `imports github.com/telroshan/go-sfml/v2/window: build constraints exclude all Go files in [...]`](#error-imports-githubcomtelroshango-sfmlv2window-build-constraints-exclude-all-go-files-in-)
      - [Error: `csfml-window-2.dll could not be found`](#error-csfml-window-2dll-could-not-be-found)
  - [Bat script](#bat-script)
  - [API](#api)
  - [Modules](#modules)
  - [Examples](#examples)
    - [Basic example](#basic-example)
    - [Other examples](#other-examples)
- [Building go-sfml](#building-go-sfml)
  - [Download & compile SFML + CSFML](#download--compile-sfml--csfml)
    - [Troubleshooting](#troubleshooting-1)
      - [Error: `/usr/bin/env: ‘bash\r’: No such file or directory`](#error-usrbinenv-bashr-no-such-file-or-directory)
      - [Error: `CMake Error [...] Could not find X11` _(or a similar error with just another name instead of `X11`)_](#error-cmake-error--could-not-find-x11-or-a-similar-error-with-just-another-name-instead-of-x11)
  - [Setup swig](#setup-swig)
    - [Option 1 - Install it from your package manager](#option-1---install-it-from-your-package-manager)
    - [Option 2 - Build it locally](#option-2---build-it-locally)
    - [Troubleshooting](#troubleshooting-2)
      - [Error: `/usr/bin/env: ‘bash\r’: No such file or directory`](#error-usrbinenv-bashr-no-such-file-or-directory-1)
      - [Error: `Cannot find pcre-config script from PCRE (Perl Compatible Regular Expressions)`](#error-cannot-find-pcre-config-script-from-pcre-perl-compatible-regular-expressions)
  - [Build go bindings](#build-go-bindings)
    - [Troubleshooting](#troubleshooting-3)
      - [Error: `/usr/bin/env: ‘bash\r’: No such file or directory`](#error-usrbinenv-bashr-no-such-file-or-directory-2)
      - [Error: `swig: command not found`](#error-swig-command-not-found)
      - [Error: `Unable to find 'swig.swg'`](#error-unable-to-find-swigswg)
      - [Error: `patchelf: command not found`](#error-patchelf-command-not-found)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

---

## Portability

I have only tested these on Windows 10

Feel free to open issues and/or PRs if you're running into problems on other platforms

## Usage

### Installation
#### First setup
I'll cover Windows in this section _(I build go-sfml using [WSL](https://docs.microsoft.com/en-us/windows/wsl/install), but build my go apps on Windows)_

1. Download [CSFML 2.5.1](https://www.sfml-dev.org/download/csfml/) and extract it wherever you like _(assuming `C:\CSFML_2.5.1` for the next steps)_. I'm downloading the 64 bits version since I'm on a 64 bits Windows
2. Download and install the [GCC compiler](https://gcc.gnu.org/install/binaries.html) _(in my case, the [mingw x86_64-win32-sjlj](https://github.com/niXman/mingw-builds-binaries/releases) one)_
3. We now need to define the `CGO_CFLAGS` and `CGO_LDFLAGS` environment variables so the linker knows where the CSFML headers and compiled libraries are at

Assuming CSFML is extracted at `C:\CSFML_2.5.1`, we can run the following in a command line :
```
set CGO_CFLAGS="-IC:\CSFML_2.5.1\include"
set CGO_LDFLAGS="-LC:\CSFML_2.5.1\lib\gcc"
```
Feel free to set those variables system-wide so you don't have to define them on the fly everytime

Notice the **-I** and **-L** before the paths, with no whitespace in between

4. Build the example to ensure your setup is working
```bash
cd examples/basic_window
go get && go build
```
5. Copy the CSFML DLLs next to your executable, for the basic example we only need `csfml-window-2.dll` and `csfml-graphics-2.dll` that you'll find into `C:\CSFML_2.5.1\bin`
6. Run the exe, a red window should appear!
#### Get the module
For future projects, simply run
```bash
go get github.com/telroshan/go-sfml/v2
```
And import the modules you need
```Go
import (
  "github.com/telroshan/go-sfml/v2/graphics"
  "github.com/telroshan/go-sfml/v2/window"
  "github.com/telroshan/go-sfml/v2/system"
  "github.com/telroshan/go-sfml/v2/audio"
)
```
#### Troubleshooting
##### Error: `SFML/Window.h: No such file or directory`
&#8594; Make sure you have downloaded the CSFML version that matches your system _(don't download the 32 bits one if you're on a 64 bits Windows)_

&#8594; Make sure you included the subfolder matching your compiler _(gcc here)_ for `CGO_LDFLAGS`, and didn't make it point to the `lib` root folder itself

&#8594; Make sure you defined the environment variables in the same command line that you ran `go build` into _(the variables defined with the [`set`](https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/set_1) command won't exist outside of the command line they've been set into)_

&#8594; Make sure you typed the variable names correctly, i.e `CGO_CFLAGS` and `CGO_LDFLAGS`

&#8594; Make sure you didn't invert the paths between the two variables. `CGO_CFLAGS` should point to the `include` folder whereas `CGO_LDFLAGS` should point to the libs

&#8594; Make sure you made no typo with the syntax of -I and -L for those variables
##### Error: `imports github.com/telroshan/go-sfml/v2/window: build constraints exclude all Go files in [...]`
&#8594; You need to define the environment variable [`CGO_ENABLED`](https://pkg.go.dev/cmd/cgo). As the doc says:
> The cgo tool is enabled by default for native builds on systems where it is expected to work. It is disabled by default when cross-compiling. You can control this by setting the CGO_ENABLED environment variable when running the go tool: set it to 1 to enable the use of cgo, and to 0 to disable it.
```
set CGO_ENABLED=1
```
##### Error: `csfml-window-2.dll could not be found`
&#8594; You probably didn't copy CSFML DLLs next to your executable, as mentioned in step 5

### Bat script
Alternatively to steps 3 to 5 from the [previous section](#installation), you could use a bat script to automate that process for you, such as [the one in the basic window example](https://github.com/Telroshan/go-sfml/blob/master/examples/basic_window/build.bat)
```bat
@ECHO OFF

rem This script sets the environment variables to be able to build the app, and copies the CSFML DLLs over if there aren't any in the folder

rem Edit the CSFML_PATH variable to match the path of your CSFML installation
set CSFML_PATH=C:\CSFML_2.5.1
rem Edit the COMPILER_NAME variable if you're not using gcc
set COMPILER_NAME=gcc

set CGO_CFLAGS="-I%CSFML_PATH%\include"
set CGO_LDFLAGS="-L%CSFML_PATH%\lib\%COMPILER_NAME%"

go get
if %ERRORLEVEL% NEQ 0 (echo go get failed && exit /b %ERRORLEVEL%)

go build
if %ERRORLEVEL% NEQ 0 (echo go build failed && exit /b %ERRORLEVEL%)

echo Build complete

if not exist "%~dp0*.dll" (
  echo No DLLs in folder, getting them from CSFML folder
  xcopy /s "%CSFML_PATH%\bin" "%~dp0"
  if %ERRORLEVEL% NEQ 0 (echo failed to copy DLLs && exit /b %ERRORLEVEL%)
)
```

### API
The generated APIs very closely follow those of [SFML's C bindings](https://www.sfml-dev.org/download/csfml/): the [tutorials](http://www.sfml-dev.org/tutorials/) & [documentation](http://www.sfml-dev.org/documentation/) available for the official C++ implementation, as well as an editor with a well configured Go-autocompletion, will get you a long way

### Modules

The original C & C++ implementations of SFML come with 5 modules: [*Audio*](https://www.sfml-dev.org/documentation/2.4.0/group__audio.php), [*Graphics*](https://www.sfml-dev.org/documentation/2.4.0/group__graphics.php), [*Network*](https://www.sfml-dev.org/documentation/2.4.0/group__network.php), [*System*](https://www.sfml-dev.org/documentation/2.4.0/group__system.php) and [*Window*](https://www.sfml-dev.org/documentation/2.4.0/group__window.php)

Of these 5 modules:
- ***Audio***, ***Graphics*** & ***Window*** come with complete Go packages counterparts
- ***System*** also has a dedicated Go package, but only contains the `sfVector2` & `sfVector3` classes; everything else is already available in one form or another in Go's standard library
- ***Network*** has been entirely discard, use Go's standard library instead

### Examples
#### Basic example
Here's a straightforward example of creating a window and handling events:
```Go
package main

import (
	"runtime"

	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
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
			if ev.GetEvType() == window.SfEventType(window.SfEvtClosed) {
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

#### Other examples
You'll find other examples in the [examples folder](https://github.com/Telroshan/go-sfml/tree/master/examples)
1. [Basic window](https://github.com/Telroshan/go-sfml/tree/master/examples/basic_window) : just the same as above
2. [Tennis](https://github.com/Telroshan/go-sfml/tree/master/examples/tennis) : go version of [SFML's tennis example](https://github.com/SFML/SFML/tree/master/examples/tennis)

Feel free to open PRs if you have any example you'd like to share!

## Building go-sfml
If you just want to use go-sfml for SFML 2.5.1 into your go project, you may want to read the [Usage section](#usage) instead

If you want to build your own bindings for a different version of SFML, then this section is for you!

_Note_: the following steps were realized in Windows' [Ubuntu bash](https://docs.microsoft.com/en-us/windows/wsl/install). Feel free to open issues and/or PRs if you're running into problems on other Unix-based platforms

### Download & compile SFML + CSFML
1. Install [SFML dependencies](https://www.sfml-dev.org/tutorials/2.5/compile-with-cmake.php#installing-dependencies) first
2. Run the `sfml.sh` script to handle the process of downloading SFML/CSFML and compiling them for you

#### Troubleshooting
##### Error: `/usr/bin/env: ‘bash\r’: No such file or directory`
&#8594; You're probably having [CRLF](https://developer.mozilla.org/en-US/docs/Glossary/CRLF) issues _(happened to me when cloning the repo on Windows initially before switching to WSL)_

&#8594; Use `dos2unix` _(install it if you don't have the command)_ on the 3 scripts : `dos2unix sfml.sh swig.sh build.sh`

##### Error: `CMake Error [...] Could not find X11` _(or a similar error with just another name instead of `X11`)_
<pre>
<b>Could not find</b> X11
Call Stack (most recent call first):
  src/SFML/Window/CMakeLists.txt (find_package)
</pre>
  &#8594; You probably didn't install every [SFML dependency](https://www.sfml-dev.org/tutorials/2.5/compile-with-cmake.php#installing-dependencies). Don't forget the development headers as well! For example on Ubuntu, you'll want to install `libx11-dev`, `xorg-dev`, `libudev-dev`, and so on

### Setup swig
#### Option 1 - Install it from your package manager
1. Depending on your platform, you may simply download the available package. For example on Ubuntu, `sudo apt install swig`
#### Option 2 - Build it locally
1. Run `sudo apt install libpcre3-dev` _(swig requires this package)_
2. Run the `swig.sh` script
3. Check where swig thinks its lib is, by running `./swig/swig -swiglib`. It should output <code><i>${path/to/go-sfml}</i>/swig/Lib</code>
4. If the output doesn't match, fix that by overriding the `SWIG_LIB` environment variable. You may run <code>export SWIG_LIB=<i>${path/to/go-sfml}</i>/swig/Lib</code> to override the var just for your current session, or [make it persistent](https://help.ubuntu.com/community/EnvironmentVariables#Persistent_environment_variables)
5. Run `./swig/swig -swiglib` again to ensure swig has the correct path to its own lib
6. Update the `build.sh` script and change the [line 23](https://github.com/telroshan/go-sfml/blob/master/build.sh#L23) : the script is looking for a [global command](#option-1---install-it-from-your-package-manager) `swig`, you must replace that path by `./swig/swig` to use the local build instead
#### Troubleshooting
##### Error: `/usr/bin/env: ‘bash\r’: No such file or directory`
&#8594; You're probably having [CRLF](https://developer.mozilla.org/en-US/docs/Glossary/CRLF) issues _(happened to me when cloning the repo on Windows initially before switching to WSL)_

&#8594; Use `dos2unix` _(install it if you don't have the command)_ on the 3 scripts : `dos2unix sfml.sh swig.sh build.sh`
##### Error: `Cannot find pcre-config script from PCRE (Perl Compatible Regular Expressions)`
&#8594; You probably didn't install libpcre3-dev as mentioned in step 1. Run `sudo apt install libpcre3-dev` and try again

### Build go bindings
1. Run `sudo apt install patchelf` _(the build script uses this package to fix the missing links from the built CSFML libraries to the SFML ones)_
2. Run `build.sh`
#### Troubleshooting
##### Error: `/usr/bin/env: ‘bash\r’: No such file or directory`
&#8594; You're probably having [CRLF](https://developer.mozilla.org/en-US/docs/Glossary/CRLF) issues _(happened to me when cloning the repo on Windows initially before switching to WSL)_

&#8594; Use `dos2unix` _(install it if you don't have the command)_ on the 3 scripts : `dos2unix sfml.sh swig.sh build.sh`
##### Error: `swig: command not found`
&#8594; You probably either did not [install the `swig` package](#option-1---install-it-from-your-package-manager), or [built it locally](#option-2---build-it-locally) but forgot to follow the step 6 to update the path used by the `build.sh` script
##### Error: `Unable to find 'swig.swg'`
&#8594; You probably went for the swig [local build](#option-2---build-it-locally), but didn't check for its lib path. Please follow steps 3 to 5 of that section
##### Error: `patchelf: command not found`
&#8594; You probably didn't install `patchelf` as mentioned in step 1

You're now ready to go!

Feel free to open issues and/or PRs if you're running into problems that are not mentioned in the troubleshooting sub-sections

