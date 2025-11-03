package handler

import (
	"net/http"
	"strconv"

	"github.com/Sina-Mahmoodmoradi/blog/internal/delivery/http/middleware"
	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"github.com/gin-gonic/gin"
)




type tagHandler struct{
	useCase *usecase.TagUseCase
	tokenManager usecase.TokenManager
}

type AssignTagRequest struct{
	Names []string `json:"names"`
}

type TagResponse struct{
	ID uint `json:"id"`
	Name string `json:"name"`
}

type AssignTagResponse struct{
	Tags []TagResponse `json:"tags"`
}

func NewTagHandler(useCase *usecase.TagUseCase, tokenManager usecase.TokenManager)*tagHandler{
	return &tagHandler{
		useCase: useCase,
		tokenManager: tokenManager,
	}
}


func (h *tagHandler)RegisterRoutes(r *gin.Engine){
	auth := r.Group("/posts/:postId/tags")
	auth.Use(middleware.AuthMiddleware(h.tokenManager))
	{
		auth.POST("/",h.AssignTags)
		// auth.GET("/",h.GetComments)
		// auth.GET("/:id",h.GetCommentById)
		// auth.PATCH("/:id",h.Update)
		// auth.DELETE("/:id",h.Delete)
	}

}


func (h *tagHandler)AssignTags(c *gin.Context){
	var req AssignTagRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	userID,ok := c.Get("userID")
	if !ok{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized"})
		return
	}

	postIDStr := c.Param("postId")
	intPostID,err:= strconv.Atoi(postIDStr)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid postId"})
		return
	}
	if intPostID<0{
		c.JSON(http.StatusBadRequest,gin.H{"error":"id is positive"})
		return
	}
	postID := uint(intPostID)

	ctx := c.Request.Context()
	tags,err:=h.useCase.AssignTagsToPost(ctx,userID.(uint),postID,req.Names)

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return	
	}

	TagResponses := make([]TagResponse, len(tags))
	for i, tag := range tags {
		TagResponses[i] = TagResponse{
			ID:   tag.ID,
			Name: tag.Name,
		}
	}
	c.JSON(http.StatusOK, AssignTagResponse{Tags: TagResponses})

}