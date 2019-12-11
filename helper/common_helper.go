package helper

import "github.com/Nerzal/gocloak/v3"

import "log"

// FindByEmail takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func FindByEmail(slice []*gocloak.User, email string) (*gocloak.User, bool) {
	for _, item := range slice {
		if item.Email == email {
			log.Println(" found user ", email)
			return item, true
		}
		log.Println(" compare user ", item.Email, ":", email)

	}
	return nil, false
}
