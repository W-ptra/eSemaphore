package utils

import (
	"eSemaphore-backend/model"
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJwt(user model.User, jwtKey []byte, jwtExpired int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    	user.Id,
		"name":  	user.Name,
		"email": 	user.Email,
		"imageUrl": user.ImageUrl,
		"exp":   	time.Now().Add(time.Hour * time.Duration(jwtExpired)).Unix(),
	})

	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}

func GetJwt(tokenString string, jwtKey []byte) (model.User, error) {
	token, err := parseJWT(tokenString, jwtKey)
	if err != nil {
		return model.User{}, err
	}

	claims,ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid{
		return model.User{}, fmt.Errorf("invalid token")
	}

	user := model.User{
		Id:  		claims["id"].(string) ,
		Name:  		claims["name"].(string),
		Email: 		claims["email"].(string),
		ImageUrl: 	claims["imageUrl"].(string),
	}
	return user, nil
}

func decodeJwt(jwtKey []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	}
}

func parseJWT(tokenString string, jwtKey []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenString, decodeJwt(jwtKey))
}