package crud

import (
	"log"
	"online-supermarket/controllers/ent"
	"online-supermarket/controllers/ent/user"
	"time"
)

func (crud Crud) AddUser(u *ent.User) (*ent.User, error) {
	newUser, err := crud.Client.User.Create().
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(crud.Ctx)
	if err != nil {
		log.Println("on AddUser in controllers/crud/user.go: ", err)
		return nil, err
	}
	return newUser, nil
}

func (crud Crud) GetUserByEmail(email *string) (*ent.User, error) {
	if fetchedUser, err := crud.Client.User.Query().Where(user.Email(*email)).First(crud.Ctx); err != nil {
		log.Println("on GetUserByEmail in controllers/crud/user.go: ", err)
		return nil, err
	} else {
		return fetchedUser, nil
	}
}
