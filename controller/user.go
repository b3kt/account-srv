package controller

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/b3kt/account-srv/helper"
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
			log.Println("Password mismatch")
			response.Header.Message = "Password does not match with conform password"
			response.Header.Error = true
			response.Header.StatusCode = http.StatusOK
			response.Header.Timestamp = time.Now()
			c.JSON(http.StatusOK, response)
			return
		}

		var user model.User
		user.FirstName = request.FirstName
		user.LastName = request.LastName
		user.Username = request.Username
		user.Email = request.Email
		user.Password = request.Password

		if err := user.Signup(); err != nil {
			log.Println("Signup failure")
			response.Header.Message = err.Error()
			response.Header.Error = true
			response.Header.StatusCode = http.StatusOK
			response.Header.Timestamp = time.Now()
			c.JSON(http.StatusOK, response)
		} else {
			log.Println("Signup completed")
			response.Header.Message = "Signup completed"
			response.Header.Error = false
			response.Header.StatusCode = http.StatusOK
			response.Header.Timestamp = time.Now()
			c.JSON(http.StatusOK, response)
		}
	} else {
		log.Println("Invalid format")
		response.Header.Message = "Invalid request format"
		response.Header.Error = true
		response.Header.StatusCode = http.StatusBadRequest
		response.Header.Timestamp = time.Now()
		c.JSON(http.StatusBadRequest, response)
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
			log.Println("Login failed")
			response.Header.Message = err.Error()
			response.Header.Error = true
			response.Header.StatusCode = http.StatusBadRequest
			response.Header.Timestamp = time.Now()
			c.JSON(http.StatusOK, response)
		} else {
			log.Println("Login success")
			response.Header.Message = "Successful logged in"
			response.Header.Error = false
			response.Header.StatusCode = http.StatusOK
			response.Header.Timestamp = time.Now()
			response.AccessToken = token
			response.Body.Username = request.Username
			c.JSON(http.StatusOK, response)
		}
	} else {
		log.Println("invalid request format")
		response.Header.Message = "Invalid request format"
		response.Header.Error = true
		response.Header.StatusCode = http.StatusBadRequest
		response.Header.Timestamp = time.Now()
		c.JSON(http.StatusBadRequest, response)
	}
}

// Recovery to recover user password -- getting recovery token
func (ctrl *UserController) Recovery(c *gin.Context) {
	var request vo.RecoveryRequestMsg
	var response vo.RecoveryResponseMsg

	if err := c.BindJSON(&request); err == nil {

		var user model.User
		user.Email = request.Email

		// should be sent recovery token to requested email
		log.Println("token generated")
		response.Header.Message = "Token generated"
		response.Header.Error = true
		response.Header.StatusCode = http.StatusBadRequest
		response.Header.Timestamp = time.Now()

		// this value should be stored to database
		response.Body.RecoveryToken = strings.ToUpper(helper.GenerateRecoveryToken())
		response.Body.Expire = time.Now().AddDate(0, 0, 1)

		c.JSON(http.StatusBadRequest, response)

	} else {
		log.Println("invalid request format")
		response.Header.Message = "Invalid request format"
		response.Header.Error = true
		response.Header.StatusCode = http.StatusBadRequest
		response.Header.Timestamp = time.Now()
		c.JSON(http.StatusBadRequest, response)
	}
}

// ResetPass to reset user password -- using recovery token
func (ctrl *UserController) ResetPass(c *gin.Context) {

}
