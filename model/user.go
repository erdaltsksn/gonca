package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Validate validates the model.
func (u *User) Validate() error {
	if err := validation.ValidateStruct(u,
		validation.Field(&u.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(&u.Password,
			validation.Required,
			validation.Length(8, 64),
		),
	); err != nil {
		return err
	}

	// Everything is fine
	return nil
}

// BeforeCreate validate the model and then generates an `UUID`,
// hashes the password and set them.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if err := u.Validate(); err != nil {
		return err
	}

	var tmpUser User
	if err := tx.Where(&User{Email: u.Email}).First(&tmpUser).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("This email is already registered")
	}

	generatedUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	u.ID = generatedUUID.String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	// Everything is fine
	return nil
}
