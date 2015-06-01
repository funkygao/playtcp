all:client server

client:client.go
	go build -o client client.go

server:server.go
	go build -o server server.go

clean:
	rm -f client server
