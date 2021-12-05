package main

import (
	"AoC2021/aoc_fun"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	var from_x, from_y, to_x, to_y int
	fmt.Sscanf(line, "%d,%d -> %d,%d", &from_x, &from_y, &to_x, &to_y)

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

func count_points(data Data, check_diagonals bool, use_array bool) int {
	hash_key := func(x int, y int) int {
		return int(x<<32 | y)
	}
	dir := func(a, b int) int {
		if a == b {
			return 0
		} else if a < b {
			return 1
		}
		return -1
	}

	points_hash := make(map[int]int)
	points_array := make([][]int, 1000)
	for idx := range points_array {
		points_array[idx] = make([]int, 1000)
	}

	sum := 0
	for _, rec := range data.records {
		if !check_diagonals && !(rec.from_x == rec.to_x || rec.from_y == rec.to_y) {
			continue
		}
		x1, y1, x2, y2 := rec.from_x, rec.from_y, rec.to_x, rec.to_y
		for {
			if use_array {
				points_array[x1][y1]++
				if points_array[x1][y1] == 2 {
					sum++
				}
			} else {
				hash := hash_key(x1, y1)
				if el := points_hash[hash]; el < 2 {
					points_hash[hash] = el + 1
					if el+1 == 2 {
						sum++
					}
				}
			}

			if x1 == x2 && y1 == y2 {
				break
			}

			x1 += dir(x1, x2)
			y1 += dir(y1, y2)
		}
	}

	return sum
}

func d05_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	return count_points(data, false, true)
}

func d05_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	return count_points(data, true, true)
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d05_1(data))
	log.Printf("02: %d", d05_2(data))
}
