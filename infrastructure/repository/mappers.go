package repository

import (
	"github.com/Sina-Mahmoodmoradi/blog/infrastructure/repository/models"
	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
)





func ToEntity(u *models.User) *entity.User{
	if u==nil{
		return nil
	}

	return &entity.User{
		ID: u.ID,
		Username: u.Username,
		Email: u.Email,
		PasswordHash: u.PasswordHash,
	}
}

func ToModel(u *entity.User) *models.User{
	if u==nil{
		return nil
	}

	return &models.User{
		ID: u.ID,
		Username: u.Username,
		Email: u.Email,
		PasswordHash: u.PasswordHash,	
	}
}