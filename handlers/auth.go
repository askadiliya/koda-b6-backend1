package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"backend-demo/models"
	"backend-demo/utils"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results,omitempty"`
}

type AuthInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context) {
	var input AuthInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "invalid request body",
		})
		return
	}

	// Validasi kosong
	if input.Email == "" || input.Password == "" {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "email and password cannot be empty",
		})
		return
	}

	// Validasi format email
	if !utils.IsValidEmail(input.Email) {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "invalid email format",
		})
		return
	}

	// Cek email sudah ada
	for _, u := range models.Users {
		if u.Email == input.Email {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "email already registered",
			})
			return
		}
	}

	user := models.User{
		ID:       models.AutoID,
		Email:    input.Email,
		Password: input.Password,
	}

	models.AutoID++
	models.Users = append(models.Users, user)

	ctx.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "register success",
		Results: user,
	})
}

func Login(ctx *gin.Context) {
	var input AuthInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "invalid request body",
		})
		return
	}

	if input.Email == "" || input.Password == "" {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "email and password cannot be empty",
		})
		return
	}

	for _, u := range models.Users {
		if u.Email == input.Email && u.Password == input.Password {
			ctx.JSON(http.StatusOK, Response{
				Success: true,
				Message: "login success",
				Results: u,
			})
			return
		}
	}

	ctx.JSON(http.StatusUnauthorized, Response{
		Success: false,
		Message: "invalid email or password",
	})
}