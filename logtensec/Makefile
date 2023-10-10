# Makefile for echo container
FILES = logtensec.go Dockerfile
all: $(FILES)
	@echo "NOTE: building for amd64"
	mkdir -p build
	cp $(FILES) build
	cd build && \
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o logtensec logtensec.go && \
	upx ./logtensec && \
	docker build -t logtensec .
clean:
	rm -rf build
rm:
	docker rmi logtensec
