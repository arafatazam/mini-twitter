package dto

import (
	"regexp"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type NewUserReq struct{
	Handle string		`json:"handle"`
	Name string			`json:"name"`
	Email string		`json:"email"`
	Password string		`json:"password"`
}

func (usr NewUserReq) Validate() error {
	return v.ValidateStruct(&usr,
		v.Field(&usr.Handle,v.Required, v.Match(regexp.MustCompile("^[a-zA-Z0-9_]{1,15}$"))),
		v.Field(&usr.Name, v.Required, v.Min(8)),
		v.Field(&usr.Email, v.Required, is.Email),
		v.Field(&usr.Password, v.Required, v.Min(8)),
	)
}