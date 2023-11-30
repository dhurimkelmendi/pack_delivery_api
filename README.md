# pack_delivery_api
## This is an API that uses the packages below to create an API:
-	HTTP Router: https://github.com/go-chi/chi
-	CORS checks: https://github.com/go-chi/cors
-	Request id generation(KSUID): https://github.com/segmentio/ksuid
-	Logging: https://github.com/sirupsen/logrus

## Running the API
To be able to use postgres locally without SSL, comment the following lines of code in db/db.go:
```
	TLSConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
```
This disables ssl connection with the postgres instance locally. It is used in the deployed version of the app to ensure communication with the DB is encrypted.

There are two options to run the API:
### Using go cli commands
1. Make sure you have Postgres running
2. Add your env variables(check config/config.go) for db credentials
3. Run `make migrate`(this will run migrations and then run the server)
4. Go to `localhost:8080` to access the API
### Using Docker
1. Run `docker-compose up`. This will create two containers: one for the Postgres db, and another for the API.
2. Go to `localhost:8080` to access the API

## Running the tests
To run tests, run `go test ./controllers ./services`

To see the API up and running, visit https://pack-delivery-api.onrender.com/api/v1/products and make a POST request. Make sure you provide the correct request body so that it passes the API validations.


## Endpoints
/POST `/api/v1/products`
Enables the end user to order products defined by the request body. It will return the response as stated below.

### Request body
```
{
	"amount": int
}
```
### Response body
```
{
	"amount_of_packs": int,
	"pack_size": int
}
```
