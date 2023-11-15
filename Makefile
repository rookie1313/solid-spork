pkg:
	go mod tidy

server:
	go run main.go

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15-alpine

initdb:
	docker exec -it postgres createdb --username=root --owner=root solid_spork
