build:
	go build -o bin/ecs

test:
	go test ./tests/**/*.go -v
