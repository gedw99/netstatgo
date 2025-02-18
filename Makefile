.PHONY: build-mac-intel build-mac-silicon build-linux build-windows build-all clean

build-mac-intel:
	GOOS=darwin GOARCH=amd64 go build -o netstatgo-macos-amd64

build-mac-silicon:
	GOOS=darwin GOARCH=arm64 go build -o netstatgo-macos-arm64

build-linux:
	GOOS=linux GOARCH=amd64 go build -o netstatgo-linux-amd64

build-windows:
	GOOS=windows GOARCH=amd64 go build -o netstatgo-windows-amd64.exe

build-all: build-mac-intel build-mac-silicon build-linux build-windows

clean:
	rm -f netstatgo-*
	
run:
	go run main.go