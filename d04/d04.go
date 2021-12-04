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
	boards_shoots []map[int]int
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
	board_shoot := make(map[int]int)
	for idx, number_str := range strings.Fields(line) {
		number, ok := strconv.ParseUint(number_str, 10, 8)
		if ok != nil {
			log.Fatal(ok)
		}
		board = append(board, int(number))
		board_shoot[int(number)] = idx
	}
	data.boards = append(data.boards, board)
	data.boards_shoots = append(data.boards_shoots, board_shoot)
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
	for idx, board_str := range input_split {
		if idx == 0 {
			parse_bingo(input_split[0], &data)
			continue
		}
		parse_board(board_str, &data)
	}

	return data
}

func is_bingo(board []int) int {
	b := board

	is_win := false
	switch {
	// rows
	case b[0]+b[1]+b[2]+b[3]+b[4] == -5:
		is_win = true
	case b[5]+b[6]+b[7]+b[8]+b[9] == -5:
		is_win = true
	case b[10]+b[11]+b[12]+b[13]+b[14] == -5:
		is_win = true
	case b[15]+b[16]+b[17]+b[18]+b[19] == -5:
		is_win = true
	case b[20]+b[21]+b[22]+b[23]+b[24] == -5:
		is_win = true
	// columns
	case b[0]+b[5]+b[10]+b[15]+b[20] == -5:
		is_win = true
	case b[1]+b[6]+b[11]+b[16]+b[21] == -5:
		is_win = true
	case b[2]+b[7]+b[12]+b[17]+b[22] == -5:
		is_win = true
	case b[3]+b[8]+b[13]+b[18]+b[23] == -5:
		is_win = true
	case b[4]+b[9]+b[14]+b[19]+b[24] == -5:
		is_win = true
	}
	if !is_win {
		return 0
	}

	sum := 0
	for _, val := range b {
		if val > 0 {
			sum += val
		}
	}
	return sum
}

func d04_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	for idx, shoot := range data.bingo {
		for jdx, board := range data.boards {
			if board_idx, ok := data.boards_shoots[jdx][shoot]; ok {
				board[board_idx] = -1
			}
		}

		if idx < 5 {
			continue
		}

		for _, board := range data.boards {
			if res := is_bingo(board); res > 0 {
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
		for jdx, board := range data.boards {
			if _, ok := win_boards[jdx]; ok {
				continue
			}
			if board_idx, ok := data.boards_shoots[jdx][shoot]; ok {
				board[board_idx] = -1
			}
		}

		if idx < 5 {
			continue
		}

		for jdx, board := range data.boards {
			if _, ok := win_boards[jdx]; ok {
				continue
			}
			if res := is_bingo(board); res > 0 {
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
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d04_1(data))
	log.Printf("02: %d", d04_2(data))
}
