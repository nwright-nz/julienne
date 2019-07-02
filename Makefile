 # Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    BINARY_NAME=julienne 
    BINARY_WINDOWS=julienne.exe
    all: build
    build:
	$(GOBUILD) -o $(BINARY_NAME) -v
    clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_WINDOWS)
    
    # Cross compilation
    build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
    build-windows:
	GOOS=windows GOARCH=386 $(GOBUILD) -o $(BINARY_WINDOWS) -v
