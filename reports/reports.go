package reports

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Stats(results []string) (float64, float64, int) {
	var responses []string
	var statuscodes []string
	if len(results) == 0 {
		fmt.Println("No stats to create as http requests & results were empty")
		os.Exit(1)
	}
	for _, r := range results {
		fields := strings.Fields(r)
		responses = append(responses, fields[0])
		statuscodes = append(statuscodes, fields[2])
	}

	min, _ := strconv.ParseFloat(strings.TrimSuffix(responses[0], "s"), 64)
	max := min

	for _, t := range responses {
		trimed := strings.TrimSpace(t)
		num, err := strconv.ParseFloat(strings.TrimSuffix(trimed, "s"), 64)
		if err != nil {
			fmt.Println("Error parsing response time:", t, err)
			continue
		}
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	var not_ok_status []string
	for _, s := range statuscodes {
		if strings.TrimSpace(s) != "200" {
			not_ok_status = append(not_ok_status, s)
		}
	}
	return min, max, len(not_ok_status)
}

func GenerateCSV(filename string, data []string) {
	if filename == "" {
		filename = "surge-" + strconv.FormatInt(time.Now().Unix(), 10) + ".csv"
	} else {
		filename = filename
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	header := []string{"ResponseTime", "ResponseBytes", "HTTPStatusCode"}
	if err := writer.Write(header); err != nil {
		panic(err)
	}
	for _, row := range data {
		fields := strings.Fields(row)
		if err := writer.Write(fields); err != nil {
			panic(err)
		}
	}
	fmt.Printf("CSV report is in the file %s\n", filename)
}
