package models

import "github.com/Sina-Mahmoodmoradi/blog/internal/entity"





func ToEntity(u *User) *entity.User{
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

func ToModel(u *entity.User) *User{
	if u==nil{
		return nil
	}

	return &User{
		ID: u.ID,
		Username: u.Username,
		Email: u.Email,
		PasswordHash: u.PasswordHash,	
	}
}