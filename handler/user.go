package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
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
		response = helper.APIResponse("Unable to register the user.",http.StatusServiceUnavailable,"error",nil)
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	formattedUser := user.FormatUser(userData, "")
	response = helper.APIResponse("User is successfully registered",http.StatusOK,"success",formattedUser)
	fmt.Println(response)

	c.JSON(http.StatusOK, response)
}