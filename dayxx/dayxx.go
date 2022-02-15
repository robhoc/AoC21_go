package dayxx

import (
	"fmt"
	"time"
)

var day string = "dayxx"

var path string = day + "/inputs/input.txt"
var path01 string = day + "/inputs/input01.txt"
var path02 string = day + "/inputs/input02.txt"

func Solve() {
	var res01 string = Test01(path01)
	if res01 == "" {
		fmt.Println("Test 01 successful")

		start := time.Now()
		fmt.Println("Solution 01: " + Solve01(path))
		elapsed := time.Since(start)
		fmt.Printf("Part 01 took %s", elapsed)

		var res02 string = Test02(path02)
		if res02 == "" {
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
	fmt.Println("using path " + path)
	return ""
}

func Solve02(path string) string {
	fmt.Println("using path " + path)
	return ""
}

func Test01(path string) string {
	return Solve01(path)
}

func Test02(path string) string {
	return Solve02(path)
}
