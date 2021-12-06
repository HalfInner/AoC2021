package main

import (
	"AoC2021/aoc_fun"
	"io/ioutil"
	"log"
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

func d06_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	for day := 0; day < 80; day++ {
		len := len(data.states)
		for idx := 0; idx < len; idx++ {
			s := data.states[idx]
			if s == 0 {
				s = 6 + 1
				data.states = append(data.states, 8)
			}
			s--
			data.states[idx] = s
		}
	}

	return len(data.states)
}

func d06_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	var phases [9]int
	for _, t := range data.states {
		phases[t]++
	}

	days := 256
	for day := 0; day < days; day++ {
		last := 0
		p_n := len(phases)
		for idx := 0; idx < p_n-1; idx++ {
			if idx == 0 {
				last = phases[idx]
			}
			phases[idx] = phases[idx+1]
		}
		phases[p_n-1] = last
		phases[p_n-3] += last
	}

	sum := 0
	for _, val := range phases {
		sum += val
	}
	return sum
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d06_1(data))
	log.Printf("02: %d", d06_2(data))
}
