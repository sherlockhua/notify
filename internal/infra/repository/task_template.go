package repository

import (
	"context"
	"gorm.io/gorm"
	"notify/internal/common"
	"notify/internal/domain/entity"
	"notify/internal/domain/repository"
	"time"
)

type TemplateRepositoryImpl struct {
	db *gorm.DB //+
}

type TemplateModel struct {
	ID           int64             `gorm:"column:id" json:"id"`
	TemplateID   int64             `gorm:"column:template_id" json:"template_id"`
	TemplateName string            `gorm:"column:template_name" json:"template_name"`
	TemplateDesc string            `gorm:"column:template_desc" json:"template_desc"`
	CreateTime   time.Time         `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
	TaskType     int               `gorm:"column:task_type" json:"task_type"`
	NotifyType   int               `gorm:"column:notify_type" json:"notify_type"`
	TaskStatus   common.TaskStatus `gorm:"column:task_status;type:tinyint" json:"task_status"`
	StrategyData string            `gorm:"column:strategy_data" json:"strategy_data"` // 策略数据
}

func (t *TemplateModel) TableName() string {
	return "task_template"
}
func (t *TemplateModel) ToBizModel(ctx context.Context) *entity.TaskTemplate {
	taskTemplate := &entity.TaskTemplate{
		TemplateID:    t.TemplateID,
		TemplateName:  t.TemplateName,
		TaskDesc:      t.TemplateDesc,
		CreateTime:    t.CreateTime,
		TaskType:      common.TaskType(t.TaskType),
		NotifyChannel: common.NotifyChannel(t.NotifyType),
		TaskStatus:    t.TaskStatus,
		StrategyData:  t.StrategyData,
	}
	return taskTemplate
}

func ToTemplateModel(taskTemplate *entity.TaskTemplate) *TemplateModel {
	return &TemplateModel{
		TemplateID:   taskTemplate.TemplateID,
		TemplateName: taskTemplate.TemplateName,
		TemplateDesc: taskTemplate.TaskDesc,
		CreateTime:   taskTemplate.CreateTime,
		TaskType:     int(taskTemplate.TaskType),
		NotifyType:   int(taskTemplate.NotifyChannel),
		TaskStatus:   taskTemplate.TaskStatus,
		StrategyData: taskTemplate.StrategyData,
	}
}

func NewTemplateRepository(db *gorm.DB) repository.TemplateRepository {
	return &TemplateRepositoryImpl{db: db}
}

func (r *TemplateRepositoryImpl) CreateTemplate(ctx context.Context, template *entity.TaskTemplate) error {
	templateModel := ToTemplateModel(template)
	if err := r.db.WithContext(ctx).Create(templateModel).Error; err != nil {
		return err
	}
	return nil
}
func (r *TemplateRepositoryImpl) GetTemplate(ctx context.Context, templateId int64) (*entity.TaskTemplate, error) {
	templateModel := &TemplateModel{}
	if err := r.db.WithContext(ctx).Where("template_id =?", templateId).First(templateModel).Error; err != nil {
		return nil, err
	}
	return templateModel.ToBizModel(ctx), nil
}

func (r *TemplateRepositoryImpl) UpdateTemplate(ctx context.Context, template *entity.TaskTemplate) error {
	templateModel := ToTemplateModel(template)
	if err := r.db.WithContext(ctx).Where("template_id =?", template.TemplateID).Updates(templateModel).Error; err != nil {
		return err
	}
	return nil
}
func (r *TemplateRepositoryImpl) DeleteTemplate(ctx context.Context, templateId int64) error {
	if err := r.db.WithContext(ctx).Where("template_id =?", templateId).Delete(&TemplateModel{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *TemplateRepositoryImpl) GetTemplateList(ctx context.Context, offset, size int32, conditions map[string]interface{}) ([]*entity.TaskTemplate, error) {
	templateModels := make([]*TemplateModel, 0)
	if err := r.db.WithContext(ctx).Where(conditions).Offset(int(offset)).Limit(int(size)).Find(&templateModels).Error; err != nil {
		return nil, err
	}
	templates := make([]*entity.TaskTemplate, 0, len(templateModels))
	for _, templateModel := range templateModels {
		templates = append(templates, templateModel.ToBizModel(ctx))
	}
	return templates, nil
}
