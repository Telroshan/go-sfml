.PHONY: toc

toc:
	docker run --rm -it -v ${PWD}:/usr/src jorgeandrada/doctoc --github
