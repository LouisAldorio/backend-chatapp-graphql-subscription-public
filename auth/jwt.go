package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type SubscriberClaim struct {
	SubscriberID int
	jwt.StandardClaims
}

var jwtKey = []byte("secret")

func CreateToken(subscriberId int) (string, error) {
	var signingMethod = jwt.SigningMethodHS256
	var expiredTime = time.Now().AddDate(7, 0, 0).UnixNano() / int64(time.Millisecond)

	customClaim := SubscriberClaim{
		SubscriberID: subscriberId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime,
		},
	}

	token := jwt.NewWithClaims(signingMethod, customClaim)

	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", gqlerror.Errorf(fmt.Sprintf("%s", err))
	}

	return signedToken, nil
}

func ValidateToken(t string) (*jwt.Token, error) {
	token, _ := jwt.ParseWithClaims(t, &SubscriberClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}

		return jwtKey, nil
	})


	return token, nil
}
