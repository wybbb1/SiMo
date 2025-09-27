package repo

import (
	"gorm.io/gorm"
	
	"github.com/wybbb1/SiMo/internal/log"
)

type LoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *LoginRepository {
	return &LoginRepository{db: db}
}

func (r *LoginRepository) GetUserByUsername(username string) *User {
	var user User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		log.Logger.Error("GetUserByUsername failed",
			zap.String("error", result.Error.Error()),
			zap.String("username", username),
		)
		return nil
	}
	return &user
}