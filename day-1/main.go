package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input := getInput("input.txt")
	sum := 0

	for _, k := range input {
		lineResult := sumLine(k)
		println("line: ", k)
		println("lineResult: ", lineResult)
		println()
		sum += lineResult
	}
	fmt.Println("result: ", sum)
}

func sumLine(line string) int {
	leftVal := ""
	rightVal := ""

	forward := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	backward := strings.NewReplacer(
		"eno", "1",
		"owt", "2",
		"eerht", "3",
		"ruof", "4",
		"evif", "5",
		"xis", "6",
		"neves", "7",
		"thgie", "8",
		"enin", "9",
	)

	replacedForward := forward.Replace(line)

	lineBackward := Reverse(line)
	replacedBackward := backward.Replace(lineBackward)

	for _, y := range replacedForward {
		if unicode.IsDigit(y) {
			leftVal = string(y)
			break
		}
	}

	for _, y := range replacedBackward {
		if unicode.IsDigit(y) {
			rightVal = string(y)
			break
		}
	}

	combinedDigit := leftVal + rightVal
	result, err := strconv.ParseInt(combinedDigit, 10, 64)
	if err != nil {
		fmt.Println("error converting string to int")
		fmt.Println(err)
		return 0
	}
	intResult := int(result)

	return intResult
}

func getInput(filename string) []string {

	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
