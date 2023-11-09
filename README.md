# pack_delivery_api
## This is an API that uses the packages below to create an API:
-	HTTP Router: https://github.com/go-chi/chi
-	CORS checks: https://github.com/go-chi/cors
-	Request id generation(KSUID): https://github.com/segmentio/ksuid
-	Logging: https://github.com/sirupsen/logrus

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
To run tests, run `go test ./controllers ./services`

To see the API up and running, visit https://pack-delivery-api.onrender.com/api/v1/products and make a POST request. Make sure you provide the correct request body so that it passes the API validations.