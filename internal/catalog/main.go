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

func (repo *CatalogRepository) List() ([]CatalogItem, error) {
	var items []CatalogItem
	err := repo.db.Model(&CatalogItem{}).Order("created_at DESC").Find(&items).Error

	return items, err
}
