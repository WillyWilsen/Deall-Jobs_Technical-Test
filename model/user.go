package model

import (
    "golang.org/x/crypto/bcrypt"
    "regexp"
)

type User struct {
    Name     string `json:"name,omitempty"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

func ValidateEmail(email string) bool {
    // Regular expression for basic email validation
    pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    regex := regexp.MustCompile(pattern)
    return regex.MatchString(email)
}

func (u *User) HashPassword(plainPassword string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    return nil
}

func (u *User) VerifyPassword(plainPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))
}