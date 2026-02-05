package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(id string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

type jwtCustomClaim struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func NewJwtSwervice() JwtService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "jwt",
	}
}

// GenerateToken implements [JwtService].
func (service *jwtService) GenerateToken(id string) string {
	claims := &jwtCustomClaim{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 7).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	generatedToken, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return generatedToken
}

// ValidateToken implements [JwtService].
func (service *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method %v", t.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey != "" {
		secretKey = "jwt"
	}
	return secretKey
}
