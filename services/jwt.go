package services

import (
	"github.com/gin-gonic/gin"
	"os"
	"time"
	"webServer/forms"

	"github.com/dgrijalva/jwt-go"
)

var Tokens = map[string]forms.LoginUserCommand{}

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

// Claims defines jwt claims
type Claims struct {
	UserID string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken handles generation of a jwt code
// @returns string -> token and error -> err
func GenerateToken(userEmail string) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(1440 * time.Minute)
	// Define the payload and exp time
	claims := &Claims{
		UserID: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key encoding
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}
func CreateToken(authData forms.LoginUserCommand, userEmail string, c *gin.Context) string {
	jwtToken, err := GenerateToken(userEmail)
	if err != nil {
		c.JSON(403, gin.H{"message": "There was a problem logging you in, try again later"})
		return err.Error()
	}

	Tokens[jwtToken] = authData

	return jwtToken
}

// DecodeToken handles decoding a jwt token
func DecodeToken(tkStr string) (string, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tkStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}

	if !tkn.Valid {
		return "", err
	}

	return claims.UserID, nil
}
