package controller

import (
	"final-project/model"
	"final-project/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type photoController struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoController(photoRepository repository.PhotoRepository) *photoController {
	return &photoController{
		photoRepository: photoRepository,
	}
}

func (pc *photoController) GetAllPhoto(ctx *gin.Context) {
	photos, err := pc.photoRepository.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, photos)
}

func (pc *photoController) CreatePhoto(ctx *gin.Context) {
	var newPhoto model.Photo

	err := ctx.ShouldBindJSON(&newPhoto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	savedPhoto, err := pc.photoRepository.Save(newPhoto)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, savedPhoto)
}

func (pc *photoController) DeletePhoto(ctx *gin.Context) {
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

	deletedPhoto := &model.Photo{Id: id}

	if err := pc.photoRepository.Delete(deletedPhoto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, "message : Photo deleted sucesfully")
}

func (pc *photoController) GetPhotoById(ctx *gin.Context) {
	stringId := ctx.Param("id")

	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var getPhoto model.Photo
	getPhoto.Id = id

	err = pc.photoRepository.GetById(&getPhoto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, getPhoto)
}

func (pc *photoController) UpdatePhoto(ctx *gin.Context) {
	stringId := ctx.Param("id")

	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updatedPhoto model.Photo

	err = ctx.ShouldBindJSON(&updatedPhoto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	updatedPhoto.Id = id

	_, err = pc.photoRepository.Update(&updatedPhoto)
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

	ctx.JSON(http.StatusOK, updatedPhoto)
}
