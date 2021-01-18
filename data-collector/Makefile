.PHONY: build test clean

image=data-collector
version=latest

build:
	docker build -t $(image):$(version) . -f Dockerfile

run:
	docker run -p 2137:1883 $(image):$(version)

version:
	@echo $(imagename) $(version)