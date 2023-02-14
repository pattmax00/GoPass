# The current version number of the program
VERSION := 1.3.1

# List of OS and architecture combinations to build
BUILD_OS_ARCH := \
	darwin/amd64 \
	freebsd/386 \
	freebsd/amd64 \
	freebsd/arm \
	linux/386 \
	linux/amd64 \
	linux/arm \
	plan9/386 \
	plan9/amd64 \
	plan9/arm \
	windows/386 \
	windows/amd64 \
	windows/arm

# Build all targets
all: $(BUILD_OS_ARCH)

# Build targets in the form of `OS/ARCH`
darwin/amd64:
	GOOS=darwin \
	GOARCH=amd64 \
	go build -ldflags="-s -w" -o "gopass-darwin-amd64-$(VERSION)" main.go

freebsd/386:
	GOOS=freebsd \
	GOARCH=386 \
	go build -ldflags="-s -w" -o "gopass-freebsd-386-$(VERSION)" main.go

freebsd/amd64:
	GOOS=freebsd \
	GOARCH=amd64 \
	go build -ldflags="-s -w" -o "gopass-freebsd-amd64-$(VERSION)" main.go

freebsd/arm:
	GOOS=freebsd \
	GOARCH=arm \
	go build -ldflags="-s -w" -o "gopass-freebsd-arm-$(VERSION)" main.go

linux/386:
	GOOS=linux \
	GOARCH=386 \
	go build -ldflags="-s -w" -o "gopass-linux-386-$(VERSION)" main.go

linux/amd64:
	GOOS=linux \
	GOARCH=amd64 \
	go build -ldflags="-s -w" -o "gopass-linux-amd64-$(VERSION)" main.go

linux/arm:
	GOOS=linux \
	GOARCH=arm \
	go build -ldflags="-s -w" -o "gopass-linux-arm-$(VERSION)" main.go

plan9/386:
	GOOS=plan9 \
	GOARCH=386 \
	go build -ldflags="-s -w" -o "gopass-plan9-386-$(VERSION)" main.go

plan9/amd64:
	GOOS=plan9 \
	GOARCH=amd64 \
	go build -ldflags="-s -w" -o "gopass-plan9-amd64-$(VERSION)" main.go

plan9/arm:
	GOOS=plan9 \
	GOARCH=arm \
	go build -ldflags="-s -w" -o "gopass-plan9-arm-$(VERSION)" main.go

windows/386:
	GOOS=windows \
	GOARCH=386 \
	go build -ldflags="-s -w" -o "gopass-windows-386-$(VERSION)" main.go

windows/amd64:
	GOOS=windows \
	GOARCH=amd64 \
	go build -ldflags="-s -w" -o "gopass-windows-amd64-$(VERSION)" main.go

windows/arm:
	GOOS=windows \
	GOARCH=arm \
	go build -ldflags="-s -w" -o "gopass-windows-arm-$(VERSION)" main.go
