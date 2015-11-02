all: built-in built-check built-out

built-in: in/main.go
	GOARCH=amd64 GOOS=linux go build -o built-in in/main.go

built-out: out/main.go
	GOARCH=amd64 GOOS=linux go build -o built-out out/main.go

built-check: check/main.go
	GOARCH=amd64 GOOS=linux go build -o built-check check/main.go
