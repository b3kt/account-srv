package model

import (
	"log"
	"time"

	"github.com/Nerzal/gocloak/v3"
	"github.com/b3kt/account-srv/config"
)

// User the user model
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
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

// GetUserByEmail gets the user by his email
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
		for i := range users {
			if users[i].Email == email {
				log.Println("Email already used", users[i].Email)
				return nil, ErrUserExists
			}
		}
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

	client.CreateUser(token.AccessToken, config.KeycloakAdmin.Realm, user)
	if err != nil {
		panic("Oh no!, failed to create user :(")
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
