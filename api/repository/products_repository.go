package repository

import (
	"hardware/api/models"

	"github.com/jinzhu/gorm"
)

type ProductsReposity interface {
	Save(*models.Product) (*models.Product, error)
}

type productsReposityImpl struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *productsReposityImpl {
	return &productsReposityImpl{db}
}

func (r *productsReposityImpl) Save(product *models.Product) (*models.Product, error) {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Product{}).Create(product).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return product, tx.Commit().Error
}
