#write here path for your project
PROJECT :=
GIT_COMMIT := $(shell git rev-parse HEAD)
VERSION := 1.0
APP_NAME := cinema

all: run

run:
	go install -ldflags="-X '$(PROJECT)/version.Version=$(VERSION)' \
	-X '$(PROJECT)/version.Commit=$(GIT_COMMIT)'" && $(APP_NAME)