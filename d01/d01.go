package main

import (
	"AoC2021/aoc_fun"
	"bufio"
	"log"
	"os"
	"strconv"
)

func read_data() []int {
	input_file_name := aoc_fun.GetDefaultInputFilePath()
	if len(os.Args) == 2 {
		input_file_name = os.Args[1]
	}

	file, err := os.Open(input_file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []int
	for scanner.Scan() {
		var current, _ = strconv.Atoi(scanner.Text())
		data = append(data, current)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func d01_1(data []int) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	increases := 0
	for i := 0; i < len(data)-1; i++ {
		if data[i] < data[i+1] {
			increases++
		}
	}

	return increases
}

func d01_2(data []int) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	increases := 0
	for i := 0; i < len(data)-3; i++ {
		curr_sum := data[i] + data[i+1] + data[i+2]
		next_sum := data[i+1] + data[i+2] + data[i+3]

		if curr_sum < next_sum {
			increases++
		}
	}
	return increases
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d01_1(data))
	log.Printf("02: %d", d01_2(data))
}
