package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/marmotedu/miniblog/internal/miniblog/store"
	"github.com/marmotedu/miniblog/internal/pkg/errno"
	v1 "github.com/marmotedu/miniblog/internal/pkg/miniblog/v1"
	"github.com/marmotedu/miniblog/internal/pkg/model"
	"regexp"
)

type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
}

type userBiz struct {
	ds store.IStore
}

func (u userBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, r)
	err := u.ds.Users().Create(ctx, &userM)
	if err != nil {
		match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error())
		if match {
			return errno.ErrUserAlreadyExist
		}
	}
	return nil
}

func New(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}
