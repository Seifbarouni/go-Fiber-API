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

func Login(userToAuthenticate *m.User) (string,string, error) {
	var user m.User
	if result := d.DB.Where("email = ?", userToAuthenticate.Email).First(&user); result.Error != nil || result.RowsAffected == 0 {
		return "","", errors.New("user not found, please register")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userToAuthenticate.Password)); err != nil {
		return "","", errors.New("wrong email or password")
	}
	return generateTokenPair(user)
}

func Register(newUser *m.User) (string,string, error) {
	// check if user exists
	var user m.User
	if result := d.DB.Where("email = ?", newUser.Email).First(&user); result.Error == nil && result.RowsAffected != 0 {
		return "","", errors.New("user already exists, please login")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = string(hash)
	if result := d.DB.Create(&user); result.Error != nil {
		return "","", errors.New("cannot create user")
	}
	// generate the jwt token
	return generateTokenPair(user)
}

func generateTokenPair(user m.User) (string,string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"id":   user.ID,
		"email": user.Email,
		"exp":  time.Now().Add(time.Minute * 10).Unix(),
	})
	
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString,err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", errors.New("cannot generate token")
	}
	refreshTokenString,err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", errors.New("cannot generate refresh token")
	}

	return tokenString,refreshTokenString, nil
	
}

/* func RefreshToken(refreshToken string) (string, error) {
	var user m.User
	if result := d.DB.Where("id = ?", getUserIDFromRefreshToken(refreshToken)).First(&user); result.Error != nil || result.RowsAffected == 0 {
		return "", errors.New("user not found")
	}
	return generateTokenPair(user)
} */
