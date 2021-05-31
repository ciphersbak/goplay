package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/godror/godror"
	"github.com/joho/sqltocsv"
)

var db *sql.DB

type rolename struct {
	rolename string
}

func InitDB() {
	var err error

	db, err = sql.Open("godror", `user="USERID" password="PWD" connectString="SERVER:1521/SERVICENAME"`)
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
	waitGroup.Add(5)
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
	go func() {
		newQuery()
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
	filename := "Roles_" + t.Format(layout) + ".csv"

	var roles []rolename
	for rows.Next() {
		var role rolename
		err := rows.Scan(&role.rolename)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Printf("The role names are: %s\n", role)
		roles = append(roles, role)
		// result := strings.Join(roles[:], ",")
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
	err = file.Sync()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n===> done writing to file")
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
	filename := "QueryOP_" + currentTime.Format(layout) + ".csv"
	// text := "PUT SOMETHING HERE"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	var operPSWD, operPSWDSALT string
	for rows.Next() {
		rows.Scan(&operPSWD, &operPSWDSALT)
		text := operPSWD + ", " + operPSWDSALT
		_, err := file.WriteString(text)
		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}
		// fmt.Println("Hashed password: " + operPSWD + " <> SALT password " + operPSWDSALT)
	}
	fmt.Printf("\nFile Name: %s\n", file.Name())
	// fmt.Printf("\nLength: %d bytes", len)
}

func newQuery() {
	queryText := `
		SELECT A.SETID, A.BANK_CD, B.BRANCH_NAME, C.BANK_ACCT_KEY, A.BNK_ID_NBR, B.BRANCH_ID, C.BANK_ACCOUNT_NUM, C.DESCR AS "Bank Account Descr", 
		C.CURRENCY_CD, C.BUSINESS_UNIT_GL, C.ACCT_STATUS, B.DESCR "Branch Descr", C.RFC_CODE, C.PYCYCL_PAY_LIMIT, D.COUNTERPARTY, D.BANK_CD_CUST,
		A.DESCR, A.DESCRSHORT, A.DESCR_AC, A.BANK_TYPE, A.HOLIDAY_LIST_ID, A.CR_RATING_TYPE, A.CR_RATING, 
		A.BANK_STATUS, A.AR, A.AP, A.TR, A.BI, A.NETTING_STATUS, A.LEGAL_NET, A.DEALING_SW, A.ISSUER_SW, A.TREASURY_SW, A.BROKER_SW, A.TREASURY_CPTY_SW, 
		A.BANKING_SW, A.BANK_ID_QUAL, A.RECON_TRANS_CODE, A.IMMEDIATE_DEST, A.IMMEDIATE_ORIGIN, A.COUNTRY, A.ADDRESS1, A.ADDRESS2, 
		A.ADDRESS3, A.ADDRESS4, A.CITY, A.NUM1, A.NUM2, A.HOUSE_TYPE, A.ADDR_FIELD1, A.ADDR_FIELD2, A.ADDR_FIELD3, A.COUNTY, A.STATE, A.POSTAL, 
		A.GEO_CODE, A.IN_CITY_LIMIT, A.COUNTRY_CODE, A.PHONE, A.EXTENSION, A.FAX, A.ALC, A.PAYEE_VALIDATE_REQ, A.DESCRLONG,
		B.SETID, B.BANK_CD, B.BRANCH_NAME, B.BRANCH_STATUS, B.DESCR_AC, B.DESCRSHORT, B.BANK_ID_QUAL, B.VAT_RGSTRD_FLG, 
		B.VAT_RGSTRN_ID, B.VAT_SUSPENSION_FLG, B.COUNTRY, B.ADDRESS1, B.ADDRESS2, B.ADDRESS3, B.ADDRESS4, B.CITY, B.NUM1, B.NUM2, B.HOUSE_TYPE, 
		B.ADDR_FIELD1, B.ADDR_FIELD2, B.ADDR_FIELD3, B.COUNTY, B.STATE, B.POSTAL, B.GEO_CODE, B.IN_CITY_LIMIT, B.COUNTRY_CODE, B.PHONE, B.EXTENSION, 
		B.FAX, B.DESCRLONG, C.BANK_CD_CPTY, C.DESCRSHORT, C.BRANCH_NAME, C.PYMNT_HANDLING_CD, 
		C.PYMNT_OVRD_AMT, C.RECON_TYPE_FLG, C.RECON_HEADER, C.CUR_RT_TYPE, C.RATE_INDEX, C.SETTLEMENT_ID,  
		C.LAST_PYMNT_ID_USED, C.PYMNT_ID_REF_LEN, C.FRACT_RTG_NUM, C.DFI_ID_NUM, C.DFI_ID_QUAL, C.NPL_BNK_CHRT_VALUE, C.IU_PYBL_NPL_ACCT, 
		C.IU_RCVBL_NPL_ACCT, C.CSH_PL_CLRG_ACCT, C.CSH_PL_ACCT, C.PL_BNK_ACCT_FLG, C.AR, C.AP, C.TR, C.BI, C.EX, C.PP_SW, C.RCN_ACCTG_SW, C.ENABLE_WF_SW, 
		C.DEPOSIT_TYPE, C.FORECAST, C.TARGET_BALANCE, C.PRE_RECONCILE, C.CASH_CNTL_USE_FLG, C.DRAFT_CNTL_USE_FLG, C.DRAFT_CNTL_AR_FLG, C.DIT_USE_FLG, 
		C.CHECK_DIGIT, C.INT_BASIS, C.MARGIN_PTS, C.CREATION_DT, C.OPRID_ENTERED_BY, C.OPRID_LAST_UPDT, C.DEPOSIT_BU, C.DEBIT_RATE, C.CREDIT_RATE, 
		C.BANK_ACCT_QUAL, C.PREFERRED_LANGUAGE, C.PAYER_ID_NUM, C.BANK_ACCT_TYPE, C.DEFAULT_ACCT_SW, C.CHRG_BANK_CD, C.CHRG_BANK_ACCT_KEY, C.CREDIT_RT_INDEX, 
		C.CREDIT_RT_TYPE, C.CREDIT_MARGIN_PTS, C.DEBIT_RT_INDEX, C.DEBIT_RT_TYPE, C.DEBIT_MARGIN_PTS, C.SCHEDULE, C.POOL_ID, C.INTEREST_SW, 
		C.BCH_CHRG_PAYEE_FLG, C.IBAN_CHECK_DIGIT, C.PYCYCL_PAY_LMTCURR, C.SINGLE_PAY_LIMIT, C.SINGLE_PAY_LMTCURR, 
		C.IPAC_SENDER_DO, C.BNK_RCN_CONTROL, C.BNK_STMT_CONTROL, C.IBAN_ID,
		(SELECT RTRIM(XMLAGG(XMLELEMENT(E, CH.BANK_ACCT_LED_TYPE || ' - '|| X1.XLATLONGNAME || ' - ' || CH.ACCOUNT || ' - ' || CH.OPERATING_UNIT || ' - ' || CH.FUND_CODE || ' - ' || 
							CH.DEPTID || ' - ' || CH.CHARTFIELD1 || ' - ' || CH.CHARTFIELD2, '; ' || chr(13)).EXTRACT('//text()')ORDER BY CH.BANK_ACCT_LED_TYPE, CH.ACCOUNT)
							.getclobval(),',') AS "A2"
		 FROM (PS_BANK_ACCT_CHRT CH LEFT OUTER JOIN PSXLATITEM X1 ON X1.FIELDNAME = 'BANK_ACCT_LED_TYPE' AND X1.FIELDVALUE = CH.BANK_ACCT_LED_TYPE AND X1.EFF_STATUS = 'A')
		 WHERE D.SETID = CH.SETID AND D.BANK_CD = CH.BANK_CD AND D.BANK_ACCT_KEY = CH.BANK_ACCT_KEY AND D.BANK_CD_CPTY = CH.BANK_CD_CPTY
									 AND D.COUNTERPARTY = CH.COUNTERPARTY) AS "Chartfields",
		(SELECT RTRIM(XMLAGG(XMLELEMENT(E, M.PYMNT_METHOD || ' - ' || M.FORM_ID || ' - ' || M.EFT_LAYOUT_CD || ' - ' || M.STL_THROUGH || ' - ' || M.FORMAT_ID, ';  ' || chr(13)).EXTRACT('//text()')ORDER BY M.PYMNT_METHOD).getclobval(),',') AS "A1"
		 FROM PS_BANK_ACCT_MTHD M WHERE C.SETID = M.SETID AND C.BANK_CD = M.BANK_CD AND C.BANK_ACCT_KEY = M.BANK_ACCT_KEY AND C.BANK_CD_CPTY = M.BANK_CD_CPTY) AS "Payment Methods",
		(SELECT RTRIM(XMLAGG(XMLELEMENT(E, U.BNK_ID_NBR || ' - ' || U.COUNTRY || ' - ' || U.CURRENCY_CD || ' - ' || U.UN_PYMNT_TYPE || ' - ' || U.UN_PYMNT_METHOD || ' - ' || U.FAS_1006, '; ' || chr(13)).EXTRACT('//text()')ORDER BY U.BANK_CD, U.UN_PYMNT_TYPE).getclobval(),',') AS "A2"
		 FROM PS_UN_TR_BKID_MAP U WHERE C.SETID = U.SETID AND C.BANK_CD = U.BANK_CD) AS "Payment Channel Rules"
		FROM PS_BANK_CD_TBL A, PS_BANK_BRANCH_TBL B, PS_BANK_ACCT_DEFN C, PS_BANK_ACCT_CPTY D
		WHERE B.SETID = A.SETID 
		AND B.BANK_CD = A.BANK_CD 
		AND C.SETID = B.SETID 
		AND C.BANK_CD = A.BANK_CD 
		AND C.BRANCH_NAME = B.BRANCH_NAME 
		AND C.SETID = D.SETID
		AND C.BANK_CD = D.BANK_CD
		AND C.BANK_CD_CPTY = D.BANK_CD_CPTY
		AND C.BANK_ACCT_KEY = D.BANK_ACCT_KEY
		AND A.BANK_TYPE = 'E' 
		AND A.BANKING_SW = 'Y'
		AND A.SETID = :1`

	dbQUERY, err := db.Prepare(queryText)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbQUERY.Close()
	rows, err := dbQUERY.Query("SHARE")
	if err != nil {
		fmt.Println("Error processing BANK query")
		fmt.Println(err)
		return
	}
	defer rows.Close()
	currentTime := time.Now()
	const layout = "2006-01-02_150405.000000000"
	filename := "newQuery_" + currentTime.Format(layout) + ".csv"
	// text := "PUT SOMETHING HERE"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	fileerr := sqltocsv.WriteFile(filename, rows)
	if fileerr != nil {
		panic(fileerr)
	}
}
