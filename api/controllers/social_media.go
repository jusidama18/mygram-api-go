package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/parameters"
	"github.com/jusidama18/mygram-api-go/api/responses"
	"github.com/jusidama18/mygram-api-go/services"
)

type SocialMediaController struct {
	svc *services.SocialMediaService
}

func NewSocialMediaController(svc *services.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{
		svc: svc,
	}
}

func (sm *SocialMediaController) CreateSocialMedia(c *gin.Context) {
	var req parameters.SocialMediaCreate

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

	socialMediaResp, err := sm.svc.CreateSocialMedia(req, user.ID)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}
	responses.SuccessWithData(c, http.StatusCreated, socialMediaResp, "social media successfully created")
}

func (sm *SocialMediaController) GetAllSocialMedia(c *gin.Context) {
	socialMedias, err := sm.svc.GetAllSocialMedia()
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusOK, socialMedias, "successfully get all social medias")
}

func (sm *SocialMediaController) UpdateSocialMedia(c *gin.Context) {
	var req parameters.SocialMediaUpdate

	smID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.BadRequestError(c, "ID must be a number")
		return
	}

	err = c.ShouldBindJSON(&req)
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

	resp, err := sm.svc.UpdateSocialMedia(req, smID, user.ID)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusAccepted, resp, "social media updated successfully")
}

func (sm *SocialMediaController) DeleteSocialMedia(c *gin.Context) {

	smID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.BadRequestError(c, "ID must be a number")
		return
	}

	user, err := GetUser(c)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	err = sm.svc.DeleteSocialMedia(smID, user.ID)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Success(c, http.StatusAccepted, "social media successfully deleted.")
}
