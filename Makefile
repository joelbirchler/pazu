run:
	go run main.go

fmt:
	go ftm ./...

lint:
	golint ./...

test:
	go test -v ./...