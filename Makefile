VERSION = $(shell git describe --tags)
VER = $(shell git describe --tags --abbrev=0)
DATE = $(shell date -u '+%Y-%m-%d_%H:%M:%S%Z')
FLAG_MODULE = GO111MODULE=on
FLAGS_SHARED = $(FLAG_MODULE) CGO_ENABLED=0 GOARCH=amd64
FLAGS_LD=-ldflags "-X github.com/dimus/glgh.Build=${DATE} \
                  -X github.com/dimus/glgh.Version=${VERSION}"
GOCMD=go
GOINSTALL=$(GOCMD) install $(FLAGS_LD)
GOBUILD=$(GOCMD) build $(FLAGS_LD)
GOCLEAN=$(GOCMD) clean
GOGET = $(GOCMD) get

all: install

test: deps install
	$(FLAG_MODULE) go test ./...

deps:
	$(GOCMD) mod download; \

build:
	cd glgh; \
	$(GOCLEAN); \
	$(FLAGS_SHARED) $(GOBUILD);

dc: build
	docker-compose build;

release:
	cd glgh; \
	$(GOCLEAN); \
	$(FLAGS_SHARED) GOOS=linux $(GOBUILD); \
	tar zcvf /tmp/glgh-${VER}-linux.tar.gz glgh; \
	$(GOCLEAN);

install:
	cd glgh; \
	$(FLAGS_SHARED) $(GOINSTALL);

help:
	@echo
	@echo Build commands:
	@echo "  make [all]     - Install glgh"
	@echo "  make install   - Install glgh"
	@echo "  make help      - This help"
	@echo
