package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go-ms-template/src/application/services"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

type OrderHandler interface {
	CreateOrder(ctx *gin.Context)
	GetOrderByID(ctx *gin.Context)
}

type orderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) OrderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}

func (c *orderHandler) CreateOrder(ctx *gin.Context) {
	var orderDTO dto.CreateOrderDTO

	if err := ctx.ShouldBindJSON(&orderDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	result, err := c.orderService.CreateOrder(orderDTO)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error while save data"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *orderHandler) GetOrderByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	order, err := c.orderService.GetOrderById(idParam)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	ctx.JSON(http.StatusOK, order)
}
