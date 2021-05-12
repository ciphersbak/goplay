package main

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/godror/godror"
)

var db *sql.DB

type rolename struct {
	rolename string
}

func InitDB() {
	var err error

	db, err = sql.Open("godror", `user="UserID" password="pswd" connectString="SERVERNAME:1521/SERVICENAME"`)
	if err != nil {
		fmt.Println(err)
	}
	// defer db.Close()
	if err = db.Ping(); err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		return
	}
}

func check(e error, name string) {
	if e != nil {
		fmt.Println("Error in function: " + name)
		return
	}
}

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(4)
	fmt.Println("Starting sync calls...")
	start := time.Now()
	InitDB()

	go func() {
		queryDate()
		waitGroup.Done()
	}()

	go func() {
		queryDBName()
		waitGroup.Done()
	}()

	go func() {
		queryPSRoles()
		waitGroup.Done()
	}()

	go func() {
		queryPSOPRDEFN()
		waitGroup.Done()
	}()

	waitGroup.Wait()
	defer db.Close()
	elapsedTime := time.Since(start)
	fmt.Println("Total time for Execution: " + elapsedTime.String())
	time.Sleep(time.Second)
}

func queryDate() {

	rows, err := db.Query("select sysdate from dual")
	if err != nil {
		// fmt.Println("Error running Date query")
		// fmt.Println(err)
		check(err, "queryDate")
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {
		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
}

func queryDBName() {

	rows, err := db.Query("select ora_database_name from dual")
	if err != nil {
		fmt.Println("Error running DBName query")
		fmt.Println(err)
		// check(err, "queryDBName")
		return
	}
	// check(err, "queryDBName")
	defer rows.Close()

	var dbname string
	for rows.Next() {
		rows.Scan(&dbname)
	}
	fmt.Printf("The DB name is: %s\n", dbname)
}

func queryPSRoles() {
	rows, err := db.Query("select rolename from psroleuser where roleuser = 'VP1'")
	if err != nil {
		// fmt.Println("Error running PS query")
		// fmt.Println(err)
		check(err, "queryPSRoles")
		return
	}
	defer rows.Close()

	var roles []rolename
	for rows.Next() {
		var role rolename
		err := rows.Scan(&role.rolename)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Printf("The role names are: %s\n", role)
		roles = append(roles, role)
	}
}

func writeToFile(filename string, data string) {
	file, err := os.Create(filename)
	check(err, "writeToFile")
	defer file.Close()
}

func queryPSOPRDEFN() {
	dbQUERY, err := db.Prepare("select operpswd, operpswdsalt from psoprdefn where oprid = :1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbQUERY.Close()
	rows, err := dbQUERY.Query("VP1")
	if err != nil {
		fmt.Println("Error processing PSOPRDEFN query")
		fmt.Println(err)
		return
	}
	defer rows.Close()
	var operPSWD, operPSWDSALT string
	for rows.Next() {
		rows.Scan(&operPSWD, &operPSWDSALT)
		fmt.Println("Hashed password: " + operPSWD + " <> SALT password " + operPSWDSALT)
	}
}
