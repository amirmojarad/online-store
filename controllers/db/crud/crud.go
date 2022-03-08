package crud

import (
	"context"
	"online-supermarket/controllers/ent"
)

type Crud struct {
	Ctx    context.Context
	Client *ent.Client
}
