package pack

import (
	"github.com/AppleGrey/Kitex-Demo/kitex_gen/user"
	"github.com/AppleGrey/Kitex-Demo/model/model"
)

// Users Convert orm_gen.User list to user_gorm.User list
func Users(models []*model.User) []*user.User {
	users := make([]*user.User, 0, len(models))
	for _, m := range models {
		if u := User(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

// User Convert model.User to user_gorm.User
func User(model *model.User) *user.User {
	if model == nil {
		return nil
	}
	//gender, err := strconv.ParseInt(model.Gender, 10, 64)
	//if err != nil {
	//	return nil
	//}
	return &user.User{
		UserId:    int64(model.UID),
		Name:      model.Name,
		Gender:    user.Gender(1),
		Age:       int64(18),
		Introduce: model.Introduction,
	}
}
