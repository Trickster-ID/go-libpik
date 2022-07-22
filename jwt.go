package libpik

import (
	"fmt"
	"log"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

//Getjwtdata(token, secretKey).(int) to return user id
func Getjwtdata(token string, secret string) interface{} {
	respar, _ := jwt.Parse(token, func(tok *jwt.Token)(interface{}, error){
		if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method %v", tok.Header["alg"])
		}
		return []byte(secret), nil
	})
	
	if respar.Valid {
		claims := respar.Claims.(jwt.MapClaims)
		userid := fmt.Sprintf("%v", claims["user_id"])
		log.Println("active userid : ", userid)
		id, errconv := strconv.Atoi(userid)
		if errconv != nil{
			return 0
		}
		return id
	}else{
		return 0
	}
}