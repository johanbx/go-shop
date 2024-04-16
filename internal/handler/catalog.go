package handler

import (
	"johanbx/go-web-server-2/internal/catalog"
	"johanbx/go-web-server-2/internal/db"
	"log"

	"github.com/gin-gonic/gin"
)

type CatalogItemRequestModel struct {
	Name string `json:"name" form:"name" validate:"required,min=2,max=100"`
}

func PostCreateCatalogItem(c *gin.Context) {
	var requestModel CatalogItemRequestModel
	if err := ShouldPayloadBind(c, &requestModel); err != nil {
		return
	}

	log.Printf("requestModel: %+v", requestModel)
	if err := ShouldValidate(c, requestModel); err != nil {
		return
	}

	databaseModel := catalog.CatalogItem{
		Name: requestModel.Name,
	}
	repo := catalog.NewCatalogRepository(db.DB)
	repo.Create(&databaseModel)

	c.JSON(200, databaseModel)
}

func GetAllCatalogItems(c *gin.Context) {
	result := map[string]interface{}{}
	db.DB.Model(&catalog.CatalogItem{}).Find(&result)
	c.JSON(200, result)
}
