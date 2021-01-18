.PHONY: build test clean

image="data-collector"
version=latest

build:
	docker build -t $(imagename):$(version) . -f Dockerfile

version:
	@echo $(imagename) $(version)