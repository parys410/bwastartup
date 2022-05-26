package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// Tangkap input dari user
	// Map input dari user ke struct RegisterUserInput
	// Struct di atas di-passing sebagai parameter ke service

	var input user.RegisterUserInput
	var response helper.Response
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response = helper.APIResponse("Unable to process. Please check your request.",http.StatusUnprocessableEntity,"error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userData, err := h.userService.RegisterUser(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response = helper.APIResponse("Unable to register the user.",http.StatusServiceUnavailable,"error",errorMessage)
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	formattedUser := user.FormatUser(userData, "")
	response = helper.APIResponse("User is successfully registered",http.StatusOK,"success",formattedUser)
	// fmt.Println(response)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// user memasukkan input (email & password)
	// input ditangkap handler
	// mapping dari input user ke input struct
	// input struct passing ke service
	// di service mencari dgn bantuan repository user dengan email x
	// jika ketemu maka perlu mencocokkan password

	var input user.LoginUserInput
	var response helper.Response
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response = helper.APIResponse("login failed",http.StatusUnprocessableEntity,"error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response = helper.APIResponse("login failed",http.StatusServiceUnavailable,"error",errorMessage)
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	formattedUser := user.FormatUser(loggedinUser, "")
	response = helper.APIResponse("successfully logged in",http.StatusOK,"success",formattedUser)

	c.JSON(http.StatusOK, response)
}