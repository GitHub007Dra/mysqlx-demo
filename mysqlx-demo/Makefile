all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo .

run:
	docker-compose up --build mysqlx-demo

clean:
	go clean
