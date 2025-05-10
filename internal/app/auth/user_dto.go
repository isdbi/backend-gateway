package auth

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignupUserRequest struct {
	Email      string
	Password   string
	Name       string
	FamilyName string
	Age        int
	Sex        string
}

func (r SignupUserRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(12, 255)),
		validation.Field(&r.Name, validation.Required, validation.Length(3, 32)),
		validation.Field(&r.Age, validation.Min(12)),
		validation.Field(&r.FamilyName, validation.Required, validation.Length(3, 22)),
	)
}

type LoginUserRequest struct {
	Email    string
	Password string
}

func (r LoginUserRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(12, 255)),
	)
}
