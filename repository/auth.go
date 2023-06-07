package repository

import (
    "gorm.io/gorm"
    "github.com/WillyWilsen/Deall-Jobs_Technical-Test.git/model"
)

type AuthRepository interface {
    GetByEmail(email string) (*model.User, error)
    Register(user *model.User) error
}

type AuthRepositoryImpl struct {
    db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
    return &AuthRepositoryImpl{
        db: db,
    }
}

func (r *AuthRepositoryImpl) GetByEmail(email string) (*model.User, error) {
    var user model.User
    if err := r.db.Table("users").Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *AuthRepositoryImpl) Register(user *model.User) error {
    return r.db.Create(user).Error
}

