package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ScanFile(filename string) []string {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return nil
	}
	defer f.Close()

	var lines []string = make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return nil
	}
	return lines
}

func ScanFileAsInt(filename string) []int {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return nil
	}
	defer f.Close()

	var lines []int = make([]int, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {

		var line string = sc.Text()
		intVar, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("ERROR reading as int: ", line)
			return nil
		}
		lines = append(lines, intVar)

	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return nil
	}
	return lines
}

//slices https://stackoverflow.com/a/21328187/835567
