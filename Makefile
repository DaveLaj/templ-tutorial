lint: 
	golangci-lint run --sort-results --show-stats --allow-parallel-runners

templgen:
	templ generate

run: lint
	go run main/main.go

build: lint
	go build -o bin/main.exe main/main.go

templwatch: 
	templ generate --watch --proxy="http://localhost:8080" --cmd="make run"
