ldflags=-s -w

default: fmt build

fmt:
	gofmt -s -l -w .

lint:
	golangci-lint run

build:
	go build -tags='netgo' -ldflags='$(ldflags)'

prof: build
	./SandGo -cpuprofile cpu.prof
	go tool pprof hashbench cpu.prof

test:
	go test

bench:
	go test -run=XXX -bench=.
