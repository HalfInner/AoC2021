package main

import (
	"AoC2021/aoc_fun"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Data struct {
	enhancement_algorithm [512]int
	visible_picture       []int
	columns, rows         int
}

func parse_algorithm(line string, data *Data) {
	line = strings.ReplaceAll(line, "\n", "")
	for idx, r := range line {
		if r == '#' {
			data.enhancement_algorithm[idx] = 1
		} else {
			data.enhancement_algorithm[idx] = 0
		}
	}
}
func parse_picture(line string, data *Data) {
	columns := strings.Index(line, "\n")
	line = strings.ReplaceAll(line, "\n", "")
	for _, r := range line {
		if r == '#' {
			data.visible_picture = append(data.visible_picture, 1)
		} else {
			data.visible_picture = append(data.visible_picture, 0)
		}
	}
	data.columns = columns
	data.rows = len(data.visible_picture) / columns
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
	input_split := strings.Split(input, "\n\n")

	var data Data
	parse_algorithm(input_split[0], &data)
	parse_picture(input_split[1], &data)

	return data
}

func enhence_pic(enhancement_algorithm [512]int, rows, columns, enhancements int, initial_pic []int) int {
	binary_matrix := func(x, y, rows, columns int, picture []int, second_time bool) int {
		y_last, x_last := rows-1, columns-1
		l_avail, r_avail, u_avail, d_avail := x > 0, x < x_last, y > 0, y < y_last

		point_to_idx := func(x, y int) int {
			return y*columns + x
		}
		_binary_matrix := func(cond bool, x, y int) int {
			if !cond || x >= columns || x < 0 || y >= rows || y < 0 {
				if enhancement_algorithm[0] == 1 && !second_time {
					return enhancement_algorithm[511]
				}
				return enhancement_algorithm[0] // what with out of border
			}
			return picture[point_to_idx(x, y)]
		}
		val := _binary_matrix(true, x, y) << 4
		val += _binary_matrix(l_avail, x-1, y) << 5
		val += _binary_matrix(r_avail, x+1, y) << 3
		val += _binary_matrix(u_avail, x, y-1) << 7
		val += _binary_matrix(d_avail, x, y+1) << 1
		val += _binary_matrix(l_avail && u_avail, x-1, y-1) << 8
		val += _binary_matrix(l_avail && d_avail, x-1, y+1) << 2
		val += _binary_matrix(r_avail && u_avail, x+1, y-1) << 6
		val += _binary_matrix(r_avail && d_avail, x+1, y+1) << 0
		return val
	}

	var out_picture []int
	extension := 2
	out_picture_rows, out_picture_column := rows+extension, columns+extension
	for y := 0; y < out_picture_rows; y++ {
		for x := 0; x < out_picture_column; x++ {
			code := binary_matrix(x-extension/2, y-extension/2, rows, columns, initial_pic, false)
			out_picture = append(out_picture, enhancement_algorithm[code])
		}
	}

	var out_picture_2 []int
	sum := 0
	for idx := 1; idx < enhancements; idx++ {
		sum = 0
		out_picture_rows_2, out_picture_column_2 := out_picture_rows+extension, out_picture_column+extension
		for y := 0; y < out_picture_rows_2; y++ {
			for x := 0; x < out_picture_column_2; x++ {
				code := binary_matrix(x-extension/2, y-extension/2, out_picture_rows, out_picture_column, out_picture, (idx%2) == 1)
				out_picture_2 = append(out_picture_2, enhancement_algorithm[code])
				sum += out_picture_2[len(out_picture_2)-1]
			}
		}
		out_picture, out_picture_2 = out_picture_2, make([]int, 0)
		out_picture_rows, out_picture_column = out_picture_rows_2, out_picture_column_2
	}

	return sum
}

func d20_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	return enhence_pic(data.enhancement_algorithm, data.rows, data.columns, 2, data.visible_picture)
}

func d20_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	return enhence_pic(data.enhancement_algorithm, data.rows, data.columns, 50, data.visible_picture)
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d20_1(data))
	log.Printf("02: %d", d20_2(data))
}
