package mysql

import "github.com/AppleGrey/Kitex-Demo/model/model"

func CreateUser(users []*model.User) error {
	return DB.Create(users).Error
}
