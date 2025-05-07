package repository

import (
	"context"

	"notify/internal/domain/entity"
)

type TemplateRepository interface {
	GetTemplate(ctx context.Context, templateId int64) (*entity.TaskTemplate, error)
	CreateTemplate(ctx context.Context, template *entity.TaskTemplate) error
	UpdateTemplate(ctx context.Context, template *entity.TaskTemplate) error
	DeleteTemplate(ctx context.Context, templateId int64) error
	GetTemplateList(ctx context.Context, offset, size int32, conditions map[string]interface{}) ([]*entity.TaskTemplate, error)
}
