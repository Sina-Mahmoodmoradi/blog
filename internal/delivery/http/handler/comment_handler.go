package handler

import (
	"net/http"
	"strconv"

	"github.com/Sina-Mahmoodmoradi/blog/internal/delivery/http/handler/dto"
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
type UpdateCommentRequest = CreateCommentRequest

type CommentResponse struct{
	ID uint `json:"id"`
	Content string `json:"content"`
	AuthorID uint `json:"author_id"`
}

type GetCommentsResponse struct{
	Comments []*CommentResponse  `json:"posts"`
	Total int 			     `json:"total"`
	Page int 			     `json:"Page"`
	Limit int 			     `json:"limit"`
}


func NewCommentHandler(useCase *usecase.CommentUseCase,tokenManager usecase.TokenManager) *commentHandler{
	return &commentHandler{
		useCase: useCase,
		tokenManager: tokenManager,
	}
}

func (h *commentHandler)RegisterRoutes(r *gin.Engine){
	auth := r.Group("/posts/:postId/comments")
	auth.Use(middleware.AuthMiddleware(h.tokenManager))
	{
		auth.POST("/",h.Create)
		auth.GET("/",h.GetComments)
		// auth.GET("/:id",h.GetCommentById)
		auth.PATCH("/:id",h.Update)
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
	idStr := c.Param("postId")
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
	c.JSON(http.StatusOK,&CommentResponse{
		ID: comment.ID,
		AuthorID: comment.AuthorID,
		Content: comment.Content,
	})
}


func (h *commentHandler)GetComments(c *gin.Context){
	q := dto.PaginationQuery{
		Page: 1,
		Limit: 10,
	}

	if err:=c.ShouldBindQuery(&q);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	idStr := c.Param("postId")
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
	paginatedComments,err := h.useCase.GetAllComments(ctx,&usecase.GetCommentsRequest{
		PostID: postID,
		Page: q.Page,
		Limit: q.Limit,
	})
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	comments := make([]*CommentResponse,0,len(paginatedComments.Comments))
	for _,comment:=range paginatedComments.Comments{
		comments = append(comments, &CommentResponse{
			ID: comment.ID,
			Content: comment.Content,
			AuthorID: comment.AuthorID,
		})
	}

	c.JSON(http.StatusOK,&GetCommentsResponse{
		Comments: comments,
		Total: paginatedComments.Total,
		Page: paginatedComments.Page,
		Limit: paginatedComments.Limit,
	})
	
}

// منطق مربوط به گرفتن یک کامنت خاص به نظرم زیادی اومد و کامنتش کردم
// func (h *commentHandler)GetCommentById(c *gin.Context){
// 	idStr := c.Param("id")
// 	idInt,err := strconv.Atoi(idStr)
// 	if err!=nil{

// 	}
// 	if idInt<=0{
// 		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid post id"})
// 		return
// 	}
// 	id := uint(idInt)
// }

func (h *commentHandler)Update(c *gin.Context){
	var req UpdateCommentRequest
	if err:= c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}


	idStr := c.Param("id")
	idInt,err := strconv.Atoi(idStr)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid comment id"})
		return
	}
	if idInt<=0{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid comment id"})
		return
	}
	id := uint(idInt)

	userID,ok := c.Get("userID")
		if !ok{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized"})
		return
	}

	ctx := c.Request.Context()

	comment,err := h.useCase.UpdateComment(ctx,userID.(uint),id,req.Content)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,&CommentResponse{
		ID: comment.ID,
		AuthorID: comment.AuthorID,
		Content: comment.Content,
	})

}