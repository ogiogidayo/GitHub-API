package services

import (
	"context"
	"github.com/ogiogidayo/GitHub-API/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{exec: exec},
	}
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type Services interface {
	UserService
	// issueテーブルを扱うIssueServiceなど、他のサービスインターフェースができたらそれらを追加していく
}

type services struct {
	*userService
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}
