package repository

import "gorm.io/gorm"

type Repositories struct {
}

func InitRepos(db *gorm.DB) *Repositories {
	return &Repositories{}
}
