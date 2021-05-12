package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func main() {
	db, err := sql.Open("godror", `user="YOURID" password="YOURPWD" connectString="SERVER:1521/SERVICENAME"`)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {
		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
}
