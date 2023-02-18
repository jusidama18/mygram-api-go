package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/parameters"
	"github.com/jusidama18/mygram-api-go/api/responses"
	"github.com/jusidama18/mygram-api-go/models"
	"github.com/jusidama18/mygram-api-go/services"
)

type UserController struct {
	svc *services.UserService
}

func NewUserController(svc *services.UserService) *UserController {
	return &UserController{
		svc: svc,
	}
}

// @Summary Register New User
// @Description Register New User by Data Provided
// @Tags Users
// @Accept json
// @Produce json
// @Param mygram body parameters.UserRegister true "Register User"
// @Success 200 {object} responses.Response{data=models.User}
// @Router /users/register [post]
func (u *UserController) RegisterUser(c *gin.Context) {
	var req parameters.UserRegister

	err := c.ShouldBindJSON(&req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	errs := parameters.Validate(req)
	if errs != nil {
		responses.BadRequestError(c, errs)
		return
	}

	user, err := u.svc.RegisterUser(&req)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusCreated, user, "user successfully created")
}

// @Summary Login User
// @Description Login User by Data Provided
// @Tags Users
// @Accept json
// @Produce json
// @Param mygram body parameters.UserLogin true "Login User"
// @Success 200 {object} responses.Response{data=models.User}
// @Router /users/login [post]
func (u *UserController) Login(c *gin.Context) {
	var req parameters.UserLogin

	err := c.ShouldBindJSON(&req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	errs := parameters.Validate(req)
	if errs != nil {
		responses.BadRequestError(c, errs)
		return
	}

	token, err := u.svc.Login(req.Email, req.Password)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusOK, gin.H{
		"token": token,
	}, "login successful")
}

func (u *UserController) Check(c *gin.Context) {
	userInfo, exists := c.Get("userInfo")
	if !exists {
		responses.InternalServerError(c, fmt.Errorf("context error"))
	}

	user := userInfo.(*models.User)

	responses.SuccessWithData(c, http.StatusOK, user, "success")
}

// @Summary Update User
// @Description Update User by Data Provided
// @Tags Users
// @Accept json
// @Produce json
// @Param mygram body parameters.UserUpdate true "Update User"
// @Success 200 {object} responses.Response{data=models.User}
// @Router /users [put]
func (u *UserController) UpdateUser(c *gin.Context) {
	var req parameters.UserUpdate

	// Binding request body ke variabel req
	err := c.ShouldBindJSON(&req)
	if err != nil {
		responses.BadRequestError(c, err.Error())
		return
	}

	// Validasi request body apakah sudah benar
	// Variabel errs berbeda dengan err
	errs := parameters.Validate(req)
	if errs != nil {
		responses.BadRequestError(c, errs)
		return
	}

	user, err := GetUser(c)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responseUser, err := u.svc.UpdateUser(user, &req)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.SuccessWithData(c, http.StatusOK, responseUser, "user updated successfully")
}

// @Summary Delete User
// @Description Delete User through the authentication process must be done with the help of JsonWebToken.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response{data=string}
// @Router /users [delete]
func (u *UserController) DeleteUser(c *gin.Context) {
	user, err := GetUser(c)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	err = u.svc.DeleteUser(user)
	if err != nil {
		responses.InternalServerError(c, err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "Your account has been successfully deleted")
}
