VERSION=`git describe --tags --match v[0-9]* 2> /dev/null`
EXE=livemidi

.PHONY: all
all: run

.PHONY: run
run: build
	@./build/$(EXE)

.PHONY: build
build:
	@CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o build/$(EXE) ./cmd/livemidi/main.go

.PHONY: test
test:
	@go test ./... -v

.PHONY: install
install:
	@cp ./build/$(EXE) /usr/bin/$(EXE)

.PHONY: dist
dist: dist-linux dist-win

.PHONY: dist-linux
dist-linux: dist-linux-amd64

.PHONY: dist-linux-amd64
dist-linux-amd64:
	@rm -f build/$(EXE)
	@CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -o build/$(EXE) ./cmd/livemidi/main.go
	@tar -czvf "$(EXE)-${VERSION}-linux-amd64.tar.gz" build/$(EXE)

.PHONY: dist-win
dist-win: dist-win-amd64

.PHONY: dist-win-amd64
dist-win-amd64:
	@rm -f build/$(EXE).exe
	@CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_LDFLAGS="-static -static-libgcc -static-libstdc++" CGO_ENABLED=1 GOARCH=amd64 GOOS=windows go build -ldflags="-H windowsgui" -o build/$(EXE).exe ./cmd/livemidi/main.go
	@zip "$(EXE)-${VERSION}-win-amd64.zip" build/$(EXE).exe

.PHONY: clean
clean:
	@rm -Rf build/*
