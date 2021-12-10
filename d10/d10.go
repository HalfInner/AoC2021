package main

import (
	"AoC2021/aoc_fun"
	"container/list"
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

func d10_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	point_per_bracket := func(c rune) int {
		switch c {
		case ')':
			return 3
		case ']':
			return 57
		case '}':
			return 1197
		case '>':
			return 25137
		default:
			log.Panic("no rune")
			return -1
		}
	}

	open_bracets := map[rune]rune{
		'(': ')',
		'<': '>',
		'[': ']',
		'{': '}',
	}

	close_bracets := map[rune]rune{
		')': '(',
		'>': '<',
		']': '[',
		'}': '{',
	}

	stack := list.New()
	sum := 0
	for _, line := range data.states {
		for _, b := range line {
			if val, ok := open_bracets[b]; ok {
				stack.PushFront(val)
			}
			if val, ok := close_bracets[b]; ok {
				front := stack.Front()
				if front.Value.(rune) == open_bracets[val] {
					stack.Remove(front)
				} else {
					sum += point_per_bracket(open_bracets[val])
					break
				}
			}
		}
	}

	return sum
}

func d10_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	point_per_bracket := func(c rune) int {
		switch c {
		case ')':
			return 1
		case ']':
			return 2
		case '}':
			return 3
		case '>':
			return 4
		default:
			log.Panicf("No rune '%c'", c)
			return -1
		}
	}

	open_bracets := map[rune]rune{
		'(': ')',
		'<': '>',
		'[': ']',
		'{': '}',
	}

	close_bracets := map[rune]rune{
		')': '(',
		'>': '<',
		']': '[',
		'}': '{',
	}

	var sums []int
	for _, line := range data.states {
		stack := list.New()
		to_skip := false
		for _, b := range line {
			if val, ok := open_bracets[b]; ok {
				stack.PushFront(val)
			}
			if val, ok := close_bracets[b]; ok {
				front := stack.Front()
				if front.Value.(rune) == open_bracets[val] {
					stack.Remove(front)
				} else {
					to_skip = true
					break
				}
			}
		}
		if to_skip {
			continue
		}

		partial_sum := 0
		for last := stack.Front(); last != nil; last = last.Next() {
			partial_sum = partial_sum*5 + point_per_bracket(last.Value.(rune))
		}
		if partial_sum > 0 {
			sums = append(sums, partial_sum)
		}
	}

	sort.Ints(sums)
	medium_sum := sums[len(sums)/2]
	return medium_sum
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d10_1(data))
	log.Printf("02: %d", d10_2(data))
}
