EXECUTABLE		:= nicknamego
TARGET_DIR		:= target
BIN_DIR			:= bin
CONF_DIR		:= conf
DOC_DIR			:= doc
LIB_DIR			:= lib
CONFIG_FILE		:= config.toml
#----------------------------

GOPATH ?= $(shell go env GOPATH)

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

GO              := GO111MODULE=on go
GOBUILD         := $(GO) build -o $(EXECUTABLE) .

.PHONY: integrate

integrate: build
	# create directories by script
	bash assemly.sh mkdir_loop $(TARGET_DIR) $(TARGET_DIR)/$(BIN_DIR) $(TARGET_DIR)/$(CONF_DIR) $(TARGET_DIR)/$(DOC_DIR) $(TARGET_DIR)/$(LIB_DIR)
	# copy files
	yes | cp -t $(TARGET_DIR)/$(CONF_DIR) $(CONFIG_FILE)
	# move files 
	yes | mv -t $(TARGET_DIR)/$(LIB_DIR) $(EXECUTABLE)

build: clean
	$(GOBUILD)

clean:
	-rm -rf $(TARGET_DIR)