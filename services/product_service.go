package services

import (
	"context"

	"github.com/dhurimkelmendi/pack_delivery_api/payloads"
)

// ProductService is a struct that contains references to the dependencies
type ProductService struct {
}

var productServiceDefaultInstance *ProductService

// GetProductServiceDefaultInstance returns the default instance of ProductService
func GetProductServiceDefaultInstance() *ProductService {
	if productServiceDefaultInstance == nil {
		productServiceDefaultInstance = &ProductService{}
	}

	return productServiceDefaultInstance
}

// CreateProductOrder creates a product order using the provided payload
func (s *ProductService) CreateProductOrder(ctx context.Context, createProduct *payloads.CreateProductOrderPayload) ([]*payloads.CreatedPackOrderPayload, error) {
	createdPackOrders := make([]*payloads.CreatedPackOrderPayload, 0, len(payloads.AvailablePackSizes))
	createdPackOrdersMap := map[int32]int{}
	s.createPackOrdersMapFromProductAmount(createProduct.Amount, createdPackOrdersMap)
	for key, value := range createdPackOrdersMap {
		if value == 0 {
			continue
		}
		createdPackOrderForSize := &payloads.CreatedPackOrderPayload{
			PackSize:      key,
			AmountOfPacks: value,
		}
		createdPackOrders = append(createdPackOrders, createdPackOrderForSize)
	}
	return createdPackOrders, nil
}

// recursive method that generates a map of packSizes used and amount of packs for each size
func (s *ProductService) createPackOrdersMapFromProductAmount(createProductAmount int32, createdPackOrdersMap map[int32]int) {
	for i := len(payloads.AvailablePackSizes) - 1; i >= 0; i-- {
		packSize := payloads.AvailablePackSizes[i]
		if createProductAmount >= packSize {
			amountOfPacks, _ := createdPackOrdersMap[packSize]
			createdPackOrdersMap[packSize] = amountOfPacks + 1
			s.createPackOrdersMapFromProductAmount(createProductAmount-packSize, createdPackOrdersMap)
			return
		} else if i == 0 && createProductAmount <= packSize {
			createdPackOrdersMap[packSize]++
			return
		}
	}
	return
}
