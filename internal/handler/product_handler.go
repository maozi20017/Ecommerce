package handler

import (
	"Ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	produceService *service.ProductService
}

func NewProductHandler(produceService *service.ProductService) *ProductHandler {
	return &ProductHandler{produceService: produceService}
}

func (h *ProductHandler) GetProduct(c *gin.Context) {

}
