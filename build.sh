#!/usr/bin/env bash

## [Requires 'sfml.sh' & 'swig.sh' to have been run successfully]
## Builds the cgo bindings and install the final go packages

set -e

SFML="SFML"
CSFML="CSFML"
SFML_MODULES=(Audio Graphics System Window)

for m in "${SFML_MODULES[@]}"; do
	mm=$(echo $m | tr '[:upper:]' '[:lower:]')

	echo -n "Fixing missing library links to SFML..."
	patchelf --set-rpath "$PWD/$SFML/lib" "$CSFML/lib/libcsfml-$mm.so"
	echo " OK."

	mkdir -p "$PWD/$mm"

	echo -n "building bindings for SFML's $m module..."
	cp "$PWD/interfaces/$m.i" "$PWD/$CSFML/include/$SFML/$m/$m.i"
	swig -go -cgo -intgosize 64 -I"$PWD/$CSFML/include" "$PWD/$CSFML/include/$SFML/$m/$m.i"
	cp "$PWD/$CSFML/include/$SFML/$m/$mm.go" "$PWD/$CSFML/include/$SFML/$m/${m}_wrap.c" "$PWD/$mm"
	sed -i "/import \"C\"/i \/\/ #cgo LDFLAGS: -lcsfml-$mm" "$PWD/$m/$mm.go"
	echo " OK."

	echo -n "compiling go package for SFML's $m module..."
	GO111MODULE=off CGO_LDFLAGS="-L$PWD/CSFML/lib -lcsfml-$mm" CGO_CFLAGS="-I$PWD/$CSFML/include" go install "./$mm"
	echo " OK."
done

echo "Build completed."