package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"got/examples/lgorm/model"
	"gorm.io/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main4()  {
	db,err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/october?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {

	}
	
	pconfig := model.GrowthPublisherChannelConfig{
		ID:              0,
		Line:            2,
		CaS:             "tg_jrtt",
		CaSName:         "今日头条",
		CaN:             "jrtt1",
		CaNName:         "头条1",
		IsCasAll:        1,
		DataType:        0,
		PublisherType:   0,
		AccountConfigID: 0,
		IsValid:         0,
		CreatorID:       0,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		Category:        0,
		IsDelete:        0,
	}

	db.Migrator().CreateTable(&model.GrowthPublisherChannelConfig{})
	//db.CreateTable(&model.GrowthPublisherChannelConfig{})

	result := db.Create(&pconfig)
	db.Migrator()
	fmt.Println(result)




}