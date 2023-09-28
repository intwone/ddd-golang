package cryptography

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTCryptography struct {
	SecretKey []byte
}

func NewJWTCryptography(secretKey []byte) CryptographyInterface {
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
	parsedToken, err := jwt.Parse(removeBearer(token), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(c.SecretKey), nil
		}

		return nil, errors.New("error to parse token")
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		return nil, errors.New("error to cast parsed token claims to jwt MapClaims")
	}

	userID, ok := claims["user_id"].(string)

	if !ok {
		return nil, errors.New("user_id not found in token claims")
	}

	return &userID, nil
}

func removeBearer(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
