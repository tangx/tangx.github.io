
version ?= 3.12 

build:
	docker build -f 1.Dockerfile -t example.com/demo/arg:$(version) . --no-cache --build-arg=VERSION=$(version)

build.1:
	version=3.11 make build
	version=3.12 make build


build.2:
	docker build -f 2.Dockerfile -t example.com/demo/arg:$(version) . --no-cache --build-arg=VERSION=$(version)
