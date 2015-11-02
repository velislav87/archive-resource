all: built-in built-check

built-in: in/main.go
	GOARCH=amd64 GOOS=linux go build -o built-in in/main.go

built-check: out/main.go
	GOARCH=amd64 GOOS=linux go build -o built-out out/main.go

built-check: check/main.go
	GOARCH=amd64 GOOS=linux go build -o built-check check/main.go
