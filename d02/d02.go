package main

import (
	"AoC2021/aoc_fun"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	direction string
	far       int
}

func parse_move(line string) Move {
	v := strings.Split(line, " ")
	far, _ := strconv.Atoi(v[1])
	return Move{v[0], far}
}

func read_data() []Move {
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
	var data []Move
	for scanner.Scan() {
		mv := parse_move(scanner.Text())
		data = append(data, mv)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func d02_1(data []Move) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	depth := 0
	horizontal := 0
	for _, mv := range data {
		switch mv.direction {
		case "forward":
			horizontal += mv.far
		case "down":
			depth += mv.far
		case "up":
			depth -= mv.far
		}
	}

	return depth * horizontal
}

func d02_2(data []Move) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	depth := 0
	horizontal := 0
	aim := 0
	for _, mv := range data {
		switch mv.direction {
		case "forward":
			horizontal += mv.far
			depth += aim * mv.far
		case "down":
			aim += mv.far
		case "up":
			aim -= mv.far
		}
	}

	return depth * horizontal
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d02_1(data))
	log.Printf("02: %d", d02_2(data))
}
