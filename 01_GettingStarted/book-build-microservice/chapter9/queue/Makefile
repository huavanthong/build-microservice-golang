build_read_linux:
		echo "Build Products read server"
		CGO_ENABLED=0 GOOS=linux go build -o ./reader/server ./reader/.

build_read_win:
		echo "Build Products read server"
		go build -o ./reader/server ./reader/. 

build_write_linux:
		echo "Build Products write server"
		CGO_ENABLED=0 GOOS=linux go build -o ./writer/server ./writer/. 

build_write_win:
		echo "Build Products write server"
		go build -o ./writer/server ./writer/. 

run_linux: build_write_linux build_read_linux
	docker-compose up --build

run_win: build_write_win build_read_win
	docker-compose up --build
