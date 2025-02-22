package main

import (
	"context"
	user "github.com/AppleGrey/Kitex-Demo/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (resp *user.DeleteUserResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUser(ctx context.Context, req *user.QueryUserRequest) (resp *user.QueryUserResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	// TODO: Your code here...
	return
}
