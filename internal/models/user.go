package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Password  string         `json:"password" gorm:"not null"`
	Name      string         `json:"name" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Filter ListUsersParams `json:"-" gorm:"-"`
}

func (i *User) TableName() string {
	return "users"
}

func (i *User) BeforeCreate(tx *gorm.DB) error {
	if i.ID == "" {
		i.ID = uuid.New().String()
	}

	// Hash password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(i.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	i.Password = string(hashedPassword)

	// Set CreatedAt & UpdatedAt explicitly if you want
	now := time.Now()
	i.CreatedAt = now
	i.UpdatedAt = now

	return nil
}

func (i *User) BeforeUpdate(tx *gorm.DB) error {
	i.UpdatedAt = time.Now()

	var old User
	if err := tx.First(&old, "id = ?", i.ID).Error; err != nil {
		return err
	}

	if i.Password != old.Password {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(i.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		i.Password = string(hashedPassword)
	}

	return nil
}

func (i *User) FilterData() ListUsersParams {
	return i.Filter
}

func (i *User) ApplyFilter(filter ListUsersParams) {
	i.Filter = filter
}

type ListUsersParams struct {
	Name     string
	Username string
	Limit    int
	Offset   int
}
