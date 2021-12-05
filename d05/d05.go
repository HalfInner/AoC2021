package main

import (
	"AoC2021/aoc_fun"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	from_x int
	from_y int
	to_x   int
	to_y   int
}

type Data struct {
	records []Record
}

func parse(line string, data *Data) {
	nearby_line_str := strings.Split(line, " -> ")
	if len(nearby_line_str) != 2 {
		return
	}
	from_str := strings.Split(nearby_line_str[0], ",")
	to_str := strings.Split(nearby_line_str[1], ",")

	from_x, _ := strconv.ParseUint(from_str[0], 10, 64)
	from_y, _ := strconv.ParseUint(from_str[1], 10, 64)
	to_x, _ := strconv.ParseUint(to_str[0], 10, 64)
	to_y, _ := strconv.ParseUint(to_str[1], 10, 64)

	data.records = append(data.records, Record{int(from_x), int(from_y), int(to_x), int(to_y)})
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

	input := string(file)
	input_split := strings.Split(input, "\n")

	var data Data
	for _, line := range input_split {
		parse(line, &data)
	}

	return data
}

func count_points(data Data, check_diagonals bool) int {
	hash_key := func(x int, y int) int {
		return int(int64(x)<<32 | int64(y))
	}
	dir := func(a, b int) int {
		if a == b {
			return 0
		} else if a < b {
			return 1
		}
		return -1
	}

	points := make(map[int]int)
	for _, rec := range data.records {
		if !check_diagonals && !(rec.from_x == rec.to_x || rec.from_y == rec.to_y) {
			continue
		}
		x1, y1, x2, y2 := rec.from_x, rec.from_y, rec.to_x, rec.to_y
		for {
			points[hash_key(x1, y1)]++
			if x1 == x2 && y1 == y2 {
				break
			}

			x1 += dir(x1, x2)
			y1 += dir(y1, y2)
		}
	}

	sum := 0
	for _, val := range points {
		if val >= 2 {
			sum += 1
		}
	}

	return sum
}

func d05_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	return count_points(data, false)
}

func d05_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	return count_points(data, true)
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d05_1(data))
	log.Printf("02: %d", d05_2(data))
}
