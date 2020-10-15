package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel is a model definition, including fields `ID` as UUID, `CreatedAt`,
// `UpdatedAt`, and `DeletedAt`, which could be embedded in your models.
//
// USAGE:
// 	type User struct {
// 		database.BaseModel
//
// 		Email    string `json:"email" gorm:"UNIQUE;NOT NULL"`
// 		Password string `json:"password" gorm:"NOT NULL"`
// 	}
type BaseModel struct {
	ID string `json:"id" gorm:"PRIMARY_KEY"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

// BeforeSave generates and sets an `UUID` for `ID` column.
func (baseModel *BaseModel) BeforeSave(tx *gorm.DB) error {
	generatedUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	baseModel.ID = generatedUUID.String()

	// Everything is fine
	return nil
}
