package services

import (
	"context"
	"github.com/ogiogidayo/GitHub-API/graph/model"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}
