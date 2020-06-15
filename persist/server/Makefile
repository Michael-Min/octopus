.PHONY: go-build

go-build: itemsaver.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w' -o itemsave ./itemsaver.go
