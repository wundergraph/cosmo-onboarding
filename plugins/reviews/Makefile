
.PHONY: build build-linux test generate install-wgc

install-wgc:
	@which wgc > /dev/null 2>&1 || npm install -g wgc@latest

make: build

test: install-wgc
	wgc router plugin test .

generate: install-wgc
	wgc router plugin generate .

publish: generate
	wgc router plugin publish .

build: install-wgc
	wgc router plugin build . --debug

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/linux_arm64 ./src
