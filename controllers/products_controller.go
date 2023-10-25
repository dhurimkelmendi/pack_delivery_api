package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dhurimkelmendi/pack_delivery_api/api"
	"github.com/dhurimkelmendi/pack_delivery_api/payloads"
	"github.com/dhurimkelmendi/pack_delivery_api/services"
)

// A ProductsController handles HTTP requests that deal with product.
type ProductsController struct {
	Controller
	productService *services.ProductService
}

var productsControllerDefaultInstance *ProductsController

// GetProductsControllerDefaultInstance returns the default instance of ProductController.
func GetProductsControllerDefaultInstance() *ProductsController {
	if productsControllerDefaultInstance == nil {
		productsControllerDefaultInstance = NewProductController(services.GetProductServiceDefaultInstance())
	}

	return productsControllerDefaultInstance
}

// NewProductController create a new instance of a product controller using the supplied services
func NewProductController(productService *services.ProductService) *ProductsController {
	controller := Controller{
		errCmp:    api.NewErrorComponent(api.CmpController),
		responder: api.GetResponderDefaultInstance(),
	}

	return &ProductsController{
		Controller:     controller,
		productService: productService,
	}
}

// CreateProductOrder creates a new product and returns product details with an authentication token
func (c *ProductsController) CreateProductOrder(w http.ResponseWriter, r *http.Request) {
	errCtx := c.errCmp(api.CtxCreateProduct, r.Header.Get("X-Request-Id"))
	product := &payloads.CreateProductOrderPayload{}
	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		c.responder.Error(w, errCtx(api.ErrCreatePayload, errors.New("cannot decode product")), http.StatusBadRequest)
		return
	}

	if err := product.Validate(); err != nil {
		c.responder.Error(w, errCtx(api.ErrInvalidRequestPayload, errors.New("request body not valid, missing required fields")), http.StatusBadRequest)
		return
	}

	createdProduct, err := c.productService.CreateProductOrder(context.Background(), product)
	if err != nil {
		c.responder.Error(w, errCtx(api.ErrCreateProductOrder, err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	c.responder.JSON(w, r, createdProduct, http.StatusCreated)
}
