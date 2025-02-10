# Project - Golang Webapp

This app is a small [Gin](https://gin-gonic.com/) based webserver to display a message.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

TODO:
 - [x] make wrapper
 - [x] sast - github action
 - [x] lint - github action
 - [x] build - github action
 - [x] release - github action
 - [x] chain - github action

Fluff:
- [x] Pass TAG in trigger payload
- [ ] Generate Favicon.ico
- [ ] Publish scan results
- [ ] Added health check


## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create container
```bash
make docker-run
```

Shutdown Container
```bash
make docker-down
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```

## Linting
Running locally...

### golangci-lint
Docker
```
docker run -t --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.63.4 golangci-lint run -v
```
Go Binary
```
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4
```

### semgrep
Docker 
```
docker run -it --rm -v $(pwd):/app -w /app semgrep/semgrep semgrep scan --config=p/gosec --disable-version-check --oss-only --metrics=off --verbose .
```
Additional args for later development
`--error`
`--junit-xml`
`--junit-xml-output=VAL`
