package model

import (
	"log"
	"time"

	"github.com/Nerzal/gocloak/v3"
	"github.com/b3kt/account-srv/config"
	"github.com/b3kt/account-srv/helper"
)

// User the user model
type User struct {
	// ID        uint      `gorm:"primary_key" json:"id"`
	UserID    string    `json:"userID"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetFirstByID gets the user by his ID
func (u *User) GetFirstByID(id string) error {
	// db := DB().Where("id=?", id).First(u)

	// if db.RecordNotFound() {
	// 	return ErrDataNotFound
	// } else if db.Error != nil {
	// 	return db.Error
	// }

	return nil
}

// CheckUserByEmailAndUsername gets the user by his email
func (u *User) CheckUserByEmailAndUsername(email string, username string) error {

	client := gocloak.NewClient(config.KeycloakAdmin.BaseURL)
	accessToken, err := client.LoginAdmin(config.KeycloakAdmin.Username, config.KeycloakAdmin.Password, config.KeycloakAdmin.AdminRealm)
	if err != nil {
		panic("Something wrong with the credentials or url")
		return ErrSystemConfigurationFailure
	}

	users, err := client.GetUsers(accessToken.AccessToken, config.KeycloakAdmin.Realm, gocloak.GetUsersParams{
		Email: email,
	})
	if err != nil {
		log.Println("Failed while fetch user by username", username)
		return ErrSystemIntegrationFailure
	}

	users, err = client.GetUsers(accessToken.AccessToken, config.KeycloakAdmin.Realm, gocloak.GetUsersParams{
		Username: username,
	})
	if err != nil {
		log.Println("Failed while fetch user by username", username)
		return ErrSystemIntegrationFailure
	}

	if users != nil {
		for i := range users {
			if users[i].Email == email {
				log.Println("Email already used already exists", users[i].Email)
				return ErrUserExists
			}
			if users[i].Username == username {
				log.Println("Username already used already exists", users[i].Email)
				return ErrUserExists
			}
		}
	}

	return nil
}

// CheckUserByEmail gets the user by his email
func (u *User) CheckUserByEmail(email string) (bool, error) {
	user, err := u.GetUserByEmail(email)
	if err != nil {
		panic("Error while checking user")
	} else {
		if user != nil {
			return false, ErrUserExists
		}
	}
	return true, nil
}

// GetUserByEmail - gets user by his email
func (u *User) GetUserByEmail(email string) (*User, error) {
	client := gocloak.NewClient(config.KeycloakAdmin.BaseURL)
	accessToken, err := client.LoginAdmin(config.KeycloakAdmin.Username, config.KeycloakAdmin.Password, config.KeycloakAdmin.AdminRealm)
	if err != nil {
		panic("Something wrong with the credentials or url")
	}

	users, err := client.GetUsers(accessToken.AccessToken, config.KeycloakAdmin.Realm, gocloak.GetUsersParams{
		Email: u.Email,
	})
	if err != nil {
		return nil, err
	}

	if users != nil {
		foundUser, isfound := helper.FindByEmail(users, email)
		if isfound {
			user := User{
				Username:  foundUser.Username,
				FirstName: foundUser.FirstName,
				LastName:  foundUser.LastName,
				Email:     foundUser.Email,
				UserID:    foundUser.ID,
			}

			log.Println("User found used", user)
			return &user, nil
		}

		log.Println("User not found")

	}

	return nil, nil
}

// Signup a new user
func (u *User) Signup() error {
	var existinguser User
	err := existinguser.CheckUserByEmailAndUsername(u.Email, u.Username)
	if err != nil {
		log.Println("Failed while signup", err.Error())
		return err
	}

	client := gocloak.NewClient(config.KeycloakAdmin.BaseURL)
	token, err := client.LoginAdmin(config.KeycloakAdmin.Username, config.KeycloakAdmin.Password, config.KeycloakAdmin.AdminRealm)
	if err != nil {
		panic("Something wrong with the credentials or url")
	}
	user := gocloak.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Enabled:   true,
		Username:  u.Username,
	}
	password := u.Password

	client.CreateUser(token.AccessToken, config.KeycloakAdmin.Realm, user)
	if err != nil {
		panic("Oh no!, failed to create user :(")
	}

	client.SetPassword(token.AccessToken, user.ID, config.KeycloakAdmin.Realm, password, false)
	if err != nil {
		panic("Failed while setting user password :(")
	}

	execAction := gocloak.ExecuteActionsEmail{
		UserID:      user.ID,
		ClientID:    config.KeycloakAdmin.ClientID,
		Lifespan:    1000,
		RedirectURI: "",
		Actions:     nil,
	}

	client.ExecuteActionsEmail(token.AccessToken, config.KeycloakAdmin.Realm, execAction)
	if err != nil {
		panic("Failed while setting user password :(")
	}

	return nil
}

// Login a user
func (u *User) Login() (*gocloak.JWT, error) {
	client := gocloak.NewClient(config.KeycloakAdmin.BaseURL)

	log.Println("Login process begin")
	return client.Login(config.KeycloakAdmin.ClientID, config.KeycloakAdmin.ClientSecret, config.KeycloakAdmin.Realm,
		u.Username, u.Password)

}

// RecoverAccount - execute Action Email
func (u *User) RecoverAccount(token string, execActionEmail gocloak.ExecuteActionsEmail) error {
	client := gocloak.NewClient(config.KeycloakAdmin.BaseURL)

	log.Println("proceed to Recover account")
	return client.ExecuteActionsEmail(token, config.KeycloakAdmin.Realm, execActionEmail)
}

// ResetUserPassword - used to reset user password
func (u *User) ResetUserPassword(email string, recoveryToken string, password string) error {
	client := gocloak.NewClient(config.KeycloakAdmin.BaseURL)
	adminToken, err := client.LoginAdmin(config.KeycloakAdmin.Username, config.KeycloakAdmin.Password, config.KeycloakAdmin.AdminRealm)
	if err != nil {
		panic("Something wrong with the credentials or url")
	}

	user, err := u.GetUserByEmail(email)
	if err != nil {
		panic("Unable to get user")
	}

	log.Println("proceed to Reset Password1", adminToken.AccessToken)
	log.Println("proceed to Reset Password2", user.UserID)
	log.Println("proceed to Reset Password3", config.KeycloakAdmin.Realm)
	log.Println("proceed to Reset Password4", password)
	err = client.SetPassword(adminToken.AccessToken, user.UserID, config.KeycloakAdmin.Realm, password, false)
	if err != nil {
		panic("Reset password failed")
	}

	return err
}
