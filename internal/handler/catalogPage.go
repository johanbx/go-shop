package handler

import (
	"johanbx/go-web-server-2/internal/catalog"
	"johanbx/go-web-server-2/internal/db"
	"johanbx/go-web-server-2/internal/page"
	"log"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

const (
	SessionCreatedCatalogItem = "createdCatalogItem"
)

type CatalogPage struct {
	baseData gin.H
	template string
}

func (cp CatalogPage) BaseData() gin.H {
	return cp.baseData
}

func (cp CatalogPage) Template() string {
	return cp.template
}

func CatalogPageRenderer(r multitemplate.Renderer) {
	r.AddFromFiles("catalog-index", "templates/base.html", "templates/catalog.html")
}

func CatalogItemHandler(c *gin.Context) {
	repo := catalog.NewCatalogRepository(db.DB)

	var catalogItems []catalog.CatalogItem
	var err error

	catalogItems, err = repo.List()
	if err != nil {
		return
	}

	// Create base page
	catalogPage := CatalogPage{
		baseData: gin.H{
			"catalogItems": catalogItems,
		},
		template: "catalog-index",
	}

	// Default GET response
	if c.Request.Method == "GET" {
		session := c.MustGet("session").(*sessions.Session)
		createdCatalogItem, createdCatalogItemExists := session.Values[SessionCreatedCatalogItem]

		// If we just created a product
		if createdCatalogItemExists {
			delete(session.Values, SessionCreatedCatalogItem)
			session.Save(c.Request, c.Writer)

			catalogPage.baseData["title"] = "Catalog | Product created"
			catalogPage.baseData["createdCatalogItem"] = createdCatalogItem
			catalogPage.baseData["catalogItems"] = catalogItems
			page.RenderPage(c, http.StatusOK, catalogPage)
			return
		}

		// Default
		catalogPage.baseData["title"] = "Catalog | Create new product"
		page.RenderPage(c, http.StatusOK, catalogPage)
		return
	}

	// Parsing payload
	var payload struct {
		Name string `json:"name" form:"name" validate:"required,min=2,max=100"`
	}
	if err := c.ShouldBind(&payload); err != nil {
		catalogPage.baseData["title"] = "Catalog | Invalid request"
		catalogPage.baseData["errorMessage"] = "Invalid request data"
		page.RenderPage(c, http.StatusBadRequest, catalogPage)
		return
	}

	// Validation
	if validationError := ShouldValidate(c, payload); validationError != nil {
		catalogPage.baseData["title"] = "Catalog | Failed to create product"
		catalogPage.baseData["validationError"] = validationError
		page.RenderPage(c, http.StatusBadRequest, catalogPage)
		return
	}

	// Create item
	catalogItem := catalog.CatalogItem{
		Name: payload.Name,
	}
	repo.Create(&catalogItem)

	session := c.MustGet("session").(*sessions.Session)
	session.Values[SessionCreatedCatalogItem] = catalogItem
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		log.Printf("Error when saving createdCatalogItem in to session: %+v", err.Error())

		catalogPage.baseData["errorMessage"] = "internal server error"
		page.RenderPage(c, http.StatusInternalServerError, catalogPage)
		return
	}

	c.Redirect(http.StatusSeeOther, "/catalog/item")
}
