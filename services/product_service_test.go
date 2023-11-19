package services_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/dhurimkelmendi/pack_delivery_api/payloads"
	"github.com/dhurimkelmendi/pack_delivery_api/services"
)

func TestProductService(t *testing.T) {
	t.Parallel()

	service := services.GetProductServiceDefaultInstance()
	ctx := context.Background()

	t.Run("create product order", func(t *testing.T) {
		t.Run("create product order with all fields", func(t *testing.T) {
			t.Run("create product order with amount 1", func(t *testing.T) {
				expectedResult := make([]*payloads.CreatedPackOrderPayload, 0, 1)
				expectedResult = append(expectedResult, &payloads.CreatedPackOrderPayload{AmountOfPacks: 1, PackSize: 250})

				productOrderToCreate := &payloads.CreateProductOrderPayload{}
				productOrderToCreate.Amount = 1
				createdProductOrder, err := service.CreateProductOrder(ctx, productOrderToCreate)
				if err != nil {
					t.Fatalf("error while creating product %+v", err)
				}
				if !reflect.DeepEqual(expectedResult, createdProductOrder) {
					t.Fatalf("create product order failed to produce expected result. \n received: %v", createdProductOrder)
				}
			})
			t.Run("create product order with amount 251", func(t *testing.T) {

				expectedResult := make([]*payloads.CreatedPackOrderPayload, 0, 1)
				expectedResult = append(expectedResult, &payloads.CreatedPackOrderPayload{AmountOfPacks: 1, PackSize: 500})

				productOrderToCreate := &payloads.CreateProductOrderPayload{}
				productOrderToCreate.Amount = 251
				createdProductOrder, err := service.CreateProductOrder(ctx, productOrderToCreate)
				if err != nil {
					t.Fatalf("error while creating product %+v", err)
				}
				if !reflect.DeepEqual(expectedResult, createdProductOrder) {
					t.Fatalf("create product order failed to produce expected result. \n received: %v", createdProductOrder)
				}
			})
			t.Run("create product order with amount 501", func(t *testing.T) {

				expectedResult := make([]*payloads.CreatedPackOrderPayload, 0, 1)
				expectedResult = append(expectedResult, &payloads.CreatedPackOrderPayload{AmountOfPacks: 1, PackSize: 500})
				expectedResult = append(expectedResult, &payloads.CreatedPackOrderPayload{AmountOfPacks: 1, PackSize: 250})

				productOrderToCreate := &payloads.CreateProductOrderPayload{}
				productOrderToCreate.Amount = 501
				createdProductOrder, err := service.CreateProductOrder(ctx, productOrderToCreate)
				if err != nil {
					t.Fatalf("error while creating product %+v", err)
				}
				if !reflect.DeepEqual(expectedResult, createdProductOrder) {
					t.Fatalf("create product order failed to produce expected result. \n received: %v", createdProductOrder)
				}
			})
			t.Run("create product order with amount 751", func(t *testing.T) {

				expectedResult := make([]*payloads.CreatedPackOrderPayload, 0, 1)
				expectedResult = append(expectedResult, &payloads.CreatedPackOrderPayload{AmountOfPacks: 2, PackSize: 500})

				productOrderToCreate := &payloads.CreateProductOrderPayload{}
				productOrderToCreate.Amount = 751
				createdProductOrder, err := service.CreateProductOrder(ctx, productOrderToCreate)
				if err != nil {
					t.Fatalf("error while creating product %+v", err)
				}
				if !reflect.DeepEqual(expectedResult, createdProductOrder) {
					t.Fatalf("create product order failed to produce expected result. \n received: %v", createdProductOrder)
				}
			})
			t.Run("create product order with amount 12001", func(t *testing.T) {

				expectedResult := make([]*payloads.CreatedPackOrderPayload, 0, 1)

				expectedResult = append(expectedResult, &payloads.CreatedPackOrderPayload{AmountOfPacks: 2, PackSize: 5000})
				expectedResult = append(expectedResult, &payloads.CreatedPackOrderPayload{AmountOfPacks: 1, PackSize: 2000})
				expectedResult = append(expectedResult, &payloads.CreatedPackOrderPayload{AmountOfPacks: 1, PackSize: 250})

				productOrderToCreate := &payloads.CreateProductOrderPayload{}
				productOrderToCreate.Amount = 12001
				createdProductOrder, err := service.CreateProductOrder(ctx, productOrderToCreate)
				if err != nil {
					t.Fatalf("error while creating product %+v", err)
				}
				if !reflect.DeepEqual(expectedResult, createdProductOrder) {
					t.Fatalf("create product order failed to produce expected result. \n received: %v", createdProductOrder)
				}
			})
		})
	})
}
