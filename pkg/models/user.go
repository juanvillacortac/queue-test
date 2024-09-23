package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID             int64  `json:"id"`
	Email          string `json:"email"`
	HashedPassword string `json:"-"`
}

func (user User) ValidatePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)) == nil
}
