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

	t.Run("create pack sizes", func(t *testing.T) {
		createPackSizesPayload := payloads.ChangePackSizesPayload{
			Sizes: []int{23, 31, 53},
		}
		_, err := service.CreatePackSizes(ctx, &createPackSizesPayload)
		if err != nil {
			t.Fatalf("failed to insert pack_sizes: %v", err)
		}
	})
	t.Run("create product orders with 250, 500, 2000, 5000 pack_sizes", func(t *testing.T) {
		createPackSizesPayload := payloads.ChangePackSizesPayload{
			Sizes: []int{250, 500, 2000, 5000},
		}
		_, err := service.CreatePackSizes(ctx, &createPackSizesPayload)
		if err != nil {
			t.Fatalf("failed to insert pack_sizes: %v", err)
		}
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
			if len(expectedResult) != len(createdProductOrder) {
				t.Fatalf("create product order failed to produce expected result. \n received: %v", createdProductOrder)
			}
			if !reflect.DeepEqual(expectedResult, createdProductOrder) {
				t.Fatalf("create product order failed to produce expected result. \n received: %v", createdProductOrder)
			}
		})
	})
	t.Run("create product orders with 23, 31, 53 pack_sizes", func(t *testing.T) {
		createPackSizesPayload := payloads.ChangePackSizesPayload{
			Sizes: []int{23, 31, 53},
		}
		_, err := service.CreatePackSizes(ctx, &createPackSizesPayload)
		if err != nil {
			t.Fatalf("failed to insert pack_sizes: %v", err)
		}
		t.Run("create product order with amount 500k", func(t *testing.T) {

			expectedResult := make([]*payloads.CreatedPackOrderPayload, 0, 1)

			expectedResult = append(expectedResult, &payloads.CreatedPackOrderPayload{AmountOfPacks: 9434, PackSize: 53})

			productOrderToCreate := &payloads.CreateProductOrderPayload{}
			productOrderToCreate.Amount = 500000
			createdProductOrder, err := service.CreateProductOrder(ctx, productOrderToCreate)
			if err != nil {
				t.Fatalf("error while creating product %+v", err)
			}
			if !reflect.DeepEqual(expectedResult, createdProductOrder) {
				t.Fatalf("create product order failed to produce expected result. \n received: %v", createdProductOrder)
			}
		})
	})
}
