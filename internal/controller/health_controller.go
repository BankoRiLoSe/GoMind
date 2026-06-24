package controller

import (
	"net/http"

	"gomind/internal/service"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	healthService *service.HealthService
}

func NewHealthController(healthService *service.HealthService) *HealthController {
	return &HealthController{healthService: healthService}
}

func (c *HealthController) RegisterRoutes(router *gin.Engine) {
	router.GET("/health", c.Health)
}

func (c *HealthController) Health(ctx *gin.Context) {
	status, err := c.healthService.Check(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    status,
	})
}
