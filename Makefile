# Executable name
BINARY = gotop

# Go install directory
INSTALL_DIR = $(GOPATH)/bin

# Build target
build:
	go build -o ./bin/$(BINARY) ./cmd/gotop

# Run target
run: build
	./bin/$(BINARY) 

install: build
	install ./bin/$(BINARY) $(INSTALL_DIR)

# Clean target (optional - removes the binary)
clean:
	rm -f ./bin/$(BINARY)