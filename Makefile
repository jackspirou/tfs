DIST := dist
BIN := bin
EXECUTABLE := tfs

LDFLAGS += -X "main.version=$(VERSION)"
VERSION ?= $(shell cat VERSION)

RELEASES ?= $(DIST)/$(EXECUTABLE)-linux-amd64 \
	$(DIST)/$(EXECUTABLE)-linux-386 \
	$(DIST)/$(EXECUTABLE)-linux-arm \
	$(DIST)/$(EXECUTABLE)-darwin-amd64

clean:
	@go clean -i ./...
	@rm -rf $(BIN) $(DIST)

deps:
	go get -t ./...

build: $(BIN)/$(EXECUTABLE)

test:
	go test -cover ./...

release: $(RELEASES)

install: $(BIN)/$(EXECUTABLE)
	cp $< $(GOPATH)/bin/

$(BIN)/$(EXECUTABLE):
	CGO_ENABLED=0 go build -ldflags '-s -w $(LDFLAGS)' -o $@

$(BIN)/%/$(EXECUTABLE): GOOS=$(firstword $(subst -, ,$*))
$(BIN)/%/$(EXECUTABLE): GOARCH=$(subst .exe,,$(word 2,$(subst -, ,$*)))
$(BIN)/%/$(EXECUTABLE):
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags '-s -w $(LDFLAGS)' -o $@

$(DIST)/$(EXECUTABLE)-%: GOOS=$(firstword $(subst -, ,$*))
$(DIST)/$(EXECUTABLE)-%: GOARCH=$(subst .exe,,$(word 2,$(subst -, ,$*)))
$(DIST)/$(EXECUTABLE)-%: $(BIN)/%/$(EXECUTABLE)
	@mkdir -p $(DIST)
	cp $(BIN)/$*/$(EXECUTABLE) $(DIST)/$(EXECUTABLE)-$(VERSION)-$(GOOS)-$(GOARCH)

.PHONY: clean deps build test
.PRECIOUS: $(BIN)/%/$(EXECUTABLE)
