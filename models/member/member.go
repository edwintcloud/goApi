package member

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// Member model
type Member struct {
	// structs tag is used to match query params, bson matches mongo
	FirstName string `bson:"firstName" json:"firstName" structs:"firstName"`
	LastName  string `bson:"lastName" json:"lastName" structs:"lastName"`
	Email     string `bson:"email" json:"email" structs:"email"`
	Password  string `bson:"password" json:"password" structs:"password"`
}

// HashPassword hashes password before saving to database
func (Member) HashPassword(m *Member) *Member {
	// Generate hash from password
	hash, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.Password = string(hash)
	return m
}

// ComparePasswords compares a memberFound password(hash) to a member password(string)
func (Member) ComparePasswords(fm *Member, m *Member) error {
	if err := bcrypt.CompareHashAndPassword([]byte(fm.Password), []byte(m.Password)); err != nil {
		return errors.New("Invalid password!")
	}
	return nil
}

// CheckValid makes sure member struct has valid input to insert into database
func (Member) CheckValid(m *Member) error {
	// Validate Password
	if len(m.Password) < 5 {
		return errors.New("Password must be at least 6 characters!")
	}
	// Validate FirstName and LastName
	if len(m.FirstName) < 2 || len(m.LastName) < 2 {
		return errors.New("Names must be at least 3 characters!")
	}
	//Validate Email
	if match, _ := regexp.MatchString(`.+@.+\..+`, m.Email); !match {
		return errors.New("Not a valid email address!")
	}
	return nil
}
