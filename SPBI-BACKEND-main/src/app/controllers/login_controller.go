package controllers

import (
	"fmt"
	"net/http"
	"sigap-sultan-be/src/app/helper/common_helper"
	"sigap-sultan-be/src/app/helper/jwt_helper"
	"sigap-sultan-be/src/app/models"
	"sigap-sultan-be/src/app/services"
	"sigap-sultan-be/src/common"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type LoginController struct {
	LoginService *services.LoginService
}

func NewLoginController(priceService *services.LoginService) *LoginController {
	return &LoginController{
		priceService,
	}
}

func (r LoginController) Login(c *fiber.Ctx) error {
	var data interface{}
	var errorDomain *common.ErrorDomain
	var err error

	request := new(models.LoginRequestParams)

	err = c.BodyParser(request)
	if err != nil {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	}

	data, errorDomain = r.LoginService.Login(*request)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	hasil := data.(*models.LoginResponse)
	token, err := jwt_helper.GenerateToken(hasil)
	if err != nil {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	}
	hasil.Token = token

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", hasil))
}

func (r LoginController) Forgot(c *fiber.Ctx) error {
	var data *string
	var errorDomain *common.ErrorDomain
	var err error

	request := new(models.ForgotRequestParams)

	err = c.BodyParser(request)
	if err != nil {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	}

	data, errorDomain = r.LoginService.ForgotPassword(*request)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	hasil := new(models.ForgotResponse)

	hasil.Link = *data

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", hasil))
}

func (r LoginController) CheckKode(c *fiber.Ctx) error {
	var data *string
	var errorDomain *common.ErrorDomain
	var err error

	request := new(models.ForgotTokenParam)

	err = c.BodyParser(request)
	if err != nil {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	}

	data, errorDomain = r.LoginService.CheckKode(*request)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r LoginController) ResetPassword(c *fiber.Ctx) error {
	var data *string
	var errorDomain *common.ErrorDomain
	var err error

	request := new(models.ResetRequestParams)

	err = c.BodyParser(request)
	if err != nil {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	}

	if request.NewPassword != request.NewPasswordConfirm {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "New Password dan Confirm tidak sama", nil))
	}

	data, errorDomain = r.LoginService.ResetPassword(*request)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r LoginController) Profile(c *fiber.Ctx) error {
	var errorDomain *common.ErrorDomain
	var err error

	// Get Authorization header
	authHeader := c.Get("Authorization")

	// Check if the header is empty
	if authHeader == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, "Missing Authorization Header", nil))
	}

	// Validate Bearer token format
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, "Invalid Authorization Format", nil))
	}

	// Extract the token
	token := parts[1]

	// Log or process the token
	fmt.Println("Received Token:", token)

	// Continue processing (e.g., token validation)
	dataToken, err := jwt_helper.ParseToken(token)
	if err != nil {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	}

	id := int(*dataToken.Id.Id)

	data, errorDomain := r.LoginService.Profile(id)
	if errorDomain != nil {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	}

	hasil := data.(*models.ProfileResponse)

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", hasil))
}

func (r LoginController) UpdateProfile(c *fiber.Ctx) error {
	var data *string
	var errorDomain *common.ErrorDomain
	var err error

	// Get Authorization header
	authHeader := c.Get("Authorization")

	// Check if the header is empty
	if authHeader == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, "Missing Authorization Header", nil))
	}

	// Validate Bearer token format
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, "Invalid Authorization Format", nil))
	}

	// Extract the token
	token := parts[1]

	// Log or process the token
	fmt.Println("Received Token:", token)

	// Continue processing (e.g., token validation)
	dataToken, err := jwt_helper.ParseToken(token)
	if err != nil {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	}

	id := int(*dataToken.Id.Id)

	request := new(models.UpdateProfileParam)

	err = c.BodyParser(request)
	if err != nil {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	}

	data, errorDomain = r.LoginService.UpdateProfile(id, *request)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}

func (r LoginController) ChangePassword(c *fiber.Ctx) error {
	var data *string
	var errorDomain *common.ErrorDomain
	var err error

	// Get Authorization header
	authHeader := c.Get("Authorization")

	// Check if the header is empty
	if authHeader == "" {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, "Missing Authorization Header", nil))
	}

	// Validate Bearer token format
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, "Invalid Authorization Format", nil))
	}

	// Extract the token
	token := parts[1]

	// Log or process the token
	fmt.Println("Received Token:", token)

	// Continue processing (e.g., token validation)
	dataToken, err := jwt_helper.ParseToken(token)
	if err != nil {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	}

	id := int(*dataToken.Id.Id)

	request := new(models.ChangePasswordParam)

	err = c.BodyParser(request)
	if err != nil {
		return common_helper.GenerateResponseJSON(c, http.StatusInternalServerError, models.SetError().Details(http.StatusInternalServerError, err.Error(), nil))
	}

	if request.NewPassword != request.NewPasswordConfirm {
		return common_helper.GenerateResponseJSON(c, http.StatusBadRequest, models.SetError().Details(http.StatusBadRequest, "New Password dan Confirm tidak sama", nil))
	}

	data, errorDomain = r.LoginService.UpdatePassword(id, request.OldPassword, request.NewPassword)
	if errorDomain != nil {
		return common_helper.ProcessErrorDomain(c, errorDomain)
	}

	return common_helper.GenerateResponseJSON(c, http.StatusOK, models.Set().Success("", data))
}
