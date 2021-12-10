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

type Stack struct {
	capacity int
	vec      []rune
}

func (s *Stack) Size() int {
	return s.capacity
}
func (s *Stack) Empty() {
	s.capacity = 0
}

func (s *Stack) Push(value rune) {
	if s.capacity == len(s.vec) {
		s.vec = append(s.vec, value)
	} else {
		s.vec[s.capacity] = value
	}
	s.capacity++
}

func (s *Stack) Pop() {
	if s.capacity == 0 {
		log.Panic("Trying to pop from empty stack")
	}
	s.capacity--
}

func (s *Stack) Top() rune {
	if s.capacity == 0 {
		log.Panic("Trying to read from empty stack")
	}
	return s.vec[s.capacity-1]
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

	open_bracets := func(r rune) rune {
		switch r {
		case '(':
			return ')'
		case '<':
			return '>'
		case '[':
			return ']'
		case '{':
			return '}'
		default:
			return 'X'
		}
	}

	close_bracets := func(r rune) rune {
		switch r {
		case ')':
			return '('
		case '>':
			return '<'
		case ']':
			return '['
		case '}':
			return '{'
		default:
			return 'X'
		}
	}

	var stack Stack
	sum := 0
	for _, line := range data.states {
		stack.Empty()
		for _, b := range line {
			if val := open_bracets(b); val != 'X' {
				stack.Push(val)
			}
			if val := close_bracets(b); val != 'X' {
				front := stack.Top()
				if front == open_bracets(val) {
					stack.Pop()
				} else {
					sum += point_per_bracket(open_bracets(val))
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

	open_bracets := func(r rune) rune {
		switch r {
		case '(':
			return ')'
		case '<':
			return '>'
		case '[':
			return ']'
		case '{':
			return '}'
		default:
			return 'X'
		}
	}

	close_bracets := func(r rune) rune {
		switch r {
		case ')':
			return '('
		case '>':
			return '<'
		case ']':
			return '['
		case '}':
			return '{'
		default:
			return 'X'
		}
	}

	var sums []int
	var stack Stack
	for _, line := range data.states {
		stack.Empty()
		to_skip := false
		for _, b := range line {
			if val := open_bracets(b); val != 'X' {
				stack.Push(val)
			}
			if val := close_bracets(b); val != 'X' {
				if stack.Top() == open_bracets(val) {
					stack.Pop()
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
		for idx := 0; idx < stack.capacity; idx++ {
			partial_sum = partial_sum*5 + point_per_bracket(stack.vec[stack.capacity-idx-1])
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
