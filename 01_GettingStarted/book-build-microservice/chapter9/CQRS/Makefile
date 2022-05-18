build_read:
		echo "Build Products read server"
		CGO_ENABLED=0 GOOS=linux go build -o ./product-read/server ./product-read/. 

build_write:
	  echo "Build Products write server"
		CGO_ENABLED=0 GOOS=linux go build -o ./product-write/server ./product-write/. 

run: build_read build_write
	docker-compose up

build_read_win:
		echo "Build Products read server"
		go build -o ./product-read/server ./product-read/. 

build_write_win:
	  	echo "Build Products write server"
		go build -o ./product-write/server ./product-write/. 

run_win: build_read_win build_write_win
	docker-compose up