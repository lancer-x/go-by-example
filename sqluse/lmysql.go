package main

import "database/sql"
import "github.com/go-sql-driver/mysql"

func main()  {
	db,_ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_by_example?charset=utf8mb4&parseTime=True&loc=Local");
	db.Query("")

	a := make([]int, 0, 10)
	a = append(a, 5)
}
