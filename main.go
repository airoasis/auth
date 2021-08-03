package main

import (
	"fmt"
	"github.com/airoasis/auth/config"
	"github.com/airoasis/auth/model/entity"
	"github.com/airoasis/auth/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open(postgres.Open(config.DbDSN(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}

	config.DB.AutoMigrate(&entity.User{})

	r := router.SetupRouter()
	r.Run()
}

