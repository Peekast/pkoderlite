VER_MAJOR := 0
VER_MINOR := 2
VER_PATCH := 4
VERSION := $(VER_MAJOR).$(VER_MINOR).$(VER_PATCH)
PACKAGE_NAME := pklite
DATE := $(shell sh -c date)
GOVERSION := $(shell sh -c "go version")

FLAGS := "-X 'main.buildTime=$(DATE)' -X 'main.compiler=$(GOVERSION)' -X 'main.majorVersion=$(VER_MAJOR)' -X 'main.minorVersion=$(VER_MINOR)' -X 'main.patchVersion=$(VER_PATCH)' "
.PHONY: package clean

all:
	go build -ldflags=$(FLAGS) -o $(PACKAGE_NAME) main.go

clean:
	rm -rf $(BUILD_DIR)
	rm -rf $(PACKAGE_NAME)
