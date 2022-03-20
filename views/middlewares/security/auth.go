package security

import (
	"log"
	"online-supermarket/controllers/auth"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CheckAuth() func(*gin.Context) {
	return func(ctx *gin.Context) {
		service := auth.JWTAuthService()
		// header := ctx.Request.Header

		// t := strings.Split(header["Authorization"][0], " ")[1]

		// claims := jwt.MapClaims{}
		// token, err := jwt.ParseWithClaims(t, claims, func(t *jwt.Token) (interface{}, error) {
		// 	return []byte("secret"), nil
		// })
		// if err != nil {
		// 	log.Println("Error in mdware: ", err)
		// }
		// log.Printf("%v\n", token.Claims)
		header := ctx.Request.Header
		// body := ctx.Request.Body
		// log.Println("HEADER: Accept: ", header.Values("Accept"))
		// log.Println("HEADER: User-Agent: ", header.Values("User-Agent"))
		tokenString := strings.Split(header["Authorization"][0], " ")[1]
		if token, err := service.ValidateToken(tokenString); err != nil {
			log.Println("ERROR: ", err)
		} else {
			// claims, _ := token.Claims.(jwt.MapClaims)
			log.Println("MAP CLAIMS: ", token.Claims.(jwt.Claims).Valid().Error())
		}
		// bs, err := jwt.DecodeSegment(token)
		// log.Println(err, string(bs))
		// log.Println("Body: ", body)

	}
}
