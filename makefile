build:
	@go build -o main.exe ./cmd/rest-api-golang/main.go

run: build
	@./main.exe