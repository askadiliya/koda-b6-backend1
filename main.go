package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results"`
}

type Users struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UseForm struct {
	Email    string `json:"email" form:"xxx"`
	Password string `json:"password" form:"yyy"`
}

var ListUser []Users
var autoID int = 1

func main() {
	r := gin.Default()

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.Data(200, "text/html", []byte("<h1>Hello!</h1>"))
	// })
	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.Data(200, "application/json", []byte(`{"success": true, "message": "Backedn is running well"}`))
	// })

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"success": true,
	// 		"message": "ackedn is running well",
	// 	})
	// })

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, Response{
	// 		Success: true,
	// 		Message: "ackedn is running well",
	// 	})
	// })

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "list data user",
			Results: ListUser,
		})
	})

	r.POST("/users", func(ctx *gin.Context) {
		// data := Users{}
		data := UseForm{}

		err := ctx.ShouldBindJSON(&data)

		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "failed create useer",
			})
		} else {
			// ListUser = append(ListUser, data)
			// ctx.JSON(200, Response{
			// 	Success: true,
			// 	Message: "okeee admin",
			// })
			// ListUser = append(ListUser, data)
			// ctx.JSON(200, Response{
			// 	Success: true,
			// 	Message: "okeee admin",
			// })

			// convert UseForm → Users
			user := Users{
				ID:       autoID,
				Email:    data.Email,
				Password: data.Password,
			}

			autoID++
			ListUser = append(ListUser, user)

			ListUser = append(ListUser, user)

			ctx.JSON(200, Response{
				Success: true,
				Message: "okeee admin",
			})

		}
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "5" {
			ctx.JSON(200, Response{
				Success: true,
				Message: "welcome admin",
			})

		} else {
			ctx.JSON(200, Response{
				Success: false,
				Message: fmt.Sprintf("your  id is %s", id),
			})
		}
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		var updateData Users
		if err := ctx.ShouldBindJSON(&updateData); err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "invalid input",
			})
			return
		}

		for i, user := range ListUser {
			if fmt.Sprintf("%d", user.ID) == idParam {

				// Partial update (hanya update kalau tidak kosong)
				if updateData.Email != "" {
					ListUser[i].Email = updateData.Email
				}

				if updateData.Password != "" {
					ListUser[i].Password = updateData.Password
				}

				ctx.JSON(200, Response{
					Success: true,
					Message: "user updated",
					Results: ListUser[i],
				})
				return
			}
		}

		ctx.JSON(404, Response{
			Success: false,
			Message: "user not found",
		})
	})

	r.Run("localhost:8888")
	//sudah membuat app backedn meskipun
}
