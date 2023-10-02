package cryptography

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/intwone/ddd-golang/internal/constants"
)

type JWTCryptography struct {
	SecretKey string
}

func NewJWTCryptography(secretKey string) CryptographyInterface {
	return &JWTCryptography{
		SecretKey: secretKey,
	}
}

func (c *JWTCryptography) Encrypt(value string) (*string, error) {
	claims := jwt.MapClaims{
		"user_id": value,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(c.SecretKey))

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (c *JWTCryptography) Decrypt(token string) (*string, error) {
	value := removeBearer(token)

	parsedToken, err := jwt.Parse(value, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(c.SecretKey), nil
		}

		return nil, errors.New(constants.InvalidTokenError)
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		return nil, errors.New(constants.TokenClaimsError)
	}

	userID, ok := claims["user_id"].(string)

	if !ok {
		return nil, errors.New(constants.FieldNotFoundOnToken)
	}

	return &userID, nil
}

func removeBearer(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
