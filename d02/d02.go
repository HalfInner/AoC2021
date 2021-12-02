package main

import (
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

func ParseMove(line string) Move {
	v := strings.Split(line, " ")
	far, _ := strconv.Atoi(v[1])
	return Move{v[0], far}
}

func read_data() []Move {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []Move
	for scanner.Scan() {
		mv := ParseMove(scanner.Text())
		data = append(data, mv)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func d02_1(data []Move) int {
	depth := 0
	horizontal := 0
	for i := 0; i < len(data); i++ {
		mv := data[i]
		if mv.direction == "forward" {
			horizontal += mv.far
		} else if mv.direction == "down" {
			depth += mv.far
		} else if mv.direction == "up" {
			depth -= mv.far
		}
	}

	return depth * horizontal
}

func d02_2(data []Move) int {
	depth := 0
	horizontal := 0
	aim := 0
	for i := 0; i < len(data); i++ {
		mv := data[i]
		if mv.direction == "forward" {
			horizontal += mv.far
			depth += aim * mv.far
		} else if mv.direction == "down" {
			aim += mv.far
		} else if mv.direction == "up" {
			aim -= mv.far
		}
	}

	return depth * horizontal
}

func main() {
	data := read_data()

	log.Printf("01: %d", d02_1(data))
	log.Printf("02: %d", d02_2(data))
}
