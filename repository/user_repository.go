package repository

import (
	"final-project/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Add(newUser model.User) error {
	tx := ur.db.Create(&newUser)
	return tx.Error
}

func (ur *UserRepository) GetByUsername(username string) (model.User, error) {
	var user model.User

	tx := ur.db.First(&user, "username = ?", username)
	return user, tx.Error
}
