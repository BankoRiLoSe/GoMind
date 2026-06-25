package dto

type CreateKnowledgeBaseRequest struct {
	UserID      string `json:"user_id" binding:"required"`
	Name        string `json:"name" binding:"required,min=1,max=128"`
	Description string `json:"description" binding:"max=512"`
}

type KnowledgeBaseResponse struct {
	KnowledgeBaseID string `json:"knowledge_base_id"`
	UserID          string `json:"user_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
}
