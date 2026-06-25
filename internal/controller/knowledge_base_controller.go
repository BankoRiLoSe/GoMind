package controller

import (
	"net/http"

	"gomind/internal/dto"
	"gomind/internal/service"

	"github.com/gin-gonic/gin"
)

type KnowledgeBaseController struct {
	knowledgeBaseService *service.KnowledgeBaseService
}

func NewKnowledgeBaseController(knowledgeBaseService *service.KnowledgeBaseService) *KnowledgeBaseController {
	return &KnowledgeBaseController{knowledgeBaseService: knowledgeBaseService}
}

func (c *KnowledgeBaseController) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/api/v1/knowledge-bases")
	group.POST("", c.Create)
	group.GET("", c.List)
}

func (c *KnowledgeBaseController) Create(ctx *gin.Context) {
	var req dto.CreateKnowledgeBaseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Error(err.Error()))
		return
	}

	knowledgeBase, err := c.knowledgeBaseService.Create(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dto.Success(knowledgeBase))
}

func (c *KnowledgeBaseController) List(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	knowledgeBases, err := c.knowledgeBaseService.List(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, dto.Success(knowledgeBases))
}
