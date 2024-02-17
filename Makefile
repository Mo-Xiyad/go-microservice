build:
	@go build -o bin/micoservice-structure

run: build
	@./bin/micoservice-structure
