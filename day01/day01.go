package day01

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"hochstrat.info/aoc/lib"
)

var day string = "day01"

var path string = day + "/inputs/input.txt"
var path01 string = day + "/inputs/input01.txt"
var path02 string = day + "/inputs/input01.txt"

func Solve() {
	var res01 string = Test01(path01)
	if res01 == "7" {
		fmt.Println("Test 01 successful")

		start := time.Now()
		fmt.Println("Solution 01: " + Solve01(path))
		elapsed := time.Since(start)
		fmt.Printf("Part 01 took %s", elapsed)

		var res02 string = Test02(path02)
		if res02 == "5" {
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
	var lines []int = lib.ScanFileAsInt(path)
	var last int = math.MaxInt
	var count int = 0

	for _, line := range lines {

		if line > last {
			count++
		}
		last = line
	}
	return strconv.Itoa(count)
}

func Solve02(path string) string {
	var lines []int = lib.ScanFileAsInt(path)
	var last int = math.MaxInt
	var count int = 0

	for i, line := range lines {

		if i+2 <= len(lines)-1 {
			var cur int = line + lines[i+1] + lines[i+2]
			if cur > last {
				count++
			}
			last = cur
		}
	}
	return strconv.Itoa(count)
}

func Test01(path string) string {
	return Solve01(path01)
}

func Test02(path string) string {
	return Solve02(path)
}
