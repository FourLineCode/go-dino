all:
	go run main.go

build:
	go build -o dino main.go

run:
	go build -o dino main.go && ./dino

clean:
	rm dino main bin