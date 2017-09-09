#!/usr/bin/env bash

set -e

SFML="SFML"
SFML_VERSION="2.4.0"

CSFML="CSFML"
CSFML_VERSION="2.4"

SFML_MODULES=(Audio Graphics System Window)

for m in "${SFML_MODULES[@]}"; do
	mm=$(echo $m | tr '[:upper:]' '[:lower:]')

	mkdir -p "$PWD/$mm"

	echo -n "building bindings for SFML's $m module..."
	cp "$m.i" "$PWD/$CSFML/include/$SFML/$m/$m.i"
	swig -go -cgo -intgosize 64 -I"$PWD/$CSFML/include" "$PWD/$CSFML/include/$SFML/$m/$m.i" > /dev/null 2>&1
	cp "$PWD/$CSFML/include/$SFML/$m/$mm.go" "$PWD/$CSFML/include/$SFML/$m/${m}_wrap.c" "$PWD/$mm"
	echo " OK."

	echo -n "compiling go package for SFML's $m module..."
	CGO_LDFLAGS="-L$PWD/CSFML/lib -lcsfml-$mm" CGO_CFLAGS="-I$PWD/$CSFML/include" go install "./$mm" > /dev/null 2>&1
	echo " OK."
done