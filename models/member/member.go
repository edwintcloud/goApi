package member

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Member model
type Member struct {
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Username  string `bson:"username" json:"username"`
	Password  string `bson:"password" json:"password"`
}

// HashPassword hashes password before saving to database
func (Member) HashPassword(m *Member) *Member {
	// Generate hash from password
	hash, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.Password = string(hash)
	return m
}

// CheckValid makes sure member struct has valid input to insert into database
func (Member) CheckValid(m *Member) error {
	if len(m.Password) < 4 {
		return errors.New("Password must be at least 5 characters!")
	}
	return nil
}
