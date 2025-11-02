package repository

import (
	"github.com/Sina-Mahmoodmoradi/blog/infrastructure/repository/models"
	"github.com/Sina-Mahmoodmoradi/blog/internal/entity"
)





func ToEntityUser(u *models.User) *entity.User{
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

func ToModelUser(u *entity.User) *models.User{
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


func ToEntityPost(p *models.Post) *entity.Post{
	if p==nil{
		return nil
	}

	posts := &entity.Post{
		ID: p.ID,
		Title: p.Title,
		Content:p.Content,
		AuthorID: p.AuthorID,
		Comments: make([]entity.Comment, len(p.Comments)),
	}
	for i, comment := range p.Comments {
		posts.Comments[i] = *ToEntityComment(&comment)
	}
	return posts
}


func ToModelPost(p *entity.Post) *models.Post{
	if p==nil{
		return nil
	}

	posts := &models.Post{
		ID: p.ID,
		Title: p.Title,
		Content:p.Content,
		AuthorID: p.AuthorID,	
		Comments: make([]models.Comment, len(p.Comments)),
	}
	for i, comment := range p.Comments {
		posts.Comments[i] = *ToModelComment(&comment)
	}
	return posts
}


func ToEntityComment(c *models.Comment) *entity.Comment{
	if c==nil{
		return nil
	}

	return &entity.Comment{
		ID: c.ID,
		Content: c.Content,
		AuthorID: c.AuthorID,
		PostID: c.PostID,
	}
}



func ToModelComment(c *entity.Comment) *models.Comment{
	if c==nil{
		return nil
	}

	return &models.Comment{
		ID: c.ID,
		Content: c.Content,
		AuthorID: c.AuthorID,
		PostID: c.PostID,
	}
}