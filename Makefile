# Variables
export GO111MODULE ?= on
export GOPROXY ?= https://proxy.golang.org
export GOSUMDB ?= sum.golang.org

# Applications name
CGO    ?= 0
BINARY ?= mqtt-client

# Target build windows/amd64 linux/amd64
PLATFORMS ?=  darwin/amd64 linux/arm

# Macros to sub info from platforms
temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

# GO build flags
LDFLAGS="-s -w"

#
# Target build
# Can use 	#upx dist/$(BINARY)-$(os)-$(arch) --> to reduce target binary size
#
release: $(PLATFORMS)

$(PLATFORMS):
	CGO_ENABLED=$(CGO) GOOS=$(os) GOARCH=$(arch) go build $(GCFLAGS) -ldflags=$(LDFLAGS) -o dist/$(BINARY)-$(os)-$(arch) cmd/$(BINARY)/main.go

.PHONY: release $(PLATFORMS)
