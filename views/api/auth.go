package api

import (
	"log"
	"net/http"
	"online-supermarket/controllers/ent"
	"online-supermarket/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type credential struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (api API) loginUser(email, password string) bool {
	fetchedUser, err := api.Crud.GetUserByEmail(email)
	log.Println(fetchedUser)
	if err != nil {
		log.Println("on LoginUser in service/loginService.go: ", err)
		return false
	}
	b := utils.CheckPasswordHash(fetchedUser.Password, password)
	log.Println(b)
	if !b {
		log.Println("on LoginUser in views/api/auth.go -- PASSWORD INCORRECT")
		return false
	}
	return true
}

func (api API) Login(path string) {
	api.Router.POST(path, func(ctx *gin.Context) {
		t := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
		log.Println(t)
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(t, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		log.Printf("TOKEN: %v\n", token)
		log.Println("Error: ", err)
		// var cred credential
		// if err := ctx.ShouldBind(&cred); err != nil {
		// 	ctx.IndentedJSON(http.StatusBadRequest, gin.H{
		// 		"message": err,
		// 	})
		// 	return
		// }
		// isUserAuth := api.loginUser(cred.Email, cred.Password)
		// log.Println("UserAuth: ", isUserAuth)
		// if isUserAuth {
		// 	if token := auth.JWTAuthService().GenerateToken(cred.Email, true); token != "" {
		// 		ctx.IndentedJSON(http.StatusOK, gin.H{
		// 			"token": token,
		// 		})
		// 		return
		// 	} else {
		// 		ctx.IndentedJSON(http.StatusUnauthorized, nil)
		// 		return
		// 	}
		// } else {
		// 	ctx.IndentedJSON(http.StatusBadRequest, "Incorrect Password")
		// }
	})
}

func (api API) SignUp(path string) {
	api.Router.POST(path, func(ctx *gin.Context) {
		customer := ent.Customer{}
		ctx.BindJSON(&customer)
		log.Println(customer)
		if newUser, err := api.Crud.AddUser(&customer); err != nil {
			log.Println("on SignUp Function in views/api/auth.go: ", err)
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		} else {
			ctx.IndentedJSON(http.StatusCreated, newUser)
			return
		}
	})
}
