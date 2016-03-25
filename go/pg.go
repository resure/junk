package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=resure dbname=resure sslmode=disable")
	checkErr(err)
	defer db.Close()

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT COUNT(*) FROM wikidot_hits")
	checkErr(err)

	for rows.Next() {
		var num int
		err = rows.Scan(&num)
		checkErr(err)
		fmt.Println(num)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
