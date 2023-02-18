package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/parameters"
	"github.com/jusidama18/mygram-api-go/api/responses"
	"github.com/jusidama18/mygram-api-go/services"
)

type CommentController struct {
	svc *services.CommentService
}

func NewCommentController(svc *services.CommentService) *CommentController {
	return &CommentController{
		svc: svc,
	}
}

// @Summary Create Comment
// @Description Create Comment by Data Provided
// @Tags Comments
// @Accept json
// @Produce json
// @Param data body parameters.CreateComment true "Create Comment"
// @Success 200 {object} responses.Response{data=models.Comment}
// @Router /comments [post]
func (cm *CommentController) CreateComment(c *gin.Context) {
	var req parameters.CreateComment

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

	resComment, err := cm.svc.CreateComment(req, user.ID)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusCreated, resComment, "comment successfully created")
}

// @Summary Get All Comment
// @Description Get All Comment
// @Tags Comments
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response{data=[]models.CommentGetAll}
// @Router /comments [get]
func (cm *CommentController) GetAllComment(c *gin.Context) {
	resComment, err := cm.svc.GetAllComment()
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}
	responses.SuccessWithData(c, http.StatusOK, resComment, "successfully get all comments")
}

// @Summary Update Comment
// @Description Update Comment by Data Provided
// @Tags Comments
// @Accept json
// @Produce json
// @Param data body parameters.UpdateComment true "Update Comment"
// @Param id path int true "Comment ID"
// @Success 200 {object} responses.Response{data=models.CommentUpdate}
// @Router /comments/{id} [put]
func (cm *CommentController) UpdateComment(c *gin.Context) {
	var req parameters.UpdateComment

	commentID, err := strconv.Atoi(c.Param("id"))
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

	updatedComment, err := cm.svc.UpdateComment(req, commentID, user.ID)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusAccepted, updatedComment, "comment updated successfully.")
}

// @Summary Delete Comment
// @Description Delete Comment by Data Provided
// @Tags Comments
// @Accept json
// @Produce json
// @Param id path int true "Delete Comment"
// @Success 200 {object} responses.Response{data=string}
// @Router /comments/{id} [delete]
func (cm *CommentController) DeleteComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.BadRequestError(c, "ID must be a number")
		return
	}

	user, err := GetUser(c)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	err = cm.svc.DeleteComment(commentID, user.ID)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Success(c, http.StatusAccepted, "comment successfully deleted.")
}
