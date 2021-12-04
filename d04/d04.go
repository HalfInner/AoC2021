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
	bingo         []int
	boards_shoots []map[int]bool
	boards        [][]int
}

func parse_bingo(line string, data *Data) {
	var v []int
	for _, number_str := range strings.Split(line, ",") {
		number, ok := strconv.ParseUint(number_str, 10, 64)
		if ok != nil {
			log.Fatal(ok)
		}
		v = append(v, int(number))
	}
	data.bingo = v
}

func parse_board(line string, data *Data) {
	var board []int
	board_shoot := make(map[int]bool)
	for _, number_str := range strings.Fields(line) {
		number, ok := strconv.ParseUint(number_str, 10, 64)
		if ok != nil {
			log.Fatal(ok)
		}
		board = append(board, int(number))
		board_shoot[int(number)] = false
	}
	data.boards = append(data.boards, board)
	data.boards_shoots = append(data.boards_shoots, board_shoot)
}

func read_data() Data {
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	input := string(file)
	input_split := strings.Split(input, "\n\n")

	var data Data
	for idx, board_str := range input_split {
		if idx == 0 {
			parse_bingo(input_split[0], &data)
			continue
		}
		parse_board(board_str, &data)
	}

	return data
}

func is_bingo(board []int, board_shoots map[int]bool) int {
	b := board
	s := board_shoots

	is_win := false
	switch {
	// rows
	case s[b[0]] && s[b[1]] && s[b[2]] && s[b[3]] && s[b[4]]:
		is_win = true
	case s[b[5]] && s[b[6]] && s[b[7]] && s[b[8]] && s[b[9]]:
		is_win = true
	case s[b[10]] && s[b[11]] && s[b[12]] && s[b[13]] && s[b[14]]:
		is_win = true
	case s[b[15]] && s[b[16]] && s[b[17]] && s[b[18]] && s[b[19]]:
		is_win = true
	case s[b[20]] && s[b[21]] && s[b[22]] && s[b[23]] && s[b[24]]:
		is_win = true
	// columns
	case s[b[0]] && s[b[5]] && s[b[10]] && s[b[15]] && s[b[20]]:
		is_win = true
	case s[b[1]] && s[b[6]] && s[b[11]] && s[b[16]] && s[b[21]]:
		is_win = true
	case s[b[2]] && s[b[7]] && s[b[12]] && s[b[17]] && s[b[22]]:
		is_win = true
	case s[b[3]] && s[b[8]] && s[b[13]] && s[b[18]] && s[b[23]]:
		is_win = true
	case s[b[4]] && s[b[9]] && s[b[14]] && s[b[19]] && s[b[24]]:
		is_win = true
	}
	if !is_win {
		return 0
	}
	sum := 0
	for key, val := range s {
		if !val {
			sum += key
		}
	}
	return sum
}

func d04_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	for idx, shoot := range data.bingo {
		for jdx, _ := range data.boards {
			data.boards_shoots[jdx][shoot] = true
		}

		if idx < 5 {
			continue
		}

		for jdx, _ := range data.boards {
			if res := is_bingo(data.boards[jdx], data.boards_shoots[jdx]); res > 0 {
				return res * shoot
			}
		}
	}

	return -1
}

func d04_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	win_boards := make(map[int]int)
	boards_number := len(data.boards)
	last_board := 0
	for idx, shoot := range data.bingo {
		for jdx, _ := range data.boards {
			if _, ok := win_boards[jdx]; ok {
				continue
			}
			data.boards_shoots[jdx][shoot] = true
		}

		if idx < 5 {
			continue
		}

		for jdx, _ := range data.boards {
			if _, ok := win_boards[jdx]; ok {
				continue
			}
			if res := is_bingo(data.boards[jdx], data.boards_shoots[jdx]); res > 0 {
				win_boards[jdx] = res * shoot
				last_board = jdx
			}
		}
		if len(win_boards) == boards_number {
			return win_boards[last_board]
		}
	}

	return -1
}

func main() {
	data := read_data()

	log.Printf("01: %d", d04_1(data))
	log.Printf("02: %d", d04_2(data))
}
