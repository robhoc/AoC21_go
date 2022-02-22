package day02

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"hochstrat.info/aoc/lib"
)

var day string = "day02"

var path string = day + "/inputs/input.txt"
var path01 string = day + "/inputs/input01.txt"
var path02 string = day + "/inputs/input02.txt"

func Solve() {
	var res01 string = Test01(path01)
	if res01 == "150" {
		fmt.Println("Test 01 successful")

		start := time.Now()
		fmt.Println("Solution 01: " + Solve01(path))
		elapsed := time.Since(start)
		fmt.Printf("Part 01 took %s", elapsed)

		var res02 string = Test02(path02)
		if res02 == "900" {
			fmt.Println("Test 02 successful")
			start := time.Now()
			fmt.Println("Solution 2: " + Solve02(path))
			elapsed := time.Since(start)
			fmt.Printf("Part 02 took %s", elapsed)
		} else {
			fmt.Println("Part 02 failed with result", res02)
		}
	} else {
		fmt.Println("Test 01 failed with result", res01)
	}
}

func Solve01(path string) string {
	var lines []string = lib.ScanFile(path)
	var depth int = 0
	var horizontal int = 0

	for _, line := range lines {
		var iLine = strings.Split(line, " ")
		var instruction string = iLine[0]
		var value, _ = strconv.Atoi(iLine[1])

		if instruction == "up" {
			depth -= value
		} else if instruction == "down" {
			depth += value
		} else if instruction == "forward" {
			horizontal += value
		}
	}
	return strconv.Itoa(depth * horizontal)
}

func Solve02(path string) string {
	var lines []string = lib.ScanFile(path)

	var aim = 0

	var depth int = 0
	var horizontal int = 0

	for _, line := range lines {
		var iLine = strings.Split(line, " ")
		var instruction string = iLine[0]
		var value, _ = strconv.Atoi(iLine[1])
		if instruction == "up" {
			aim -= value
		} else if instruction == "down" {
			aim += value
		} else if instruction == "forward" {
			horizontal += value
			depth += aim * value
		}
	}
	return strconv.Itoa(depth * horizontal)
}

func Test01(path string) string {
	return Solve01(path)
}

func Test02(path string) string {
	return Solve02(path)
}
