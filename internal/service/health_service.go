package service

import (
	"context"

	"gomind/internal/model"
)

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) Check(ctx context.Context) (*model.HealthStatus, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return &model.HealthStatus{
		Service: "gomind",
		Status:  "running",
	}, nil
}
