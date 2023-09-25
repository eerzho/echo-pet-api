package model

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email    string `gorm:"not null;unique"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type UserStoreRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"createdAt"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func NewUserFromStoreRequest(request UserStoreRequest) (*User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	return &User{Email: request.Email, Name: request.Name, Password: string(bytes)}, err
}

func CheckPasswordHas(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(id uint, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewUserResponseFromModel(user *User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
