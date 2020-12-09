run:
	go run main.go

build:
	go build -o latom

playground:
	make build && ./latom playground
