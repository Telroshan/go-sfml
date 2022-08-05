#!/usr/bin/env bash

set -e

SWIG="swig"
SWIG_VERSION="v4.0.2"

clone_and_checkout() {
	if [ ! -d "$1" ]; then
		echo -n "cloning $1 v$2..."
		git clone https://github.com/swig/$1.git
		echo " OK."
	else
		echo "$1 has already been cloned, skipping."
	fi
	echo -n "checking out to $1's '$2' tag..."
	git -C $1 checkout $2
	echo " OK."
}
clone_and_checkout "$SWIG" "$SWIG_VERSION"

echo -n "building $SWIG:$SWIG_VERSION..."
cd $SWIG
./autogen.sh
./configure
make
cd -
echo " OK."