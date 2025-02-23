package main

import (
	"context"
	//"fmt"
	"github.com/AppleGrey/Kitex-Demo/dao/mysql"
	user "github.com/AppleGrey/Kitex-Demo/kitex_gen/user"
	"github.com/AppleGrey/Kitex-Demo/model/model"
	"github.com/AppleGrey/Kitex-Demo/model/query"
	"github.com/AppleGrey/Kitex-Demo/pack"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	resp = new(user.UpdateUserResponse)
	u := &model.User{
		Name:         req.Name,
		Introduction: req.Introduce,
	}

	_, err = query.User.WithContext(ctx).Updates(u)
	if err != nil {
		resp.Code = user.Code_DBErr
		resp.Msg = err.Error()
		return
	}

	resp.Code = user.Code_Success
	return
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (resp *user.DeleteUserResponse, err error) {
	resp = new(user.DeleteUserResponse)

	_, err = query.User.WithContext(ctx).Where(query.User.UID.Eq(int32(req.UserId))).Delete()
	if err != nil {
		resp.Code = user.Code_DBErr
		resp.Msg = err.Error()
		return
	}

	resp.Code = user.Code_Success
	return
}

// QueryUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUser(ctx context.Context, req *user.QueryUserRequest) (resp *user.QueryUserResponse, err error) {
	resp = new(user.QueryUserResponse)

	u, m := query.User, query.User.WithContext(ctx)
	if *req.Keyword != "" {
		m = m.Where(u.Introduction.Like("%" + *req.Keyword + "%"))
	}

	var total int64
	total, err = m.Count()
	if err != nil {
		resp.Code = user.Code_DBErr
		resp.Msg = err.Error()
		return
	}

	var users []*model.User
	if total > 0 {
		users, err = m.Limit(int(req.PageSize)).Offset(int(req.PageSize * (req.Page - 1))).Find()
		if err != nil {
			resp.Code = user.Code_DBErr
			resp.Msg = err.Error()
			return
		}
	}

	resp.Code = user.Code_Success
	resp.Totoal = total
	resp.Users = pack.Users(users)

	return
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	resp = new(user.CreateUserResponse)

	if err := mysql.CreateUser([]*model.User{
		{
			Name:         req.Name,
			Introduction: req.Introduce,
		},
	}); err != nil {
		resp.Msg = err.Error()
		resp.Code = user.Code_DBErr
	}

	resp.Code = user.Code_Success
	return

}
