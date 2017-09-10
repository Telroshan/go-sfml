#!/usr/bin/env bash

## [Requires 'sfml.sh' & 'swig.sh' to have been run successfully]
## Builds the cgo bindings and install the final go packages

set -e

SFML="SFML"
CSFML="CSFML"
SFML_MODULES=(Audio Graphics System Window)

for m in "${SFML_MODULES[@]}"; do
	mm=$(echo $m | tr '[:upper:]' '[:lower:]')

	mkdir -p "$PWD/$mm"

	echo -n "building bindings for SFML's $m module..."
	cp "$PWD/interfaces/$m.i" "$PWD/$CSFML/include/$SFML/$m/$m.i"
	./swig/swig -go -cgo -intgosize 64 -I"$PWD/$CSFML/include" "$PWD/$CSFML/include/$SFML/$m/$m.i" > /dev/null 2>&1
	cp "$PWD/$CSFML/include/$SFML/$m/$mm.go" "$PWD/$CSFML/include/$SFML/$m/${m}_wrap.c" "$PWD/$mm"
	gsed -i "/import \"C\"/i \/\/ #cgo LDFLAGS: -lcsfml-$mm" "$PWD/$m/$mm.go"
	echo " OK."

	echo -n "compiling go package for SFML's $m module..."
	CGO_LDFLAGS="-L$PWD/CSFML/lib -lcsfml-$mm" CGO_CFLAGS="-I$PWD/$CSFML/include" go install "./$mm" > /dev/null 2>&1
	echo " OK."
done
