package jwtservice

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/wuttinanhi/go-jwt-auth-project/config"
)

type JWTService struct {
	secretKey string
	issuer    string
	duration  int64
}

type AuthJWT struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

var jwtService *JWTService = nil

func (s *JWTService) GenerateToken(data *AuthJWT) (string, error) {
	// set token claims
	data.StandardClaims.IssuedAt = time.Now().Unix()
	data.StandardClaims.ExpiresAt = time.Now().Unix() + s.duration
	data.StandardClaims.NotBefore = time.Now().Unix()
	data.StandardClaims.Issuer = s.issuer

	// create new token
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = data

	// sign token
	return token.SignedString([]byte(s.secretKey))
}

func (s *JWTService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	return err == nil
}

func (s *JWTService) ParseToken(token string, out *AuthJWT) {
	parsed, err := jwt.ParseWithClaims(token, &AuthJWT{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		panic(err)
	}

	if claims, ok := parsed.Claims.(*AuthJWT); ok && parsed.Valid {
		*out = *claims
	}
}

// get service singleton
func GetJWTService() *JWTService {
	if jwtService == nil {
		jwtService = &JWTService{
			secretKey: config.GetConfig().JWT_SECRET_KEY,
			issuer:    config.GetConfig().JWT_ISSUER,
			duration:  config.GetConfig().JWT_EXPIRE,
		}
	}

	return jwtService
}
