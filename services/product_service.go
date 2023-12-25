package services

import (
	"context"

	"github.com/dhurimkelmendi/pack_delivery_api/db"
	"github.com/dhurimkelmendi/pack_delivery_api/models"
	"github.com/dhurimkelmendi/pack_delivery_api/payloads"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
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
		logrus.Error(err)
		return packSizes, err
	}

	_, err = dbSession.Model(&packSizes).Insert()
	if err != nil {
		logrus.Error(err)
		return packSizes, err
	}
	return packSizes, nil
}

// CreateProductOrder creates a product order using the provided payload
func (s *ProductService) CreateProductOrder(ctx context.Context, createProduct *payloads.CreateProductOrderPayload) ([]*payloads.CreatedPackOrderPayload, error) {
	availablePackSizesFromDB := make([]*models.PackSize, 0)
	err := s.db.Model(&availablePackSizesFromDB).Select()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if len(availablePackSizesFromDB) == 0 {
		return nil, pg.ErrNoRows
	}
	availablePackSizes := make([]int, 0, len(availablePackSizesFromDB))
	for _, packSize := range availablePackSizesFromDB {
		availablePackSizes = append(availablePackSizes, packSize.Size)
	}
	packSplitter := NewPackSplitter(availablePackSizes, createProduct.Amount)
	createdPackOrdersMap := packSplitter.SplitOrderIntoPacks()

	createdPackOrders := make([]*payloads.CreatedPackOrderPayload, 0, len(availablePackSizes))

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
