package main

import (
	"AoC2021/aoc_fun"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type Data struct {
	states []string
}

func parse(line string, data *Data) {
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
	var data Data
	data.states = append(data.states, strings.Split(input_n, "\n")...)

	return data
}

func d09_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	y_last := len(data.states) - 1
	x_last := len(data.states[0]) - 1
	is_lowest := func(x, y int) bool {
		l_avail := x != 0
		r_avail := x != x_last
		u_avail := y != 0
		d_avail := y != y_last

		l := !l_avail || data.states[y][x] < data.states[y][x-1]
		r := !r_avail || data.states[y][x] < data.states[y][x+1]
		u := !u_avail || data.states[y][x] < data.states[y-1][x]
		d := !d_avail || data.states[y][x] < data.states[y+1][x]
		return l && r && u && d
	}

	sum := 0
	for y, st := range data.states {
		for x, r := range st {
			if is_lowest(x, y) {
				sum += int(r-'0') + 1
			}
		}
	}

	return sum
}

func d09_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	y_last := len(data.states) - 1
	x_last := len(data.states[0]) - 1
	is_lowest := func(x, y int) bool {
		l_avail := x != 0
		r_avail := x != x_last
		u_avail := y != 0
		d_avail := y != y_last

		l := !l_avail || data.states[y][x] < data.states[y][x-1]
		r := !r_avail || data.states[y][x] < data.states[y][x+1]
		u := !u_avail || data.states[y][x] < data.states[y-1][x]
		d := !d_avail || data.states[y][x] < data.states[y+1][x]
		return l && r && u && d
	}

	hash_key := func(x int, y int) int {
		return int(x<<32 | y)
	}

	visited := make(map[int]bool)

	var measure_basin func(x, y int) int
	measure_basin = func(x, y int) int {
		visited[hash_key(x, y)] = true
		l_avail := x != 0 && data.states[y][x-1] != '9'
		r_avail := x != x_last && data.states[y][x+1] != '9'
		u_avail := y != 0 && data.states[y-1][x] != '9'
		d_avail := y != y_last && data.states[y+1][x] != '9'

		sum := 1
		if l_avail && (data.states[y][x] < data.states[y][x-1]) {
			if _, ok := visited[hash_key(x-1, y)]; !ok {
				sum += measure_basin(x-1, y)
			}
		}
		if r_avail && (data.states[y][x] < data.states[y][x+1]) {
			if _, ok := visited[hash_key(x+1, y)]; !ok {
				sum += measure_basin(x+1, y)
			}
		}
		if u_avail && (data.states[y][x] < data.states[y-1][x]) {
			if _, ok := visited[hash_key(x, y-1)]; !ok {
				sum += measure_basin(x, y-1)
			}
		}
		if d_avail && (data.states[y][x] < data.states[y+1][x]) {
			if _, ok := visited[hash_key(x, y+1)]; !ok {
				sum += measure_basin(x, y+1)
			}
		}

		return sum
	}

	var final_sums []int
	for y, st := range data.states {
		for x, _ := range st {
			if is_lowest(x, y) {
				sum := measure_basin(x, y)
				final_sums = append(final_sums, sum)
			}
		}
	}
	sort.Ints(final_sums)
	last := len(final_sums) - 1
	return final_sums[last] * final_sums[last-1] * final_sums[last-2]
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()

	log.Printf("01: %d", d09_1(data))
	log.Printf("02: %d", d09_2(data))
}
