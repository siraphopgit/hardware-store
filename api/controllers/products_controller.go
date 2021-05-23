package controllers

import (
	"encoding/json"
	"hardware/api/models"
	"hardware/api/repository"
	"hardware/api/utils"
	"io/ioutil"
	"net/http"
)

type ProductsController interface {
	PostProduct(http.ResponseWriter, *http.Request)
}

type ProductsControllerImpl struct {
	productRepository repository.ProductsReposity
}

func NewProductController(productsRepository repository.ProductsReposity) *ProductsControllerImpl {
	return &ProductsControllerImpl{productsRepository}
}

func (c *ProductsControllerImpl) PostProduct(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	product := &models.Product{}
	err = json.Unmarshal(bytes, product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = product.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	product, err = c.productRepository.Save(product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, product.ID))
	utils.WriteAsJson(w, product)
}
