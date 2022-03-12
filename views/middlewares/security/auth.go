package security

import (
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckAuth() func(*gin.Context) {
	return func(ctx *gin.Context) {
		header := ctx.Request.Header

		t := strings.Split(header["Authorization"][0], " ")[1]

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(t, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte("<YOUR VERIFICATION KEY>"), nil
		})
		if err != nil {
			log.Println("Error in mdware: ", err)
		}
		log.Printf("%v\n", token)
		// header := ctx.Request.Header
		// body := ctx.Request.Body
		// log.Println("HEADER: Accept: ", header.Values("Accept"))
		// log.Println("HEADER: User-Agent: ", header.Values("User-Agent"))
		// token := strings.Split(header["Authorization"][0], " ")[1]
		// bs, err := jwt.DecodeSegment(token)
		// log.Println(err, string(bs))
		// log.Println("Body: ", body)

	}
}
