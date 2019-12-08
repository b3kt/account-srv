package router

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"github.com/b3kt/account-srv/controller"
	"github.com/b3kt/account-srv/middleware"
	"github.com/b3kt/account-srv/model"
)

// Route makes the routing
func Route(app *gin.Engine) {
	indexController := new(controller.IndexController)
	app.GET(
		"/", indexController.GetIndex,
	)

	auth := app.Group("/auth")
	authMiddleware := middleware.Auth()
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			user, _ := c.Get("email")
			c.JSON(200, gin.H{
				"email": claims["email"],
				"name":  user.(*model.User).Username,
				"text":  "Hello World.",
			})
		})
	}

	userController := new(controller.UserController)
	app.GET(
		"/user/:id", userController.GetUser,
	)

	app.POST(
		"/signup", userController.Signup,
	)

	app.POST(
		"/login", userController.Signin,
	)

	api := app.Group("/api")
	{
		api.GET("/version", indexController.GetVersion)
	}
}
