package utils

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// CreateToken generates token with userID
func CreateToken(userID int) (string, error) {
	var err error

	expMinutes := viper.GetInt("security.jwtExpMinutes")

	if err != nil {
		return "", err
	}

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(expMinutes)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	jwtSecret := viper.GetString("security.jwtSecret")
	token, err := at.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// ValidateToken validated token according to method and secret
func ValidateToken(tokenString string) bool {
	jwtSecret := viper.GetString("security.jwtSecret")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}

	return false
}

// DecodeToken decodes the specified jwt token and returns the necessary data
func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	jwtSecret := viper.GetString("security.jwtSecret")

	token, err := StripBearerPrefix(tokenString)

	claims := jwt.MapClaims{}

	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}

// StripBearerPrefix removes bearer prefix in JWT
func StripBearerPrefix(token string) (string, error) {
	if len(token) > 6 && strings.ToUpper(token[0:7]) == "BEARER " {
		return token[7:], nil
	}
	return token, nil
}
