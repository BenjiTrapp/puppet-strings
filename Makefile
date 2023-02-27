all: build run

build:
			go fmt .
			go build .

run:
			./puppet-strings