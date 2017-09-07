#!/usr/bin/env bash

set -e

SFML="SFML"
SFML_VERSION="2.4.0"

CSFML="CSFML"
CSFML_VERSION="2.4"

## Audio
swig -go -cgo -intgosize 64 -I"$PWD/CSFML/include" $PWD/CSFML/include/SFML/Audio/Audio.i
CGO_LDFLAGS="-L$PWD/CSFML/lib -lcsfml-audio" CGO_CFLAGS="-I$PWD/CSFML/include" go install -x ./CSFML/include/SFML/Audio

## Graphics
swig -go -cgo -intgosize 64 -I"$PWD/CSFML/include" $PWD/CSFML/include/SFML/Graphics/Graphics.i
CGO_LDFLAGS="-L$PWD/CSFML/lib -lcsfml-graphics" CGO_CFLAGS="-I$PWD/CSFML/include" go install -x ./CSFML/include/SFML/Graphics

## Window
swig -go -cgo -intgosize 64 -I"$PWD/CSFML/include" $PWD/CSFML/include/SFML/Window/Window.i
CGO_LDFLAGS="-L$PWD/CSFML/lib -lcsfml-window" CGO_CFLAGS="-I$PWD/CSFML/include" go install -x ./CSFML/include/SFML/Window