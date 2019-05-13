package main

import (
    "encoding/csv"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/tealeg/xlsx"
)

type masterTestCase struct {
    ReqDesc  string
    ReqID    string
    FRICEID  string
    BPID     string
    BPName   string
    TestName string
    Desc     string
}

func main() {
    sDir := "/Users/kpanigrahi001/Downloads/MasterTestCases/AllMasterTestCases"
    files, err := ioutil.ReadDir(sDir)
    if err != nil {
        log.Fatal(err)
    }

    file, err := os.Create("/Users/kpanigrahi001/Downloads/MasterTestCases/masterTestcases.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()
    writer.Write([]string{"File Name", "Sheet Name", "Requirement Description", "Requirement ID", "FRICE ID", "BP ID", "BP Name", "Test Name", "Test Description"})

    for _, f := range files {
        if filepath.Ext(f.Name()) == ".xlsx" {
            fmt.Println(filepath.Join(f.Name()))
            excelFileName := filepath.Join(sDir, f.Name())
            xlFile, err := xlsx.OpenFile(excelFileName)
            if err != nil {
                log.Fatal(err)
            }
            for _, sheet := range xlFile.Sheets {
                if len(sheet.Rows) >= 1 && len(sheet.Rows[0].Cells) >= 1 {
                    fmt.Println("Sheet name:", sheet.Name)
                    iLineNbr := 1
                    if sheet.Rows[0].Cells[0].String() == "Requirement Description" && strings.Contains(sheet.Rows[iLineNbr].Cells[8].String(), "TC_") {
                        if sheet.Rows[iLineNbr].Cells[1].String() == "Requirement ID" {
                            iLineNbr = 2
                        }
                        masterTestCaseCur := masterTestCase{
                            ReqDesc:  sheet.Rows[iLineNbr].Cells[0].String(),
                            ReqID:    sheet.Rows[iLineNbr].Cells[1].String(),
                            FRICEID:  sheet.Rows[iLineNbr].Cells[2].String(),
                            BPID:     sheet.Rows[iLineNbr].Cells[5].String(),
                            BPName:   sheet.Rows[iLineNbr].Cells[6].String(),
                            TestName: sheet.Rows[iLineNbr].Cells[8].String(),
                            Desc:     sheet.Rows[iLineNbr].Cells[9].String()}
                        writer.Write([]string{f.Name(), sheet.Name, masterTestCaseCur.ReqDesc, masterTestCaseCur.ReqID, masterTestCaseCur.FRICEID, masterTestCaseCur.BPID, masterTestCaseCur.BPName, masterTestCaseCur.TestName, masterTestCaseCur.Desc})
                    }
                }
            }
        }
    }
}
