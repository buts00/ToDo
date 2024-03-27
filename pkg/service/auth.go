package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	Todo "github.com/buts00/ToDo"
	"github.com/buts00/ToDo/pkg/repository"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	salt       = "fnifown134bfoaoieafno34"
	tokenTTL   = 12 * time.Hour
	signingKey = "fwefw1ibrqubfeeFTAEBNBETIE52"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.MapClaims
	UserId int `json:"user-id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateNewUser(user Todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateNewUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method ")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.User(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Add(tokenTTL).Unix(),
			"iat": time.Now().Unix(),
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(signingKey))
}
