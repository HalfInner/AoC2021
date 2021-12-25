package main

import (
	"AoC2021/aoc_fun"
	"container/heap"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

type Columns struct {
	columns []int
}

type Data struct {
	rows []Columns
}

func parse(line string, data *Data) {
	v := make([]int, len(line))
	for idx := range line {
		v[idx] = int(line[idx] - '0')
	}

	data.rows = append(data.rows, Columns{v})
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
	for _, line := range strings.Fields(input_split[0]) {
		parse(line, &data)
	}

	return data
}

type Stack struct {
	capacity         int
	d_idx, d_idx_rev []int
	d                *[]int
}

func (s *Stack) Size() int {
	return s.capacity
}
func (s *Stack) Empty() {
	s.capacity = 0
}

func (s *Stack) Push(value interface{}) {
	if s.capacity == len(s.d_idx) {
		s.d_idx = append(s.d_idx, value.(int))
	} else {
		s.d_idx[s.capacity] = value.(int)
	}
	s.capacity++
}

func (s *Stack) Pop() interface{} {
	if s.capacity == 0 {
		log.Panic("Trying to pop from empty stack")
	}
	s.capacity--
	return s.d_idx[s.capacity]
}

func (s *Stack) Len() int           { return s.Size() }
func (s *Stack) Less(i, j int) bool { return (*s.d)[s.d_idx[i]] < (*s.d)[s.d_idx[j]] }
func (s *Stack) Swap(i, j int) {
	s.d_idx_rev[s.d_idx[i]], s.d_idx_rev[s.d_idx[j]] = s.d_idx_rev[s.d_idx[j]], s.d_idx_rev[s.d_idx[i]]
	s.d_idx[i], s.d_idx[j] = s.d_idx[j], s.d_idx[i]
}

func low_risk_route(map_multiplier int, data Data) int {
	const INF int = math.MaxInt - 1

	hash_key := func(x, y int) int {
		return x<<32 | y
	}

	real_width := len(data.rows[0].columns) * map_multiplier
	real_height := len(data.rows) * map_multiplier
	key_point := func(hash_key int) (x, y int) {
		return hash_key >> 32, hash_key & 0xffffffff
	}

	point_idx := func(x, y int) int {
		return y*real_width + x
	}

	idx_key := func(key int) int {
		return (key%real_width)<<32 | (key/real_width)&0xffffff
	}

	Q := make(map[int]int)
	d := make([]int, real_height*real_width)
	d_idx := make([]int, real_height*real_width)
	d_idx_rev := make([]int, real_height*real_width)
	for idx := 0; idx < real_height*real_width; idx++ {
		d[idx], d_idx[idx], d_idx_rev[idx] = INF, idx, idx
		x, y := idx%real_width, idx/real_width
		Q[hash_key(x, y)] = 1
	}
	d[0] = 0

	risk := func(x, y int) int {
		initial_width, initial_height := real_width/map_multiplier, real_height/map_multiplier
		initial_x, initial_y := x%initial_width, y%initial_height
		initial_risk := data.rows[initial_y].columns[initial_x]
		risk_shifts := x/initial_width + y/initial_height
		return (((initial_risk - 1) + risk_shifts) % 9) + 1
	}

	var s Stack
	s.capacity = len(d)
	s.d_idx = d_idx[:]
	s.d_idx_rev = d_idx_rev[:]
	s.d = &d
	heap.Init(&s)
	for range d {
		node := idx_key(heap.Pop(&s).(int))
		delete(Q, node)
		x, y := key_point(node)
		update := func(xn, yn, x, y int) {
			_, ok := Q[hash_key(xn, yn)]
			if ok && d[point_idx(xn, yn)] > d[point_idx(key_point(node))]+risk(xn, yn) {
				d[point_idx(xn, yn)] = d[point_idx(key_point(node))] + risk(xn, yn)
				heap.Fix(&s, d_idx_rev[point_idx(xn, yn)])
			}
		}
		if x > 0 {
			update(x-1, y, x, y)
		}
		if x < real_width-1 {
			update(x+1, y, x, y)
		}
		if y > 0 {
			update(x, y-1, x, y)
		}
		if y < real_height-1 {
			update(x, y+1, x, y)
		}
	}
	return d[point_idx(real_height-1, real_width-1)]
}

func d15_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	return low_risk_route(1, data)
}

func d15_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	return low_risk_route(5, data)
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d15_1(data))
	log.Printf("02: %d", d15_2(data))
}
