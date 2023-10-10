# Makefile for echo container
FILES = logminute.go Dockerfile
all: $(FILES)
	@echo "NOTE: building for amd64"
	mkdir -p build
	cp $(FILES) build
	cd build && \
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o logminute logminute.go && \
	upx ./logminute && \
	docker build -t logminute .
clean:
	rm -rf build
rm:
	docker rmi logminute
