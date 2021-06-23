package util

/*
JWT(json web token), for authentication.
*/

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zhang21/learn_gin/pkg/setting"
)

var jwtSecret = []byte(setting.JwtSecret)

// Claims is
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken generates jwt token use username and password
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	// NewWithClaims(method SigningMethod, claims Claims)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// func (t *Token) SignedString(key interface{}) generates token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parses token.
func ParseToken(token string) (*Claims, error) {
	// func (p *Parser) ParseWithClaims parses claims
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// func (m MapClaims) Valid() valids time
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
