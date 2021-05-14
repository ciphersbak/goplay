package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func createFile(filename, text string) {
	fmt.Printf("Writing to a file\n")
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	len, err := file.WriteString(text)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func main() {

	currentTime := time.Now()
	const layout = "2006-01-02_150405.000000000"
	filename := "TestFile_" + currentTime.Format(layout) + ".txt"
	text := "PUT SOMETHING HERE"
	createFile(filename, text)

}
