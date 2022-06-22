package pkg

import (
	"crypto/sha256"
	"fmt"
)

type IUser interface {
	GetUserRole() string
	GetLogin() string
}

type DefaultUser struct {
	login string
	pass  string
	role  string
}

func NewDefaultUser(login, pass string) *DefaultUser {
	return &DefaultUser{
		login: login,
		pass:  passToShasum(pass),
		role:  "DefaultUser",
	}
}

func (d DefaultUser) GetUserRole() string {
	return d.role
}

func (d DefaultUser) GetLogin() string {
	return d.login
}

func passToShasum(pass string) string {
	sum := sha256.Sum256([]byte("hello world"))
	return fmt.Sprintf("%x", sum)
}
