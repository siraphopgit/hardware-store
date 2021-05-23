package routes

import (
	"hardware/api/controllers"
	"net/http"
)

type ProductRoutes interface {
	Routes() []*Route
}

type ProductRoutesImpl struct {
	productsController controllers.ProductsController
}

func NewProductRoutes(productsController controllers.ProductsController) *ProductRoutesImpl {
	return &ProductRoutesImpl{productsController}
}

func (r *ProductRoutesImpl) Routes() []*Route {
	return []*Route{
		&Route{
			Path:    "/products",
			Method:  http.MethodPost,
			Handler: r.productsController.PostProduct,
		},
	}
}
