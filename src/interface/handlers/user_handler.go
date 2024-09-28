package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go-ms-template/src/application/service"
	"github.com/nanda03dev/go-ms-template/src/interface/dto"
)

type UserHandler interface {
	CreateUser(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (c *userHandler) CreateUser(ctx *gin.Context) {
	var userDTO dto.CreateUserDTO

	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result, err := c.userService.CreateUser(userDTO)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *userHandler) GetUserByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	user, err := c.userService.GetUserById(idParam)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
