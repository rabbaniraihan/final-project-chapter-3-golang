package controller

import (
	"final-project/model"
	"final-project/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type socialMediaController struct {
	socialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaController(socialMediaRepository repository.SocialMediaRepository) *socialMediaController {
	return &socialMediaController{
		socialMediaRepository: socialMediaRepository,
	}
}

func (sc *socialMediaController) GetAllSocialMedia(ctx *gin.Context) {

	sosmed, err := sc.socialMediaRepository.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, sosmed)
}

func (sc *socialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var newSosmed model.SocialMedia

	err := ctx.ShouldBindJSON(&newSosmed)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	savedSocialMedia, err := sc.socialMediaRepository.Save(newSosmed)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, savedSocialMedia)
}

func (sc *socialMediaController) DeleteSocialMedia(ctx *gin.Context) {
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

	deletedSocialMedia := &model.SocialMedia{Id: id}

	if err := sc.socialMediaRepository.Delete(deletedSocialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, "message : Social media deleted sucesfully")
}

func (sc *socialMediaController) GetSocialMediaById(ctx *gin.Context) {
	stringId := ctx.Param("id")

	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var getSosmed model.SocialMedia
	getSosmed.Id = id

	err = sc.socialMediaRepository.GetById(&getSosmed)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, getSosmed)
}

func (sc *socialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	stringId := ctx.Param("id")

	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updatedSosmed model.SocialMedia

	err = ctx.ShouldBindJSON(&updatedSosmed)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	updatedSosmed.Id = id

	_, err = sc.socialMediaRepository.Update(&updatedSosmed)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedSosmed)
}
