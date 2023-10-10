# Makefile for getecho container
FILES = getecho.go Dockerfile
all: $(FILES)
	@echo "NOTE: building for amd64"
	mkdir -p build
	cp $(FILES) build
	cd build && \
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o getecho getecho.go && \
	upx ./getecho && \
	docker build -t getecho .
clean:
	rm -rf build
rm:
	docker rmi getecho
