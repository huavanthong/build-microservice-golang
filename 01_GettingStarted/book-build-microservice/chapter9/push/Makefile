build_receiver_linux:
		echo "Build message receiver server"
		CGO_ENABLED=0 GOOS=linux go build -o ./publisher/server ./receiver/. 

build_receiver_win:
		echo "Build message receiver server"
		go build -o ./publisher/server ./receiver/. 

build_publisher_linux:
		echo "Build message publisher"
		CGO_ENABLED=0 GOOS=linux go build -o ./publisher/server ./publisher/. 

build_publisher_win:
		echo "Build message publisher"
		go build -o ./publisher/server ./publisher/. 

run_linux: build_receiver_linux build_publisher_linux
	docker-compose up --build

run_win: build_receiver_win build_publisher_win
	docker-compose up --build