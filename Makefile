# build
build:buildplugin buildmain
buildplugin:
	CGO_ENABLED=1 go build -o plugin.so -buildmode=plugin plugin/plugin.go
buildmain:
	CGO_ENABLED=1 go build -o main main.go
