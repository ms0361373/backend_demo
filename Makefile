# git
version := $(shell git describe --tags --dirty --match 'v*' || echo 'dev')
commit := $(shell git rev-parse --short HEAD || echo '?')
buildtime := $(shell date +%Y-%m-%dT%H:%M:%S%z)

build_flag := "-w \
	-X proton.vir000.com/Jason/backend_demo/internal/ver.Version=$(version) \
	-X proton.vir000.com/Jason/backend_demo/internal/ver.BuildTime=$(buildtime) \
	-X proton.vir000.com/Jason/backend_demo/internal/ver.Commit=$(commit)"

all: build

lint:
	golangci-lint run

test:
	go test ./... -count=1

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
	rm coverage.out

build: ./*
	GOOS=linux GOARCH=amd64 go build -o build/backend_demo -a -v -ldflags $(build_flag) ./cmd/main.go

tar:
	tar zcvf backend_demo.tar.gz build

clean:
	rm -r build
