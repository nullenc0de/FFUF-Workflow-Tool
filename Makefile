.PHONY: build install clean

build:
	go build -o bin/ffuf-workflow ./cmd/ffuf-workflow

install: build
	go install ./cmd/ffuf-workflow

clean:
	rm -rf bin
