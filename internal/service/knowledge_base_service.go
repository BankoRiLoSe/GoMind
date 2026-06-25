package service

import (
	"context"
	"fmt"
	"strings"

	"gomind/internal/dao"
	"gomind/internal/dto"
	"gomind/internal/model"

	"github.com/google/uuid"
)

type KnowledgeBaseService struct {
	knowledgeBaseDao *dao.KnowledgeBaseDao
}

func NewKnowledgeBaseService(knowledgeBaseDao *dao.KnowledgeBaseDao) *KnowledgeBaseService {
	return &KnowledgeBaseService{knowledgeBaseDao: knowledgeBaseDao}
}

func (s *KnowledgeBaseService) Create(ctx context.Context, req dto.CreateKnowledgeBaseRequest) (*dto.KnowledgeBaseResponse, error) {
	userID := strings.TrimSpace(req.UserID)
	if userID == "" {
		return nil, fmt.Errorf("user_id is required")
	}

	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, fmt.Errorf("knowledge base name is required")
	}

	knowledgeBase := &model.KnowledgeBase{
		UUID:        uuid.NewString(),
		UserID:      userID,
		Name:        name,
		Description: strings.TrimSpace(req.Description),
	}
	if err := s.knowledgeBaseDao.Create(ctx, knowledgeBase); err != nil {
		return nil, err
	}

	return &dto.KnowledgeBaseResponse{
		KnowledgeBaseID: knowledgeBase.UUID,
		UserID:          knowledgeBase.UserID,
		Name:            knowledgeBase.Name,
		Description:     knowledgeBase.Description,
	}, nil
}

func (s *KnowledgeBaseService) List(ctx context.Context, userID string) ([]dto.KnowledgeBaseResponse, error) {
	userID = strings.TrimSpace(userID)
	if userID == "" {
		return nil, fmt.Errorf("user_id is required")
	}

	knowledgeBases, err := s.knowledgeBaseDao.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.KnowledgeBaseResponse, 0, len(knowledgeBases))
	for _, knowledgeBase := range knowledgeBases {
		responses = append(responses, dto.KnowledgeBaseResponse{
			KnowledgeBaseID: knowledgeBase.UUID,
			UserID:          knowledgeBase.UserID,
			Name:            knowledgeBase.Name,
			Description:     knowledgeBase.Description,
		})
	}

	return responses, nil
}
