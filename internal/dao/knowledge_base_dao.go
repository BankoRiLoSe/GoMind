package dao

import (
	"context"
	"fmt"

	"gomind/internal/model"

	"gorm.io/gorm"
)

type KnowledgeBaseDao struct {
	db *gorm.DB
}

func NewKnowledgeBaseDao(db *gorm.DB) *KnowledgeBaseDao {
	return &KnowledgeBaseDao{db: db}
}

func (d *KnowledgeBaseDao) Create(ctx context.Context, knowledgeBase *model.KnowledgeBase) error {
	if err := d.db.WithContext(ctx).Create(knowledgeBase).Error; err != nil {
		return fmt.Errorf("create knowledge base: %w", err)
	}
	return nil
}
