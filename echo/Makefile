# Makefile for echo container
FILES = echo.go Dockerfile
all: $(FILES)
	@echo "NOTE: building for amd64"
	mkdir -p build
	cp $(FILES) build
	cd build && \
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o echo echo.go && \
	upx ./echo && \
	docker build -t echo .
clean:
	rm -rf build
rm:
	docker rmi echo
