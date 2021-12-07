package main

import (
	"AoC2021/aoc_fun"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	states []int
}

func parse(line string, data *Data) {
	for _, state := range strings.Split(line, ",") {
		res, _ := strconv.ParseUint(state, 10, 64)
		data.states = append(data.states, int(res))
	}
}

func read_data() Data {
	input_file_name := aoc_fun.GetDefaultInputFilePath()
	if len(os.Args) == 2 {
		input_file_name = os.Args[1]
	}

	file, err := ioutil.ReadFile(input_file_name)
	if err != nil {
		log.Fatal(err)
	}

	input_n := string(file)
	input := strings.Split(input_n, "\n")
	var data Data
	parse(input[0], &data)

	return data
}

func d07_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	min_fuel := math.MaxInt
	min_idx := -1
	for jdx := 0; jdx < len(data.states); jdx++ { // possible issue in logic
		sum := 0
		for idx := 0; idx < len(data.states) || idx == jdx; idx++ {
			sum += aoc_fun.Abs(data.states[jdx] - data.states[idx])
		}
		if sum < min_fuel {
			min_fuel = sum
			min_idx = jdx
		}
	}

	log.Print(min_idx)
	return min_fuel
}

func d07_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	var fuel_cost [2000]int // risky but fast
	fuel_cost[0] = 0
	for idx := 1; idx < len(fuel_cost); idx++ {
		curr_fuel := (idx)
		fuel_cost[idx] = curr_fuel + fuel_cost[idx-1]
	}

	min_fuel := math.MaxInt
	for jdx := 0; jdx < 2000; jdx++ {
		sum := 0
		for idx := 0; idx < len(data.states); idx++ {
			sum += fuel_cost[aoc_fun.Abs(jdx-data.states[idx])]
		}
		if sum < min_fuel {
			min_fuel = sum
		}
	}

	return min_fuel
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d07_1(data))
	log.Printf("02: %d", d07_2(data))
}
