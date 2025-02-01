# Executable name
BINARY = gotop

# Build target
build:
	go build -o ./bin/$(BINARY) ./cmd/gotop

# Run target
run: build
	./bin/$(BINARY) 

# Clean target (optional - removes the binary)
clean:
	rm -f ./bin/$(BINARY)