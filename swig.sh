#!/usr/bin/env bash

set -e

SWIG="swig"
OUTPUT_DIR=$PWD/$SWIG
SWIG_VERSION="rel-4.0.1"

clone_and_checkout() {
	if [ ! -d "$1" ]; then
		echo -n "cloning $1 v$2..."
		git clone git@github.com:swig/$1.git $1/src > /dev/null 2>&1
		echo " OK."
	else
		echo "$1 has already been cloned, skipping."
	fi
	echo -n "checking out to $1's '$2' tag..."
	git -C $1/src checkout $2 > /dev/null 2>&1
	echo " OK."
}
clone_and_checkout "$SWIG" "$SWIG_VERSION"

echo -n "building $SWIG:$SWIG_VERSION..."
(
cd ${SWIG}/src
./autogen.sh > /dev/null 2>&1
./configure --prefix=$OUTPUT_DIR > /dev/null 2>&1
make > /dev/null 2>&1
make install > /dev/null 2>&1
)
echo " OK."