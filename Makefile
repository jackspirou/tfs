DIST := dist
BIN := bin
EXECUTABLE := tfs

RELEASES ?= $(DIST)/$(EXECUTABLE)-linux-amd64 \
	$(DIST)/$(EXECUTABLE)-linux-386 \
	$(DIST)/$(EXECUTABLE)-linux-arm \
	$(DIST)/$(EXECUTABLE)-darwin-amd64

build: $(BIN)/$(EXECUTABLE)

test:
	go test -cover ./...

release: $(RELEASES)
