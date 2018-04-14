GO=go
FILES=`go list ./.../`

.PHONY: all vet test build clean coverage

all: vet test build bench

vet:
	$(GO) vet $(FILES)

build:
	$(GO) build -o bin/caesar

test: vet
	$(GO) test -race -cover -v $(FILES)

clean:
	rm -vrf bin
	rm coverage.out

bench:
	$(GO) test -v -bench=. -benchtime=5s

coverage:
	go tool cover -html=coverage.out
