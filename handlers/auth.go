package handlers

import (
	"net/http"
	"time"

	"backend-demo/models"
	"backend-demo/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
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

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{
		ID:       models.AutoID,
		Email:    input.Email,
		Password: string(hashedPassword),
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

		if u.Email == input.Email {

			err := bcrypt.CompareHashAndPassword(
				[]byte(u.Password),
				[]byte(input.Password),
			)

			if err != nil {
				ctx.JSON(http.StatusUnauthorized, Response{
					Success: false,
					Message: "invalid email or password",
				})
				return
			}

			
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user_id": u.ID,
				"email":   u.Email,
				"exp":     time.Now().Add(time.Hour * 1).Unix(),
			})

			tokenString, err := token.SignedString([]byte("secret_key"))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, Response{
					Success: false,
					Message: "failed to generate token",
				})
				return
			}

			
			ctx.JSON(http.StatusOK, Response{
				Success: true,
				Message: "login success",
				Results: gin.H{
					"token": tokenString,
				},
			})
			return
		}
	}

	ctx.JSON(http.StatusUnauthorized, Response{
		Success: false,
		Message: "invalid email or password",
	})
}