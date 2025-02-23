package main

import (
	"github.com/AppleGrey/Kitex-Demo/dao"
	user "github.com/AppleGrey/Kitex-Demo/kitex_gen/user/userservice"
	"log"
)

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	dao.Init()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
