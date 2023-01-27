export GO111MODULE = on

build_tags := $(strip $(BUILD_TAGS))
BUILD_FLAGS := -tags "$(build_tags)"

OUT_DIR = ./build

.PHONY: all build install clean

all: build install

build: go.sum
	go build -mod=readonly $(BUILD_FLAGS) -o $(OUT_DIR)/cryptossd ./cmd/cryptossd

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/cryptossd

lint:
	GO111MODULE=off go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run --timeout 5m0s --allow-parallel-runners
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify

clean:
	go clean
	rm -rf $(OUT_DIR)
