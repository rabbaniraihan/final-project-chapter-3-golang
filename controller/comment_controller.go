package controller

import (
	"final-project/model"
	"final-project/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentRepository repository.CommentRepository
}

func NewCommentController(commentRepository repository.CommentRepository) *commentController {
	return &commentController{
		commentRepository: commentRepository,
	}
}

func (cc *commentController) GetAllComment(ctx *gin.Context) {

	comments, err := cc.commentRepository.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

func (cc *commentController) CreateComment(ctx *gin.Context) {
	var newComment model.Comment

	err := ctx.ShouldBindJSON(&newComment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	savedComment, err := cc.commentRepository.Save(newComment)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, savedComment)
}

func (cc *commentController) DeleteComment(ctx *gin.Context) {
	//Ambil id dari param
	stringId := ctx.Param("id")

	//Convert string -> int
	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	deletedComment := &model.Comment{Id: id}

	if err := cc.commentRepository.Delete(deletedComment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, "message : Comment deleted sucesfully")
}

func (cc *commentController) GetCommentById(ctx *gin.Context) {
	stringId := ctx.Param("id")

	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var getComment model.Comment
	getComment.Id = id

	err = cc.commentRepository.GetById(&getComment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get comment by id",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, getComment)
}

func (cc *commentController) UpdateComment(ctx *gin.Context) {
	stringId := ctx.Param("id")

	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updatedComment model.Comment

	err = ctx.ShouldBindJSON(&updatedComment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	updatedComment.Id = id

	_, err = cc.commentRepository.Update(&updatedComment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedComment)
}
