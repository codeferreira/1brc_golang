package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		semicolon := strings.Index(row, ";")
		city := row[:semicolon]
		temp := row[semicolon+1:]
		fmt.Println(city, temp)
		return
	}
}
