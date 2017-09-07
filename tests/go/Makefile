BIN=go-sfml

build:
	go build -o ${BIN} .

run:
	env LD_LIBRARY_PATH=${PWD}/../../CSFML/lib:${PWD}/../../SFML/lib DYLD_FRAMEWORK_PATH=${PWD}/../../SFML/extlibs/libs-osx/Frameworks/ ./${BIN}
