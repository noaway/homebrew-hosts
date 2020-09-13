clean:
	go clean

build:
	go build -v

build-release:
	bash .tools/build.sh