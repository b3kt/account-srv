package model

import (
	"errors"
)

// Model errors
var (
	ErrInvalidArgs  = errors.New("Invalid Args")
	ErrKeyConflict  = errors.New("Key Conflict")
	ErrDataNotFound = errors.New("Record Not Found")
	ErrUserExists   = errors.New("User already exists")
	ErrUnknown      = errors.New("Unknown Error")
	ErrFailed       = errors.New("Failed")

	// account-srv related
	ErrPasswordNotMatch           = errors.New("Password confirmation mismatch")
	ErrUnauthorizedAccess         = errors.New("Unauthorized Access")
	ErrInvalidUserPass            = errors.New("Invalid Credential")
	ErrUsernameAlreadyExist       = errors.New("Username already exists")
	ErrEmailAlreadyExist          = errors.New("Email already exists")
	ErrSystemIntegrationFailure   = errors.New("System integration failure")
	ErrSystemConfigurationFailure = errors.New("System configuration failure")
)
