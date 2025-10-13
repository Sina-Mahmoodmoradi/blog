package handler

import (
	"net/http"


	"github.com/Sina-Mahmoodmoradi/blog/internal/delivery/http/middleware"
	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"github.com/gin-gonic/gin"
)



type postHandler struct{
	useCase *usecase.PostUseCase
	tokenManager usecase.TokenManager
}

type CreatePostRequest struct{
	Title string `json:"title"`
	Content string `json:"content"`
	
}

type CreatePostResponse struct{
	ID uint `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

func NewPostHandler(useCase *usecase.PostUseCase,tokenManager usecase.TokenManager) *postHandler{
	return &postHandler{
		useCase: useCase,
		tokenManager: tokenManager,
	}
}

func (h *postHandler)RegisterRoutes(r *gin.Engine){
	auth := r.Group("/posts")
	auth.Use(middleware.AuthMiddleware(h.tokenManager))
	{
		auth.POST("/",h.Create)
	}
}


func (h *postHandler)Create(c *gin.Context){
	var req CreatePostRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}

	userID,ok := c.Get("userID")
	if !ok{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized"})
	}


	ctx := c.Request.Context()
	ucResponse,err := h.useCase.CreatePost(ctx,&usecase.CreatePostRequest{
		Title:   req.Title,
		Content: req.Content,
		AuthorID:  userID.(uint),
	})
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	}

	c.JSON(http.StatusCreated,&CreatePostResponse{
		ID: ucResponse.ID,
		Title: ucResponse.Title,
		Content: ucResponse.Content,
	})


}