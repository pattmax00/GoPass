# The current version number of the program
VERSION := 1.2.2

# List of OS and architecture combinations to build
BUILD_OS_ARCH := \
	darwin/amd64 \
	linux/386 \
	linux/amd64 \
	windows/386 \
	windows/amd64

# Build all targets
all: $(BUILD_OS_ARCH)

# Build targets in the form of `OS/ARCH`
darwin/amd64:
	GOOS=darwin \
	GOARCH=amd64 \
	go build -ldflags="-s -w" -o "gopass-darwin-amd64-$(VERSION)" main.go

linux/386:
	GOOS=linux \
	GOARCH=386 \
	go build -ldflags="-s -w" -o "gopass-linux-386-$(VERSION)" main.go

linux/amd64:
	GOOS=linux \
	GOARCH=amd64 \
	go build -ldflags="-s -w" -o "gopass-linux-amd64-$(VERSION)" main.go

windows/386:
	GOOS=windows \
	GOARCH=386 \
	go build -ldflags="-s -w" -o "gopass-windows-386-$(VERSION)" main.go

windows/amd64:
	GOOS=windows \
	GOARCH=amd64 \
	go build -ldflags="-s -w" -o "gopass-windows-amd64-$(VERSION)" main.go
