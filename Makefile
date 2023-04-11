BINARY_NAME = application
exec = application

build:
	ls;
	cd ./backend && go mod download;
	cd ./backend && go build -o bin/${BINARY_NAME}

run:
	cd ./backend && ./bin/${BINARY_NAME}

run_app: build run

clean:
	go clean
