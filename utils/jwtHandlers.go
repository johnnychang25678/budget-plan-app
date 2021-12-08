package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// enum
type TokenType int

const (
	AccessToken TokenType = iota
	RefreshToken
)

type TokenResult struct {
	Token string
	Err   error
}

type customClaims struct {
	userId string
	jwt.StandardClaims
}

func GenJwtToken(userId int, t TokenType, tokenChan chan TokenResult) {
	var jwtKey string
	var expireTime time.Time
	if t == AccessToken {
		jwtKey = os.Getenv("JWT_ACCESS_KEY")
		expireTime = time.Now().Add(60 * time.Minute) // one hour
	}
	if t == RefreshToken {
		jwtKey = os.Getenv("JWT_REFRESH_KEY")
		expireTime = time.Now().AddDate(0, 3, 0) // three months
	}

	claims := customClaims{
		userId: fmt.Sprint(userId),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		tokenChan <- TokenResult{"", err}
	}
	fmt.Println("signed token:", signedToken)
	tokenChan <- TokenResult{signedToken, nil}
}

func VerifyJwtToken(token string, tokenType TokenType) (int, error) {
	// keyFunc: implement the verificaiton and returns the key

	// keyFunc := func(t *jwt.Token) (interface{}, error) {
	// 	_, ok := t.Method.(*jwt.SigningMethodHMAC) // type assertion
	// 	if ok {
	// 		if tokenType == AccessToken {
	// 			return []byte(os.Getenv("JWT_ACCESS_KEY")), nil
	// 		}

	// 		if tokenType == RefreshToken {
	// 			return []byte(os.Getenv("JWT_REFRESH_KEY")), nil
	// 		}

	// 	}
	// 	return nil, jwt.ErrInvalidKey
	// }

	// tokenClaims, err := jwt.ParseWithClaims(token, &customClaims{}, keyFunc)
	// if err != nil {
	// 	return 0, err
	// }

	// claims, ok := tokenClaims.Claims.(*jwt.Claims)
	// if !ok || !tokenClaims.Valid {
	// }

	// TODO: complete this function
	return 0, nil
}
