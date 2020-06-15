.PHONY: go-build

go-build: main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w' -o main ./main.go
