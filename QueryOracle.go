package main

import (
	"database/sql"
	"fmt"
	"log"
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

	db, err = sql.Open("godror", `user="UserID" password="PWD" connectString="SERVERNAME:1521/SERVICENAME"`)
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
	rows, err := db.Query("select rolename from psroleuser where roleuser = 'prashant.atman'")
	if err != nil {
		// fmt.Println("Error running PS query")
		// fmt.Println(err)
		check(err, "queryPSRoles")
		return
	}
	defer rows.Close()
	t := time.Now()
	const layout = "2006-01-02_150405.000000000"
	filename := "Roles_" + t.Format(layout) + ".txt"

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
	// fmt.Println(roles)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("\nFailed creating file: %s", err)
	}
	defer file.Close()
	// datawriter := bufio.NewWriter(file)
	// for _, data := range roles {
	// 	_, _ = datawriter.WriteString(data + "\n")
	// }
	// datawriter.Flush()
	len, err := file.WriteString(fmt.Sprintln(roles) + "\n")
	if err != nil {
		log.Fatalf("\nFailed writing to file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func queryPSOPRDEFN() {
	dbQUERY, err := db.Prepare("select operpswd, operpswdsalt from psoprdefn where oprid = :1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbQUERY.Close()
	rows, err := dbQUERY.Query("prashant.atman")
	if err != nil {
		fmt.Println("Error processing PSOPRDEFN query")
		fmt.Println(err)
		return
	}
	defer rows.Close()
	currentTime := time.Now()
	const layout = "2006-01-02_150405.000000000"
	filename := "QueryOP_" + currentTime.Format(layout) + ".txt"
	// text := "PUT SOMETHING HERE"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	var operPSWD, operPSWDSALT string
	for rows.Next() {
		rows.Scan(&operPSWD, &operPSWDSALT)
		text := operPSWD + " <:> " + operPSWDSALT
		_, err := file.WriteString(text)
		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}
		// fmt.Println("Hashed password: " + operPSWD + " <> SALT password " + operPSWDSALT)
	}
	fmt.Printf("\nFile Name: %s\n", file.Name())
	// fmt.Printf("\nLength: %d bytes", len)
}
