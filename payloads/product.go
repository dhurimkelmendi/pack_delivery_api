package payloads

import (
	"fmt"
	"net/http"

	"github.com/dhurimkelmendi/pack_delivery_api/models"
)

// CreateProductOrderPayload for registering a new product
type CreateProductOrderPayload struct {
	Amount int `json:"amount"`
}

// Validate ensures that all the required fields are present in an instance of *RegisterProductPayload
func (p *CreateProductOrderPayload) Validate() error {
	if p.Amount == 0 {
		return fmt.Errorf("amount is a required field")
	}
	return nil
}

// Render is used by go-chi/renderer
func (p *CreateProductOrderPayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// AvailablePackSizes shows the available packs
// var AvailablePackSizes = []int32{250, 500, 1000, 2000, 5000}

// CreatedPackOrderPayload for returning the created pack order
type CreatedPackOrderPayload struct {
	AmountOfPacks int `json:"amount_of_packs"`
	PackSize      int `json:"pack_size"`
}

// Render is used by go-chi/renderer
func (p *CreatedPackOrderPayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Equals compares two instances of type CreatedPackOrderPayload
func (p *CreatedPackOrderPayload) Equals(secondPackOrder *CreatedPackOrderPayload) bool {
	if p.AmountOfPacks != secondPackOrder.AmountOfPacks {
		return false
	}
	if p.PackSize != secondPackOrder.PackSize {
		return false
	}
	return true
}

// ChangePackSizesPayload is a struct that represents the payload that is expected when creating packSizes
type ChangePackSizesPayload struct {
	Sizes []int `json:"sizes"`
}

func (a *ChangePackSizesPayload) ToPackSizesModelArray() []*models.PackSize {
	packSizesModelArray := make([]*models.PackSize, 0, len(a.Sizes))
	for _, packSize := range a.Sizes {
		packSizeModel := &models.PackSize{
			Size: packSize,
		}
		packSizesModelArray = append(packSizesModelArray, packSizeModel)
	}
	return packSizesModelArray
}

// Validate ensures that all the required fields are present in an instance of *CreatePackSizesPayload
func (a *ChangePackSizesPayload) Validate() error {
	if len(a.Sizes) == 0 {
		return fmt.Errorf("sizes is a required field")
	}

	return nil
}
