package main

import (
	"github.com/AppleGrey/Kitex-Demo/dao/mysql"
	"gorm.io/gen"
	// reuse your gorm db
	// init db
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./model/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	mysql.Init()

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	// reuse your gorm db
	g.UseDB(mysql.DB)

	// Generate basic type-safe DAO API for generated struct `model.User` following conventions
	g.ApplyBasic(g.GenerateModel("users"))

	// Generate the code
	g.Execute()
}
