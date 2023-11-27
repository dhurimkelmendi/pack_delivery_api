#This is only meant to be used in development. DO NOT USE THIS TO SERVE PRODUCTION SERVER
serve: main.go
	go run main.go

migrate: main.go
	go run main.go migrate up

reset: main.go
	go run main.go migrate reset

downgrade: main.go
	go run main.go migrate down

test:
	go test -count=1 -parallel 1 -v ./...
