package service

import (
	"context"

	"AI-Study-Community/internal/dao"
	"AI-Study-Community/internal/model/do"
)

type User struct {
}

var UserService = &User{}

func (s *User) Create(ctx context.Context, req *do.UserReq) (*do.UserResp, error) {
	id, err := dao.UserDao.Insert(ctx, req)
	if err != nil {
		return nil, err
	}

	return &do.UserResp{
		Id: id,
	}, nil
}