tidy:
	go mod tidy

test:
	go test ./... -v

test-integers:
	go test ./integers -v