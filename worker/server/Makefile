.PHONY: go-build

go-build: worker.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w' -o worker ./worker.go
