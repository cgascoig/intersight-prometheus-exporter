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

GO_MODULE := github.com/cgascoig/intersight-prometheus-exporter
DOCKER_IMAGE_ID := ghcr.io/cgascoig/intersight-prometheus-exporter

GO_CMD ?= go
GO_BUILD_CMD := $(GO_CMD) build -v 
GO_BUILD_FLAGS := -ldflags "-X main.commit=`git rev-parse HEAD`"
GO_PATH ?= $(shell go env GOPATH)

all: build/intersight-prometheus-exporter
.PHONY: all

clean:
> rm -Rf build tmp
.PHONY: clean

# Go unit tests
test: $(shell find cmd -name \*.go -type f) $(shell find pkg -name \*.go -type f) go.mod
> $(GO_CMD) test -v $(GO_MODULE)/cmd
.PHONY: test

build/intersight-prometheus-exporter: $(shell find cmd -name \*.go -type f) $(shell find pkg -name \*.go -type f) go.mod
> mkdir -p $(@D)
> $(GO_BUILD_CMD) -o "$@" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd

build/intersight-prometheus-exporter-linux_amd64: build/intersight-prometheus-exporter
> mkdir -p $(@D)
> GOOS=linux GOARCH=amd64 $(GO_BUILD_CMD) -o "$@" $(GO_BUILD_FLAGS) $(GO_MODULE)/cmd

containers: tmp/.intersight-prometheus-exporter-docker-image.sentinel
.PHONY: containers

containers-push: tmp/.intersight-prometheus-exporter-docker-image-push.sentinel

tmp/.intersight-prometheus-exporter-docker-image.sentinel: build/intersight-prometheus-exporter-linux_amd64 Dockerfile
> mkdir -p $(@D)
> docker build . -t "$(DOCKER_IMAGE_ID):latest"
> touch $@

tmp/.intersight-prometheus-exporter-docker-image-push.sentinel: tmp/.intersight-prometheus-exporter-docker-image.sentinel
> # Check if there are any uncommitted changes
> if [[ -n $$(git status --porcelain) ]]; then echo "repo is dirty"; exit 1; fi
> image_id="$(DOCKER_IMAGE_ID):$$(git rev-parse HEAD)"
> echo "Tagging image $${image_id}"
> docker tag "$(DOCKER_IMAGE_ID):latest" "$${image_id}"
> echo "Pushing image $${image_id}"
> docker push "$${image_id}"
> touch $@