BIN=go-sfml

build:
	gcc -L${PWD}/../../CSFML/lib -L${PWD}/../../SFML/lib -lcsfml-window -lcsfml-graphics -lsfml-window -I${PWD}/../../CSFML/include -o ${BIN} main.c

run:
	env LD_LIBRARY_PATH=${PWD}/../../CSFML/lib:${PWD}/../../SFML/lib DYLD_FRAMEWORK_PATH=${PWD}/../../SFML/extlibs/libs-osx/Frameworks/ ./${BIN}
