package services_test

import (
	"context"
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

				productOrderToCreate := &payloads.CreateProductOrderPayload{}
				productOrderToCreate.Amount = 1
				createdProductOrder, err := service.CreateProductOrder(ctx, productOrderToCreate)
				if err != nil {
					t.Fatalf("error while creating product %+v", err)
				}
				if len(createdProductOrder) != 1 {
					t.Fatalf("create product order failed to produce expected result. \n expected: amount_of_packs: %d, pack_size: %d \n received: %+v", 1, 250, createdProductOrder)
				}
			})
			t.Run("create product order with amount 501", func(t *testing.T) {

				expectedResult := make([]*payloads.CreatedPackOrderPayload, 0, 1)
				expectedResult = append(expectedResult, &payloads.CreatedPackOrderPayload{AmountOfPacks: 1, PackSize: 250})

				productOrderToCreate := &payloads.CreateProductOrderPayload{}
				productOrderToCreate.Amount = 501
				createdProductOrder, err := service.CreateProductOrder(ctx, productOrderToCreate)
				if err != nil {
					t.Fatalf("error while creating product %+v", err)
				}
				if len(createdProductOrder) != 2 {
					t.Fatalf("create product order failed to produce expected result. \n received: %+v", createdProductOrder)
				}
			})
			t.Run("create product order with amount 751", func(t *testing.T) {

				expectedResult := make([]*payloads.CreatedPackOrderPayload, 0, 1)
				expectedResult = append(expectedResult, &payloads.CreatedPackOrderPayload{AmountOfPacks: 1, PackSize: 250})

				productOrderToCreate := &payloads.CreateProductOrderPayload{}
				productOrderToCreate.Amount = 751
				createdProductOrder, err := service.CreateProductOrder(ctx, productOrderToCreate)
				if err != nil {
					t.Fatalf("error while creating product %+v", err)
				}
				if len(createdProductOrder) != 2 {
					t.Fatalf("create product order failed to produce expected result. \n received: %+v", createdProductOrder)
				}
			})
			t.Run("create product order with amount 12001", func(t *testing.T) {

				expectedResult := make([]*payloads.CreatedPackOrderPayload, 0, 1)
				expectedResult = append(expectedResult, &payloads.CreatedPackOrderPayload{AmountOfPacks: 1, PackSize: 250})

				productOrderToCreate := &payloads.CreateProductOrderPayload{}
				productOrderToCreate.Amount = 12001
				createdProductOrder, err := service.CreateProductOrder(ctx, productOrderToCreate)
				if err != nil {
					t.Fatalf("error while creating product %+v", err)
				}
				if len(createdProductOrder) != 3 {
					t.Fatalf("create product order failed to produce expected result. \n received: %+v", createdProductOrder)
				}
			})
		})
	})
}
