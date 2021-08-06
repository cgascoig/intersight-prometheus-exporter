SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
# .DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

ifeq ($(origin .RECIPEPREFIX), undefined)
  $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
endif
.RECIPEPREFIX = >

GO_MODULE := github.com/cgascoig/intersight-metrics
GO_CMD ?= go
GO_BUILD_CMD := $(GO_CMD) build -v 
GO_BUILD_FLAGS := -ldflags "-X main.commit=`git rev-parse HEAD`"
GO_PATH ?= $(shell go env GOPATH)

all: build/prom-intersight-metrics
.PHONY: all

clean:
> rm -Rf build tmp
.PHONY: clean

# Go unit tests
test: $(shell find cmd -name \*.go -type f) $(shell find pkg -name \*.go -type f) go.mod
> $(GO_CMD) test -v $(GO_MODULE)/cmd
.PHONY: test

build/prom-intersight-metrics: $(shell find cmd -name \*.go -type f) $(shell find pkg -name \*.go -type f) go.mod
> mkdir -p $(@D)
> $(GO_BUILD_CMD) -o "$@" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd

build/prom-intersight-metrics-linux_amd64: build/prom-intersight-metrics
> mkdir -p $(@D)
> GOOS=linux GOARCH=amd64 $(GO_BUILD_CMD) -o "$@" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd

containers: tmp/.prom-intersight-metrics-docker-image.sentinel
.PHONY: containers

containers-push: tmp/.prom-intersight-metrics-docker-image-push.sentinel

tmp/.prom-intersight-metrics-docker-image.sentinel: build/prom-intersight-metrics-linux_amd64 Dockerfile
> mkdir -p $(@D)
> image_id="cgascoig/prom-intersight-metrics:latest"
> docker build . -t "$${image_id}"
> touch $@

tmp/.prom-intersight-metrics-docker-image-push.sentinel: tmp/.prom-intersight-metrics-docker-image.sentinel
> image_id="cgascoig/prom-intersight-metrics:latest"
> docker push "$${image_id}"