build-linux:
	env GOOS=linux GOARCH=amd64 go build -o client client.go
	env GOOS=linux GOARCH=amd64 go build -o server server.go


clean:
	rm -f client server
	docker rmi $(shell docker images -f "dangling=true" -q)
