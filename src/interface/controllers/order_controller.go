package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go-ms-template/src/application/service"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

type OrderController interface {
	CreateOrder(ctx *gin.Context)
	GetOrderByID(ctx *gin.Context)
}

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &orderController{
		orderService: orderService,
	}
}

func (c *orderController) CreateOrder(ctx *gin.Context) {
	var orderDTO dto.CreateOrderDTO

	if err := ctx.ShouldBindJSON(&orderDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result, err := c.orderService.CreateOrder(orderDTO)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *orderController) GetOrderByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	order, err := c.orderService.GetOrderById(idParam)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	ctx.JSON(http.StatusOK, order)
}
