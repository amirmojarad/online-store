package crud

import (
	"fmt"
	"log"
	"online-supermarket/controllers/ent"
)

func (crud Crud) DeleteCategory(id int) (string, error) {
	if err := crud.Client.Category.DeleteOneID(id).Exec(crud.Ctx); err != nil {
		log.Println("on DeleteCategory in controllers/db/crud/category.go: ", err)
		return "", err
	} else {
		return fmt.Sprintf("Category with id {%d} deleted completely", id), nil
	}
}

func (crud *Crud) GetAllCategories() ([]*ent.Category, error) {
	if categories, err := crud.Client.Category.Query().All(crud.Ctx); err != nil {
		log.Println("on GetAllCategories in controllers/db/crud/category.go: ", err)
		return nil, err
	} else {
		return categories, nil
	}
}

func (crud *Crud) AddCategories(categories []*ent.Category) ([]*ent.Category, error) {
	bulks := make([]*ent.CategoryCreate, len(categories))
	for i, category := range categories {
		bulks[i] = crud.Client.Category.Create().SetName(category.Name).SetDescription(category.Description).SetThumbnail(category.Thumbnail)
	}
	if addedCategories, err := crud.Client.Category.CreateBulk(bulks...).Save(crud.Ctx); err != nil {
		log.Println("on GetAllCategories in controllers/db/crud/category.go: ", err)
		return nil, err
	} else {
		return addedCategories, nil
	}
}

func (crud *Crud) AddCategory(category *ent.Category) (*ent.Category, error) {
	if newCategory, err := crud.Client.Category.Create().
		SetDescription(category.Description).
		SetName(category.Name).
		SetThumbnail(category.Thumbnail).
		Save(crud.Ctx); err != nil {

		log.Println("on AddCategory in controllers/db/crud/category.go: ", err)
		return nil, err

	} else {
		return newCategory, nil
	}
}
