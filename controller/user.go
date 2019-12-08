package controller

import (
	"log"
	"net/http"

	"github.com/b3kt/account-srv/model"
	"github.com/b3kt/account-srv/vo"
	"github.com/gin-gonic/gin"
)

// UserController is the user controller
type UserController struct{}

// Signup struct
// type Signup struct {
// 	Email     string `form:"email" json:"email" binding:"required"`
// 	Username  string `form:"username" json:"username" binding:"required"`
// 	Password  string `form:"password" json:"password" binding:"required,min=6,max=20"`
// 	Password2 string `form:"password2" json:"password2" binding:"required"`
// }

// GetUser gets the user info
func (ctrl *UserController) GetUser(c *gin.Context) {
	var user model.User

	id := c.Param("id")

	if err := user.GetFirstByID(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// Signup a new user
func (ctrl *UserController) Signup(c *gin.Context) {
	var request vo.SignupRequestMsg
	var response vo.SignupResponseMsg

	if err := c.BindJSON(&request); err == nil {
		if request.Password != request.PasswordConfirm {
			c.JSON(http.StatusOK, gin.H{"error": "Password does not match with conform password"})
			return
		}

		var user model.User
		user.FirstName = request.FirstName
		user.LastName = request.LastName
		user.Username = request.Username
		user.Email = request.Email
		user.Password = request.Password

		if err := user.Signup(); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			log.Println("FAILED", err.Error())
		} else {

			response.StatusCode = http.StatusOK
			response.Message = "Success"
			response.ErrorCode = "00"

			c.JSON(http.StatusOK, response)
			log.Println("SUCCESS", request)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// Signin a new user
func (ctrl *UserController) Signin(c *gin.Context) {
	var request vo.SigninRequestMsg
	var response vo.SigninResponseMsg

	if err := c.BindJSON(&request); err == nil {

		var user model.User
		user.Username = request.Username
		// user.Email = request.Email
		user.Password = request.Password

		if token, err := user.Login(); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			log.Println("FAILED", err.Error())
		} else {
			response.StatusCode = http.StatusOK
			response.Message = "Success"
			response.ErrorCode = "00"
			response.AccessToken = token.AccessToken
			response.Username = user.Username

			c.JSON(http.StatusOK, response)
		}
	} else {
		log.Println("FAILS")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
