package main

import (
	"AoC2021/aoc_fun"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Record struct {
	from, to string
}

type Data struct {
	records []Record
}

func parse(line string, data *Data) {
	r_str := strings.Split(line, "-")
	data.records = append(data.records, Record{r_str[0], r_str[1]})
}

func is_big_cave(s string) bool {
	return s[0] >= 'A' && s[0] <= 'Z'
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
	for _, line := range strings.Split(input_n, "\n") {
		parse(line, &data)
	}

	return data
}

type tips struct{ t []string }

func create_cyclic_graph(data Data) map[string]tips {

	graph := make(map[string]tips)
	for _, r := range data.records {
		// left side
		val, ok := graph[r.from]
		if !ok {
			val.t = make([]string, 0)
		}
		val.t = append(val.t, r.to)
		graph[r.from] = val

		// righ side
		val, ok = graph[r.to]
		if !ok {
			val.t = make([]string, 0)
		}
		val.t = append(val.t, r.from)
		graph[r.to] = val
	}

	return graph
}

func d12_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	graph := create_cyclic_graph(data)

	begin, end := "start", "end"
	visited := make(map[string]int)
	var search func(curr string) int
	search = func(curr string) int {
		if curr == end {
			return 1
		}

		visited[curr]++
		sum := 0
		for _, next := range graph[curr].t {
			if next == begin {
				continue
			}
			if !is_big_cave(next) && visited[next] > 0 {
				continue
			}
			sum += search(next)
		}
		visited[curr]--

		return sum
	}

	return search(begin)
}

func d12_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	graph := create_cyclic_graph(data)

	begin, end := "start", "end"

	type Doubled struct {
		key    string
		indent int
	}

	var doubled Doubled
	visited := make(map[string]int)

	var search func(curr string, deep int) int
	search = func(curr string, deep int) int {
		if curr == end {
			return 1
		}

		visited[curr]++
		sum := 0
		for _, next := range graph[curr].t {
			if next == begin {
				continue
			}

			if !is_big_cave(next) && visited[next] > 0 {
				if doubled.key != "" {
					continue
				}
				doubled = Doubled{next, deep}
			}

			sum += search(next, deep+1)

			if doubled.indent == deep {
				doubled.key = ""
			}
		}
		visited[curr]--

		return sum
	}
	return search(begin, 0)
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d12_1(data))
	log.Printf("02: %d", d12_2(data))
}
