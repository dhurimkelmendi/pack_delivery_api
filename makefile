#This is only meant to be used in development. DO NOT USE THIS TO SERVE PRODUCTION SERVER
serve: main.go
	go run main.go

migrate: main.go
	go run main.go migrate up

downgrade: main.go
	go run main.go migrate down

reset: main.go
	go run main.go migrate reset

test:
	go test -count=1 -parallel 1 -v ./...

swagger:
	swag f -g server.go -d server,controllers,payloads
	swag i -g server.go -d server,controllers,payloads
