BIN=go-sfml

build:
	gcc -L${PWD}/../../CSFML/lib -L${PWD}/../../SFML/lib -lcsfml-audio -lsfml-audio -lcsfml-graphics -lsfml-graphics -lcsfml-system -lsfml-system -lcsfml-window -lsfml-window -I${PWD}/../../CSFML/include -o ${BIN} main.c

run: build
	env LD_LIBRARY_PATH=${PWD}/../../CSFML/lib:${PWD}/../../SFML/lib DYLD_FRAMEWORK_PATH=${PWD}/../../SFML/extlibs/libs-osx/Frameworks/ ./${BIN}
