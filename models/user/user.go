package user

import (
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Password  string
}

// BeforeSave hashes password before saving to database
func (u *User) BeforeSave() (e error) {
	// Generate hash from password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil || len(u.Password) < 3 {
		e = errors.New("Unable to hash password")
		log.Fatal(e)
	} else {
		u.Password = string(hash)
	}
	return
}
