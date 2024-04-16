package catalog

import (
	"gorm.io/gorm"
)

type CatalogItem struct {
	gorm.Model `json:"-"`
	Name       string `form:"name" json:"name"`
}

type CatalogRepository struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) *CatalogRepository {
	return &CatalogRepository{db: db}
}

func (repo *CatalogRepository) Create(catalogItem *CatalogItem) error {
	return repo.db.Create(catalogItem).Error
}
