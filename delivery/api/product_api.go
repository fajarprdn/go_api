package api

import (
	"github.com/gin-gonic/gin"
	"go_api/delivery/apprequest"
	"go_api/delivery/commonresp"
	"net/http"
)

type ProductApi struct {
}

func (p *ProductApi) QueryStringUrl(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func (p *ProductApi) gettingWithPathVariable() gin.HandlerFunc {
	return func(c *gin.Context) {
		productId := c.Param("id")
		commonresp.NewJsonResponse(c).SendData(commonresp.NewResponseMessage("00", "Product Id", productId))

	}
}

func (p *ProductApi) posting() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productReq apprequest.ProductRequest
		if err := c.ShouldBindJSON(&productReq); err != nil {
			commonresp.NewJsonResponse(c).SendData(commonresp.NewResponseMessage("00", "Create Product", productReq))
		}
		c.JSON(http.StatusOK, gin.H{
			"message": productReq,
		})
	}
}

func (p *ProductApi) gettingWithQueryParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageNo := c.Query("page")
		itemPerPage := c.Query("itempage")
		commonresp.NewJsonResponse(c).SendData(commonresp.NewResponseMessage("00", "Product Id", gin.H{
			"pageNo":      pageNo,
			"itemPerPage": itemPerPage,
		}))
	}
}

func NewProductApi(producRoute *gin.RouterGroup) {
	api := ProductApi{}
	producRoute.GET("", api.gettingWithPathVariable())
	producRoute.GET("/:id", api.gettingWithPathVariable())
	producRoute.POST("", api.posting())
}
