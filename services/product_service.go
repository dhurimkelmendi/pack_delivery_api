package services

import (
	"context"
	"fmt"

	"github.com/dhurimkelmendi/pack_delivery_api/db"
	"github.com/dhurimkelmendi/pack_delivery_api/models"
	"github.com/dhurimkelmendi/pack_delivery_api/payloads"
	"github.com/go-pg/pg/v10"
)

// ProductService is a struct that contains references to the dependencies
type ProductService struct {
	db *pg.DB
}

var productServiceDefaultInstance *ProductService

// GetProductServiceDefaultInstance returns the default instance of ProductService
func GetProductServiceDefaultInstance() *ProductService {
	if productServiceDefaultInstance == nil {
		productServiceDefaultInstance = &ProductService{
			db: db.GetDefaultInstance().GetDB(),
		}
	}

	return productServiceDefaultInstance
}

// CreatePackSizes creates a product order using the provided payload
func (s *ProductService) CreatePackSizes(ctx context.Context, createPackSizes *payloads.ChangePackSizesPayload) ([]*models.PackSize, error) {
	createdPackSizes := make([]*models.PackSize, 0, len(createPackSizes.Sizes))
	var err error
	s.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		createdPackSizes, err = s.createPackSizes(tx, createPackSizes)
		return err
	})
	return createdPackSizes, err
}
func (s *ProductService) createPackSizes(dbSession *pg.Tx, createPackSizes *payloads.ChangePackSizesPayload) ([]*models.PackSize, error) {
	packSizes := make([]*models.PackSize, 0, len(createPackSizes.Sizes))
	if err := createPackSizes.Validate(); err != nil {
		return packSizes, err
	}
	for _, packSize := range createPackSizes.Sizes {
		packSizeModel := &models.PackSize{
			Size: packSize,
		}
		packSizes = append(packSizes, packSizeModel)
	}

	_, err := dbSession.Model(&models.PackSize{}).Where("?", true).Delete()
	if err != nil {
		return packSizes, err
	}

	_, err = dbSession.Model(&packSizes).Insert()
	if err != nil {
		return packSizes, err
	}
	return packSizes, nil
}

// CreateProductOrder creates a product order using the provided payload
func (s *ProductService) CreateProductOrder(ctx context.Context, createProduct *payloads.CreateProductOrderPayload) ([]*payloads.CreatedPackOrderPayload, error) {
	availablePackSizesFromDB := make([]*models.PackSize, 0)
	err := s.db.Model(&availablePackSizesFromDB).Select()
	if err != nil {
		return nil, err
	}
	if len(availablePackSizesFromDB) == 0 {
		return nil, pg.ErrNoRows
	}
	availablePackSizes := make([]int, 0, len(availablePackSizesFromDB))
	for _, packSize := range availablePackSizesFromDB {
		availablePackSizes = append(availablePackSizes, packSize.Size)
	}
	createdPackOrders := make([]*payloads.CreatedPackOrderPayload, 0, len(availablePackSizes))
	createdPackOrdersMap := map[int]int{}
	s.createPackOrdersMapFromProductAmount(createProduct.Amount, availablePackSizes, createdPackOrdersMap)
	s.combineSmallPacksIntoLargerOnesWhereApplicable(availablePackSizes, createdPackOrdersMap)
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
func (s *ProductService) createPackOrdersMapFromProductAmount(createProductAmount int, availablePackSizes []int, createdPackOrdersMap map[int]int) {
	for i := len(availablePackSizes) - 1; i >= 0; i-- {
		packSize := availablePackSizes[i]
		if packSize == 53 && createProductAmount == 51 {
			fmt.Print("OK")
		}
		if createProductAmount >= packSize {
			amountOfPacks, _ := createdPackOrdersMap[packSize]
			createdPackOrdersMap[packSize] = amountOfPacks + 1
			s.createPackOrdersMapFromProductAmount(createProductAmount-packSize, availablePackSizes, createdPackOrdersMap)
			return
		} else if i == 0 && createProductAmount <= packSize {
			createdPackOrdersMap[packSize]++
			return
		}
	}
	return
}

func (s *ProductService) combineSmallPacksIntoLargerOnesWhereApplicable(availablePackSizes []int, createdPackOrdersMap map[int]int) {
	for i := len(availablePackSizes) - 1; i >= 0; i-- {
		packSize := availablePackSizes[i]
		amountOfPacks, _ := createdPackOrdersMap[packSize]
		maxProductAmount := amountOfPacks * int(packSize)
		if amountOfPacks == 0 {
			continue
		}
		sumOfSmallerPackSizes := 0
		proportionsOfSmallerPackSizesToLargerPackSize := map[int]float64{}
		for j := 0; j < i; j++ {
			smallerPackSize := availablePackSizes[j]
			sumOfSmallerPackSizes += smallerPackSize
			proportionsOfSmallerPackSizesToLargerPackSize[smallerPackSize] = float64(smallerPackSize) / float64(packSize)
		}

		portionsOfLargerSizePacksWithAmountsOfPacks := map[float64]int{}
		if sumOfSmallerPackSizes >= packSize {
			for j := 0; j < i; j++ {
				smallerPackSize := availablePackSizes[j]
				proportionForSmallerPackSize := proportionsOfSmallerPackSizesToLargerPackSize[smallerPackSize]
				portionsOfLargerSizePacksWithAmountsOfPacks[proportionForSmallerPackSize*float64(createdPackOrdersMap[smallerPackSize])] = createdPackOrdersMap[smallerPackSize]
				delete(createdPackOrdersMap, smallerPackSize)
			}
			sumOfPortionedPacks := float64(0)
			for portion, amountOfPacks := range portionsOfLargerSizePacksWithAmountsOfPacks {
				sumOfPortionedPacks += portion * float64(amountOfPacks)
			}
			createdPackOrdersMap[packSize] += int(sumOfPortionedPacks)
		}

		for j := i + 1; j <= len(availablePackSizes)-1; j++ {
			largerPackSize := availablePackSizes[j]
			remainder := maxProductAmount % int(largerPackSize)
			if remainder >= 0 && remainder < maxProductAmount {
				differenceInPackSizeForLargerPack := maxProductAmount / int(availablePackSizes[j])
				createdPackOrdersMap[largerPackSize] += differenceInPackSizeForLargerPack
				differenceInPackSizeForSmallerPack := differenceInPackSizeForLargerPack * (maxProductAmount / int(packSize))
				createdPackOrdersMap[packSize] -= differenceInPackSizeForSmallerPack
				if createdPackOrdersMap[packSize] == 0 {
					delete(createdPackOrdersMap, packSize)
				}
			}
		}
	}
	return
}
