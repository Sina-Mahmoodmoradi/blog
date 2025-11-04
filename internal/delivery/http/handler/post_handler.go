package handler

import (
	"net/http"
	"strconv"
	"strings"

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



type PostResponse struct{
	ID uint `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Comments []CommentResponse `json:"comments"`
	CommentsTruncated bool `json:"comments_truncated"`
} 

type GetPostsResponse struct{
	Posts []PostResponse `json:"posts"`
	Total int 			 `json:"total"`
	Page int 			 `json:"Page"`
	Limit int 			 `json:"limit"`

}

type GetPostsQuery struct{
	Page int `form:"page" binding:"min=1"`
	Limit int `form:"limit" binding:"min=1,max=100"`
	Tags  string `form:"tags" binding:"omitempty"`

}

type UpdatePostRequest struct{
	Title *string   `json:"title"`
	Content *string `json:"content"`
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
		auth.GET("/",h.GetPosts)
		auth.GET("/:postId",h.GetPostById)
		auth.PATCH("/:postId",h.Update)
		auth.DELETE("/:postId",h.Delete)
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
		return
	}


	ctx := c.Request.Context()
	ucResponse,err := h.useCase.CreatePost(ctx,&usecase.CreatePostRequest{
		Title:   req.Title,
		Content: req.Content,
		AuthorID:  userID.(uint),
	})
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated,&CreatePostResponse{
		ID: ucResponse.ID,
		Title: ucResponse.Title,
		Content: ucResponse.Content,
	})


}


func (h *postHandler)GetPosts(c *gin.Context){
	q := GetPostsQuery{
		Page: 1,
		Limit: 10,
	}

	if err:=c.ShouldBindQuery(&q);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	tagNames := []string{}
	if q.Tags != "" {
		for tag := range strings.SplitSeq(q.Tags, ",") {
			tagNames = append(tagNames, strings.TrimSpace(tag))
		}
	}

	// userID ,ok:=c.Get("userID") 
	// if !ok{
	// 	c.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized"})
	// 	return
	// }
	ctx := c.Request.Context()
	paginatedPosts,err :=h.useCase.GetAllPosts(ctx,&usecase.GetPostsRequest{
		AuthorID: nil,
		Page: q.Page,
		Limit: q.Limit,
		TagNames: tagNames,
	})
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	posts := make([]PostResponse,0,len(paginatedPosts.Posts))
	for _,post:= range paginatedPosts.Posts{
		posts = append(posts, PostResponse{
			ID: post.ID,
			Title: post.Title,
			Content: post.Content,
		})
	}

	c.JSON(http.StatusOK,&GetPostsResponse{
		Posts: posts,
		Total: paginatedPosts.Total,
		Page:paginatedPosts.Page,
		Limit: paginatedPosts.Limit,
	})


}





func (h *postHandler)GetPostById(c *gin.Context){
	idStr := c.Param("postId")
	intID,err:= strconv.Atoi(idStr)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid id"})
		return
	}
	if intID<0{
		c.JSON(http.StatusBadRequest,gin.H{"error":"id is positive"})
		return
	}
	id := uint(intID)

	ctx := c.Request.Context()
	userID ,ok:=c.Get("userID") 
	if !ok{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized"})
		return
	}
	post,hasMoreComments,err := h.useCase.GetPost(ctx,userID.(uint),id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	postResponse := &PostResponse{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		Comments: make([]CommentResponse, len(post.Comments)),
		CommentsTruncated: hasMoreComments,
	}
	for i, comment := range post.Comments {
		postResponse.Comments[i] = CommentResponse{
			ID:      comment.ID,
			Content: comment.Content,
			AuthorID: comment.AuthorID,
		}
	}
	c.JSON(http.StatusOK,postResponse)
}


func (h *postHandler)Update(c *gin.Context){
	var req UpdatePostRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	idStr := c.Param("postId")
	intID,err:= strconv.Atoi(idStr)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid id"})
		return
	}
	if intID<0{
		c.JSON(http.StatusBadRequest,gin.H{"error":"id is positive"})
		return
	}
	id := uint(intID)

	userID ,ok:=c.Get("userID") 
	if !ok{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized"})
		return
	}

	ctx := c.Request.Context()
	post,err:= h.useCase.UpdatePost(ctx,userID.(uint),id,&usecase.UpdatePostRequest{
		Title: req.Title,
		Content: req.Content,
	})

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,&PostResponse{
		ID: post.ID,
		Title: post.Title,
		Content: post.Content,
	})
}



func (h *postHandler)Delete(c *gin.Context){
	idStr := c.Param("postId")
	intID,err:= strconv.Atoi(idStr)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid id"})
		return
	}
	if intID<0{
		c.JSON(http.StatusBadRequest,gin.H{"error":"id is positive"})
		return
	}
	id := uint(intID)

	ctx := c.Request.Context()
	userID ,ok:=c.Get("userID") 
	if !ok{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized"})
		return
	}
	if err:=h.useCase.DeletePost(ctx,userID.(uint),id);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
	
}

