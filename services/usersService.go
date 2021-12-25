package services

import (
	"errors"
	"os"
	d "projects/Go-Fiber/api/data"
	m "projects/Go-Fiber/api/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(userToAuthenticate *m.User) (string, error) {
	var user m.User
	if result := d.DB.Where("email = ?", userToAuthenticate.Email).First(&user); result.Error != nil || result.RowsAffected == 0 {
		return "", errors.New("user not found, please register")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userToAuthenticate.Password)); err != nil {
		return "", errors.New("wrong email or password")
	}
	return generateToken(user.Name, user.ID)
}

func Register(newUser *m.User) (string, error) {
	// check if user exists
	var user m.User
	if result := d.DB.Where("email = ?", newUser.Email).First(&user); result.Error == nil && result.RowsAffected != 0 {
		return "", errors.New("user already exists, please login")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = string(hash)
	if result := d.DB.Create(&user); result.Error != nil {
		return "", errors.New("cannot create user")
	}
	// generate the jwt token
	return generateToken(user.Name, user.ID)
}

func generateToken(name string, id int) (string, error) {
	claims := jwt.MapClaims{
		"name": name,
		"id":   id,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
