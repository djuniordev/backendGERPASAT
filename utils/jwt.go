package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/djuniordev/SAST-MA/dto"
	"github.com/joho/godotenv"
)

var secretKey []byte

func init() {
	// Carregar variáveis do .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Sem o arquivo .env")
	}

	// Obter a chave secreta do .env
	secretKey = []byte(os.Getenv("SECRET_KEY"))
	if secretKey == nil {
		panic("Secret key not found in .env")
	}

	log.Printf("Secret key loaded successfully")
}

// Claims estrutura para armazenar informações no token
type Claims struct {
	User      dto.UserDtoResponse `json:"user"`
	IsRefresh bool                `json:"is_refresh,omitempty"` // Campo para identificar refresh token
	jwt.StandardClaims
}

// GenerateToken gera um novo token JWT usando HS256
func GenerateToken(user dto.UserDtoResponse) (string, error) {
	expirationTime := time.Now().Add(1 * time.Minute)

	claims := &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func GenerateRefreshToken(user dto.UserDtoResponse) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		User:      user,
		IsRefresh: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	//log.Printf("Validating token: %s", tokenString)

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		//log.Printf("Token parsing error: %v", err)
		return nil, err
	}

	if !token.Valid {
		//log.Printf("Token is invalid: %s", tokenString)
		return nil, errors.New("invalid token")
	}

	//log.Printf("Token is valid, claims: %+v", claims)
	return claims, nil
}
