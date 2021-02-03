package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/october")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		id        int
		migration string
	)
	rows, err := db.Query("select id,migration from migrations")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &migration)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, migration)
	}
}
