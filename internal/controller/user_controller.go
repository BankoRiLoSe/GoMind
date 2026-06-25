package controller

import (
	"net/http"

	"gomind/internal/dto"
	"gomind/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/api/users")
	group.POST("/register", c.Register)
}

func (c *UserController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Error(err.Error()))
		return
	}

	user, err := c.userService.Register(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dto.Success(user))
}
