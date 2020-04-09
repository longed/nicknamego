EXECUTABLE		:= nicknamego
TAR_GZ			:= $(EXECUTABLE).tar.gz
TARGET_DIR		:= targets
BIN_DIR			:= bin
CONF_DIR		:= conf
DOC_DIR			:= doc
LIB_DIR			:= lib
CONFIG_FILE		:= config.toml

#--------------------------------------------------------
# build

GOPATH ?= $(shell go env GOPATH)

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

GO              := GO111MODULE=on go
GOBUILD         := $(GO) build -o $(EXECUTABLE) .

.PHONY: integrate build clean test

integrate: build
	# create directories by script
	bash assemly.sh mkdir_loop $(TARGET_DIR) $(TARGET_DIR)/$(BIN_DIR) $(TARGET_DIR)/$(CONF_DIR) $(TARGET_DIR)/$(DOC_DIR) $(TARGET_DIR)/$(LIB_DIR)
	# copy files | config-file scripts README.md docs/*
	yes | cp -t $(TARGET_DIR)/$(CONF_DIR) $(CONFIG_FILE)
	yes | cp -t $(TARGET_DIR)/$(BIN_DIR) build/nicknamego.sh
	yes | cp -t $(TARGET_DIR)/$(DOC_DIR) README.md
	yes | cp -r -t $(TARGET_DIR)/$(DOC_DIR) docs/
	# move files 
	yes | mv -t $(TARGET_DIR)/$(LIB_DIR) $(EXECUTABLE)
	# compress targets
	cd $(TARGET_DIR) &&	tar cvzf $(TAR_GZ) * && cd -

build: clean
	$(GOBUILD)

clean:
	-rm -rf $(TARGET_DIR)

test: