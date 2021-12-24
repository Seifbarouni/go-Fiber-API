package services

import (
	"errors"
	"os"
	m "projects/Go-Fiber/api/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Users []*m.User

var allUsers = Users{
	&m.User{ID: 1, Name: "test1", Email: "test1@gmail.com", Password: "123456"},
	&m.User{ID: 2, Name: "test2", Email: "test2@gmail.com", Password: "789555"},
	&m.User{ID: 3, Name: "test3", Email: "test3@gmail.com", Password: "321123"},
}

func Login(userToAuthenticate *m.User) (string, error) {
	// check if user exists
	for _, user := range allUsers {
		if user.Email == userToAuthenticate.Email {
			if user.Password == userToAuthenticate.Password {
				// generate the jwt token
				return generateToken(user.Name, user.ID)
			} else {
				return "", errors.New("wrong email or password")
			}
		}
	}
	return "", errors.New("user not found")
}

func Register(newUser *m.User) (string, error) {
	// check if user exists
	for _, user := range allUsers {
		if user.Email == newUser.Email {
			return "", errors.New("user already exists")
		}
	}
	// add user to the list (change this to a database)
	hash, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	allUsers = append(allUsers, &m.User{ID: len(allUsers) + 1, Name: newUser.Name, Email: newUser.Email, Password: string(hash)})
	// generate the jwt token
	return generateToken(newUser.Name, len(allUsers))
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
