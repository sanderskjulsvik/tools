actions:
	act

win-build:
	go build -o bin\\ .\\...

win-test:
	go test .\\...

win-deps:
	choco install golangci-lint act-cli

test:
	go test ./...
