package main

import (
	"fmt"
	"got/util"
	"strconv"

	//"github.com/guregu/null"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"got/examples/lgorm/model"
	"time"
)

func main()  {
	db,err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/go_by_example?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {

	}






	//model, rows := FirstRecord(db)
	//util.DD(model, rows)

	//createData(db)

	//db.Create(&c)
	//db.First(&c)

	//query(db, 20)
	//queryByWhere(db)
	//queryByWhereMap(db)
	//queryByTable(db)
	queryByRaw(db)
}

func FirstRecord(db *gorm.DB) (model.ReportTest, int64) {
	//model := model.ReportTest{}
	var model model.ReportTest
	result := db.First(&model)
	return model,result.RowsAffected
}

func count(db *gorm.DB) {
	var totalRows int64
	resultOrm :=db.Model(&model.ReportTest{})
	resultOrm.Count(&totalRows)
	fmt.Println(totalRows)
}

func query(db *gorm.DB, id int) {
	model := model.ReportTest{ID: uint32(id)}
	ret := db.First(&model)
	util.DD(model, ret.RowsAffected)
}

func queryByWhere(db *gorm.DB)  {
	var models []model.ReportTest
	db.Where("event_type=?", 1).Find(&models)
	util.DD(models)
}

func queryByWhereMap(db *gorm.DB) {
	var models []model.ReportTest
	db.Where(map[string]interface{}{"id" : 111}).Find(&models)
	util.DD(models)
}

func queryByTable(db *gorm.DB) {
	rows,err := db.Table("report_test").Where("id=?", 100).Select("id as data_id,event_type").Rows()
	if err != nil {

	}

	var id, event int64
	for rows.Next() {
		rows.Scan(&id, &event)
		util.DD(id, event)
	}
}

func queryByRaw(db *gorm.DB) {
	type Result struct {
		Id int
		Event_type int
	}
	var result []Result
	db.Table("report_test").Select("id", "event_type").Where("id>?", 198).Scan(&result)
	util.DD(result)
}

func createData(db *gorm.DB)  {
	models := make([]model.ReportTest, 100)
	for i := 0; i < len(models); i++ {
		models[i] = model.ReportTest{
			ID:              0,
			Type:            util.RandomStr(5),
			Callback:        util.RandomStr(20),
			Source:          "guazi",
			ConvTime:        time.Now().Format("2006-01-02 15:04:05"),
			EventType:       0,
			Header:          "",
			Body:            util.RandomStr(10),
			AccountID:       util.RandomStr(13),
			UserActionSetID: "",
			Actions:         util.RandomStr(10),
			EventTime:       strconv.Itoa(int(time.Now().Unix())),
			EventTimealt1:   "",
			JsEventType:     0,
			EventTypealt1:   0,
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
			Token:           util.RandomStr(20),
			ConversionTypes: "",
			PurchaseAmount:  "",
		}
	}

	db.Create(&models)

}

func show(args ...interface{})  {
}


