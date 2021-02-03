package main

import (
	"fmt"
	"got/lmysql"
	"log"
	"time"
)

func main() {
	dns := "root:123456@tcp(127.0.0.1:3306)/october"
	db := lmysql.Conn(dns)
	rows, _ := db.Query("select id,migration from migrations")
	println(rows)

	defer rows.Close()

	var (
		id        int
		migration string
	)

	for rows.Next() {
		err := rows.Scan(&id, &migration)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, migration)
	}

	row := db.QueryRow("select id,migration from migrations order by id desc")

	row.Scan(&id, &migration)
	println(id, migration)

	result, err := db.Exec("insert into migrations(migration,batch) values (?,?)", "lllllll", 5)
	if err != nil {
		println(err.Error())
	} else {
		println(result.LastInsertId())
	}
}

func try() {
	dns := "root:123456@tcp(127.0.0.1:3306)/october"
	d := time.NewTicker(1 * time.Second)

	start := time.After(10 * time.Second)

	for {
		select {
		case <-d.C:
			db := lmysql.Conn(dns)
			fmt.Println(db)
		case <-start:
			fmt.Println("time down")
			goto ForEnd
		default:

		}
	}
ForEnd:
	println("bye")
}
