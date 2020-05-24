package models

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

type Evidence struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	AttachmentId string    `json:"attachment_id"`
	FlagId       uuid.UUID `json:"flag_id"`
	File         os.File   `gorm:"-" json:"file"`
	Type         string    `json:"type"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func CreateEvidence(attachmentId, flagId uuid.UUID, type_ string) bool {

	db.Create(&Evidence{
		AttachmentId: attachmentId.String(),
		FlagId:       flagId,
		Type:         type_,
	})
	return true
}

// 返回自昨天开始的evidence
func FindEvidencesByFlag(flagId uuid.UUID) (evidences []Evidence) {
	// 获取当前时间
	now := time.Now()
	// 回到昨天
	yesterday := now.Add(time.Hour * -24)
	yesterday = time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, yesterday.Location())
	db.Where("flag_id = ? and created_at >= ?", flagId, yesterday).Order("created_at desc").Find(&evidences)
	return
}

func FindEvidenceByFlagIdAndAttachmentId(flagId, attachmentId uuid.UUID) (evidences []Evidence) {
	db.Where("flag_id = ? and attachment_id = ?", flagId, attachmentId).Find(&evidences)
	return
}

// BeforeCreate will set a UUID rather than numeric ID.
func (e *Evidence) BeforeCreate(scope *gorm.Scope) error {
	uuid_, _ := uuid.NewV4()
	scope.SetColumn("ID", uuid_)
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

func (e *Evidence) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
