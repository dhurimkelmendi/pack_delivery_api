package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dhurimkelmendi/pack_delivery_api/controllers"
	"github.com/dhurimkelmendi/pack_delivery_api/payloads"
	"github.com/go-chi/chi"
)

func TestProductController(t *testing.T) {
	t.Parallel()

	ctrl := controllers.GetControllersDefaultInstance()

	t.Run("create product", func(t *testing.T) {
		r := chi.NewRouter()
		r.Post("/api/v1/products", ctrl.Products.CreateProductOrder)

		bBuf := bytes.NewBuffer([]byte(fmt.Sprintf(`{"amount": %d}`, 1)))
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", bBuf)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Fatalf("expected http status code of 200 but got: %+v, %+v", res.Code, res.Body.String())
		}

		body := make([]map[string]interface{}, 0, len(payloads.AvailablePackSizes))
		dec := json.NewDecoder(strings.NewReader(res.Body.String()))
		err := dec.Decode(&body)
		if err != nil {
			t.Fatalf("error decoding response body: %+v", err)
		}

		amountOfPacks := body[0]["amount_of_packs"].(float64)
		if int32(1) != int32(amountOfPacks) {
			t.Fatalf("expected amount_of_packs to be 1, got: %+v", amountOfPacks)
		}
		packSize := body[0]["pack_size"].(float64)
		if int32(payloads.AvailablePackSizes[0]) != int32(packSize) {
			t.Fatalf("expected pack_size to be %d, got: %+v", payloads.AvailablePackSizes[0], packSize)
		}
	})
}
