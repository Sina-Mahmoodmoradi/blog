package handler

import (
	"net/http"
	"strconv"

	"github.com/Sina-Mahmoodmoradi/blog/internal/delivery/http/middleware"
	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"github.com/gin-gonic/gin"
)






type commentHandler struct{
	useCase *usecase.CommentUseCase
	tokenManager usecase.TokenManager
}

type CreateCommentRequest struct{
	Content string `json:"content"`
}

type CreateCommentResponse struct{
	ID uint `json:"id"`
	Content string `json:"content"`
	AuthorID uint `json:"author_id"`
}


func NewCommentHandler(useCase *usecase.CommentUseCase,tokenManager usecase.TokenManager) *commentHandler{
	return &commentHandler{
		useCase: useCase,
		tokenManager: tokenManager,
	}
}

func (h *commentHandler)RegisterRoutes(r *gin.Engine){
	auth := r.Group("/posts/:id/comments")
	auth.Use(middleware.AuthMiddleware(h.tokenManager))
	{
		auth.POST("/",h.Create)
		// auth.GET("/",h.GetPosts)
		// auth.GET("/:id",h.GetPostById)
		// auth.PATCH("/:id",h.Update)
		// auth.DELETE("/:id",h.Delete)
	}
}


func (h *commentHandler)Create(c *gin.Context){
	var req CreateCommentRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}

	userID,ok := c.Get("userID")
	if !ok{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized"})
		return
	}
	idStr := c.Param("id")
	intID,err:= strconv.Atoi(idStr)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid post id"})
		return
	}
	if intID<0{
		c.JSON(http.StatusBadRequest,gin.H{"error":"post id is positive"})
		return
	}
	postID := uint(intID)

	ctx := c.Request.Context()
	comment,err := h.useCase.CreateComment(ctx,&usecase.CreateCommentRequest{
		AuthorID: userID.(uint),
		PostID: postID,
		Content: req.Content,
	})
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,&CreateCommentResponse{
		ID: comment.ID,
		AuthorID: comment.AuthorID,
		Content: comment.Content,
	})
}