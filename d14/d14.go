package main

import (
	"AoC2021/aoc_fun"
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

type Rule struct {
	polimer, insertion string
}

type Data struct {
	template string
	rules    []Rule
}

func parse_template(line string, data *Data) {
	data.template = line
}

func parse_rule(line string, data *Data) {
	var polimer, insertion string
	fmt.Sscanf(line, "%s -> %s", &polimer, &insertion)
	data.rules = append(data.rules, Rule{polimer, insertion})
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
		parse_template(line, &data)
	}
	for _, line := range strings.Split(input_split[1], "\n") {
		parse_rule(line, &data)
	}

	return data
}

func d14_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	polymer := list.New()

	// create polymer
	counter := make(map[string]int)
	for _, r := range data.template {
		polymer.PushBack(string(r))
		counter[string(r)]++
	}

	// create map
	rules_mapping := make(map[string]string)
	for _, rule := range data.rules {
		rules_mapping[rule.polimer] = rule.insertion
	}

	for step := 0; step < 10; step++ {
		prev := polymer.Front()
		for {
			if prev == nil {
				break
			}
			next := prev.Next()
			if next == nil {
				break
			}

			rule := prev.Value.(string) + next.Value.(string)
			if insertion, ok := rules_mapping[rule]; ok {
				polymer.InsertAfter(insertion, prev)
				counter[insertion]++
			}
			prev = next
		}
	}

	min, max := math.MaxInt, math.MinInt
	for _, val := range counter {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return max - min
}

func d14_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	counter := make(map[string]int)
	for _, r := range data.template {
		counter[string(r)]++
	}
	for _, rule := range data.rules {
		counter[string(rule.polimer[0])]++
		counter[string(rule.polimer[1])]++
		counter[string(rule.insertion[0])]++
	}

	mapping_chars := make(map[int]int)
	number := len(counter)
	idx := 0
	for key, _ := range counter {
		next_rune := int('A' + idx)
		mapping_chars[int(key[0])] = next_rune
		idx++
	}

	code_mapping := make(map[int]int)
	for idx := range data.rules {
		a := mapping_chars[int(data.rules[idx].polimer[0])] - 'A'
		b := mapping_chars[int(data.rules[idx].polimer[1])] - 'A'
		c := mapping_chars[int(data.rules[idx].insertion[0])] - 'A'
		code_mapping[(a<<8)|(b&0xff)] = c
	}

	result := make([]int, number)
	template := ""
	for _, r := range data.template {
		cnv := mapping_chars[int(r)]
		template += string(cnv)
		result[cnv-'A']++
	}

	board_pairs := make([]int, number*number)
	for idx := 0; idx < len(template)-1; idx++ {
		a := int(template[idx] - 'A')
		b := int(template[idx+1] - 'A')
		board_pairs[a*number+b]++
	}

	board_help := make([]int, number*number)
	for step := 0; step < 40; step++ {
		for idx, val := range board_pairs {
			a := idx / number
			b := idx % number
			c, ok := code_mapping[a<<8|b]
			if !ok {
				continue
			}
			result[c] += val
			board_help[a*number+b] += -val
			board_help[a*number+c] += val
			board_help[c*number+b] += val
		}

		for idx := range board_pairs {
			board_pairs[idx] += board_help[idx]
			board_help[idx] = 0
		}
	}

	min, max := math.MaxInt, math.MinInt
	for _, val := range result {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return max - min
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d14_1(data))
	log.Printf("02: %d", d14_2(data))
}
