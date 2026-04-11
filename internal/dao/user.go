package dao

import (
	"context"

	"AI-Study-Community/internal/model/do"
	"AI-Study-Community/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"

)

type User struct {
	table string
}

var UserDao = &User{
	table: "users",
}

func (p *User) Insert(ctx context.Context, req *do.UserReq) (uint64, error) {
	user := entity.User{
		Username:      req.Username,
	}



	result, err := p.db().Data(user).Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (p *User) db() *gdb.Model {
	return Database.Model(p.table)
}