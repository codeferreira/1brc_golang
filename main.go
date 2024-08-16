package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	file, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make(map[string]Measurement)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		semicolon := strings.Index(row, ";")
		city := row[:semicolon]
		rawTemp := row[semicolon+1:]

		temp, _ := strconv.ParseFloat(rawTemp, 64)
		measurement, ok := data[city]
		if !ok {
			measurement = Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {

			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}

		data[city] = measurement
	}

	locations := make([]string, 0, len(data))
	for name := range data {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	for _, name := range locations {
		measurement := data[name]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f, ",
			name, measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)
	}
}
