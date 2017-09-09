#!/usr/bin/env bash

## [MacOS only]
## Clones and builds SFML & CSFML at the specified versions

set -e

SFML="SFML"
SFML_VERSION="2.4.0"
SFML_CMAKE_OPTIONS=""
SFML_CMAKE_OPTIONS+=" -DBUILD_SHARED_LIBS=ON"
SFML_CMAKE_OPTIONS+=" -DCMAKE_BUILD_TYPE=Release"
SFML_CMAKE_OPTIONS+=" -DCMAKE_INSTALL_FRAMEWORK_PREFIX=$PWD/$SFML/frameworks/"
SFML_CMAKE_OPTIONS+=" -DCMAKE_INSTALL_PREFIX=$PWD/$SFML/libs"
SFML_CMAKE_OPTIONS+=" -DIBTOOL=$(which ibtool)"
SFML_CMAKE_OPTIONS+=" -DSFML_BUILD_DOC=OFF"
SFML_CMAKE_OPTIONS+=" -DSFML_BUILD_EXAMPLES=ON"
SFML_CMAKE_OPTIONS+=" -DSFML_BUILD_FRAMEWORKS=OFF"
SFML_CMAKE_OPTIONS+=" -DSFML_INSTALL_XCODE_TEMPLATES=OFF"
SFML_CMAKE_OPTIONS+=" -DSFML_OPENGL_ES=OFF"

CSFML="CSFML"
CSFML_VERSION="2.4"
CSFML_CMAKE_OPTIONS=""
CSFML_CMAKE_OPTIONS+=" -DBUILD_SHARED_LIBS=ON"
CSFML_CMAKE_OPTIONS+=" -DCMAKE_BUILD_TYPE=Release"
CSFML_CMAKE_OPTIONS+=" -DCMAKE_INSTALL_PREFIX=$PWD/$CSFML/libs"
CSFML_CMAKE_OPTIONS+=" -DCSFML_BUILD_DOC=OFF"
CSFML_CMAKE_OPTIONS+=" -DCSFML_LINK_SFML_STATICALLY=ON"
CSFML_CMAKE_OPTIONS+=" -DSFML_INCLUDE_DIR=$PWD/$SFML/include"
CSFML_CMAKE_OPTIONS+=" -DCMAKE_MODULE_PATH=$PWD/$SFML/cmake/Modules/"
CSFML_CMAKE_OPTIONS+=" -DSFML_SYSTEM_LIBRARY_RELEASE=$PWD/$SFML/lib/libsfml-system.dylib"
CSFML_CMAKE_OPTIONS+=" -DSFML_WINDOW_LIBRARY_RELEASE=$PWD/$SFML/lib/libsfml-window.dylib"
CSFML_CMAKE_OPTIONS+=" -DSFML_NETWORK_LIBRARY_RELEASE=$PWD/$SFML/lib/libsfml-network.dylib"
CSFML_CMAKE_OPTIONS+=" -DSFML_GRAPHICS_LIBRARY_RELEASE=$PWD/$SFML/lib/libsfml-graphics.dylib"
CSFML_CMAKE_OPTIONS+=" -DSFML_AUDIO_LIBRARY_RELEASE=$PWD/$SFML/lib/libsfml-audio.dylib"

clone_and_checkout() {
	if [ ! -d "$1" ]; then
		echo -n "cloning $1 v$2..."
		git clone git@github.com:SFML/$1.git > /dev/null 2>&1
		echo " OK."
	else
		echo "$1 has already been cloned, skipping."
	fi
	echo -n "checking out to $1's '$2' tag..."
	git -C $1 checkout $2 > /dev/null 2>&1
	echo " OK."
}
clone_and_checkout "$SFML" "$SFML_VERSION"
clone_and_checkout "$CSFML" "$CSFML_VERSION"

clean_and_build() {
	cd "$1"
	echo "temporary moved to $PWD/"
	echo -en "\tcleansing $1..."
	git clean -fd > /dev/null 2>&1
	echo " OK."
	echo -en "\tbuilding metatools for $1..."
	cmake $2 . > /dev/null 2>&1
	echo " OK."
	echo -en "\tbuilding $1..."
	make > /dev/null 2>&1
	echo " OK."
	cd - > /dev/null 2>&1
	echo "moved back to $PWD/"
}
clean_and_build "$SFML" "$SFML_CMAKE_OPTIONS"
clean_and_build "$CSFML" "$CSFML_CMAKE_OPTIONS"
