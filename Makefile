build:
	go build -o bin/main cmd/server/main.go

clean:
	rm bin/*

run:
	go run cmd/server/main.go

