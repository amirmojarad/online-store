package api

import (
	"log"
	"net/http"
	"online-supermarket/controllers/auth"
	"online-supermarket/controllers/ent"
	"online-supermarket/utils"

	"github.com/gin-gonic/gin"
)

var JwtService = auth.JWTAuthService()

func (api *API) AuthRouter() {
	auth := api.Router.Group("/users")
	auth.POST("/login", api.loginUser())
	auth.POST("/signup", api.signUpUser())
}

func (api *API) loginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userSchema := &ent.User{}
		ctx.BindJSON(&userSchema)
		if fetchedUser, err := api.Crud.GetUserByEmail(&userSchema.Email); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"status":  400,
			})
			return
		} else if utils.CheckPasswordHash(fetchedUser.Password, userSchema.Password) {
			token := JwtService.GenerateToken(fetchedUser.Email, true)
			ctx.IndentedJSON(http.StatusOK, gin.H{
				"token": token,
			})
			return
		}
	}
}

func (api *API) signUpUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userSchema := &ent.User{}
		ctx.BindJSON(&userSchema)
		log.Println(userSchema)
		if createdUser, err := api.Crud.AddUser(userSchema); err != nil {
			log.Println("on SignUpUser in auth.go: ", err)
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"status":  400,
				"message": err.Error(),
			})
			return
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{
				"email":     createdUser.Email,
				"createdAt": createdUser.CreatedAt,
				"updatedAt": createdUser.UpdatedAt,
				"id":        createdUser.ID,
			})
			return
		}
	}
}

// type credential struct {
// 	Email    string `form:"email"`
// 	Password string `form:"password"`
// }

// func (api API) loginUser(email, password string) bool {
// 	fetchedUser, err := api.Crud.GetUserByEmail(email)
// 	log.Println(fetchedUser)
// 	if err != nil {
// 		log.Println("on LoginUser in service/loginService.go: ", err)
// 		return false
// 	}
// 	b := utils.CheckPasswordHash(fetchedUser.Password, password)
// 	log.Println(b)
// 	if !b {
// 		log.Println("on LoginUser in views/api/auth.go -- PASSWORD INCORRECT")
// 		return false
// 	}
// 	return true
// }

// func (api API) Login(path string) {
// 	api.Router.POST(path, func(ctx *gin.Context) {
// 		t := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
// 		log.Println(t)
// 		claims := jwt.MapClaims{}
// 		token, err := jwt.ParseWithClaims(t, claims, func(t *jwt.Token) (interface{}, error) {
// 			return []byte("secret"), nil
// 		})
// 		log.Printf("TOKEN: %v\n", token)
// 		log.Println("Error: ", err)
// 		// var cred credential
// 		// if err := ctx.ShouldBind(&cred); err != nil {
// 		// 	ctx.IndentedJSON(http.StatusBadRequest, gin.H{
// 		// 		"message": err,
// 		// 	})
// 		// 	return
// 		// }
// 		// isUserAuth := api.loginUser(cred.Email, cred.Password)
// 		// log.Println("UserAuth: ", isUserAuth)
// 		// if isUserAuth {
// 		// 	if token := auth.JWTAuthService().GenerateToken(cred.Email, true); token != "" {
// 		// 		ctx.IndentedJSON(http.StatusOK, gin.H{
// 		// 			"token": token,
// 		// 		})
// 		// 		return
// 		// 	} else {
// 		// 		ctx.IndentedJSON(http.StatusUnauthorized, nil)
// 		// 		return
// 		// 	}
// 		// } else {
// 		// 	ctx.IndentedJSON(http.StatusBadRequest, "Incorrect Password")
// 		// }
// 	})
// }

// func (api API) SignUp(path string) {
// 	api.Router.POST(path, func(ctx *gin.Context) {
// 		customer := ent.Customer{}
// 		ctx.BindJSON(&customer)
// 		log.Println(customer)
// 		if newUser, err := api.Crud.AddUser(&customer); err != nil {
// 			log.Println("on SignUp Function in views/api/auth.go: ", err)
// 			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 			return
// 		} else {
// 			ctx.IndentedJSON(http.StatusCreated, newUser)
// 			return
// 		}
// 	})
// }
