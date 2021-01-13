dev:
	go build main.go
	go run main.go

restart:
	docker-compose down
	docker-compose up --build -d

stop:
	docker-compose down