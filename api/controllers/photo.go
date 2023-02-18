package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/parameters"
	"github.com/jusidama18/mygram-api-go/api/responses"
	"github.com/jusidama18/mygram-api-go/services"
)

type PhotoController struct {
	svc *services.PhotoService
}

func NewPhotoController(svc *services.PhotoService) *PhotoController {
	return &PhotoController{
		svc: svc,
	}
}


// @Summary Create Photo
// @Description Create Photo by Data Provided
// @Tags Photos
// @Accept json
// @Produce json
// @Param data body parameters.CreatePhoto true "Create Photo"
// @Success 200 {object} responses.Response{data=models.Photo}
// @Router /photos [post]
func (p *PhotoController) CreatePhoto(c *gin.Context) {
	var req parameters.CreatePhoto

	err := c.ShouldBindJSON(&req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	errs := parameters.Validate(&req)
	if errs != nil {
		responses.BadRequestError(c, errs)
		return
	}

	user, err := GetUser(c)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	photo, err := p.svc.CreatePhoto(req, user.ID)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusCreated, photo, "photo successfully created")
}


// @Summary Get All Photo
// @Description Get All Photo
// @Tags Photos
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response{data=[]models.PhotoGet}
// @Router /photos [get]
func (p *PhotoController) GetAllPhotos(c *gin.Context) {
	photos, err := p.svc.GetAllPhotos()
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}
	responses.SuccessWithData(c, http.StatusOK, photos, "successfully get all photos")
}


// @Summary Update Photo
// @Description Update Photo by Data Provided
// @Tags Photos
// @Accept json
// @Produce json
// @Param data body parameters.UpdatePhoto true "Update Photo"
// @Param id path int true "Photo ID"
// @Success 200 {object} responses.Response{data=models.PhotoUpdate}
// @Router /photos/{id} [put]
func (p *PhotoController) UpdatePhoto(c *gin.Context) {
	var req parameters.UpdatePhoto

	photoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.BadRequestError(c, "ID must be a number")
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	user, err := GetUser(c)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	updatedPhoto, err := p.svc.UpdatePhoto(req, photoID, user.ID)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusAccepted, updatedPhoto, "photo updated successfully.")
}


// @Summary Delete Photo
// @Description Delete Photo by Data Provided
// @Tags Photos
// @Accept json
// @Produce json
// @Param id path int true "Delete Photo"
// @Success 200 {object} responses.Response{data=string}
// @Router /photos/{id} [delete]
func (p *PhotoController) DeletePhoto(c *gin.Context) {
	photoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.BadRequestError(c, "ID must be a number")
		return
	}

	user, err := GetUser(c)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	err = p.svc.DeletePhoto(photoID, user.ID)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Success(c, http.StatusAccepted, "photo successfully deleted.")
}
