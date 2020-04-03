PROJECT=nicknamego
EXECUTABLE=nicknamego
GOPATH ?= $(shell go env GOPATH)

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

GO              := GO111MODULE=on go
GOBUILD         := $(GO) build $(BUILD_FLAG) .

.PHONY: build clean

build:
	$(GOBUILD)

clean:
	-rm -rf $(EXECUTABLE)