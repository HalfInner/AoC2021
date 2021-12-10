package main

import (
	"AoC2021/aoc_fun"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Record struct {
	signal []string
	expect []string
}

type Data struct {
	states []Record
}

func parse(line string, data *Data) {
	if len(line) == 0 {
		return
	}
	parts := strings.Split(line, " | ")
	signal := strings.Fields(parts[0])
	expect := strings.Fields(parts[1])

	data.states = append(data.states, Record{signal, expect})
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

func d08_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	unique_numbers := 0
	for _, record := range data.states {
		signals := record.expect
		for _, signal := range signals {
			switch len(signal) {
			case 2, 4, 3, 7:
				unique_numbers++
			}
		}
	}

	return unique_numbers
}

func d08_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	// static line analisys:
	//	 a=>8 b=>6 c=>8 d=>7 e=>4 f=>9 g=>7
	make_mapping := func(a, b, c, d, e, f, g rune) map[rune]rune {
		mapping := make(map[rune]rune)
		mapping[a] = 'a'
		mapping[b] = 'b'
		mapping[c] = 'c'
		mapping[d] = 'd'
		mapping[e] = 'e'
		mapping[f] = 'f'
		mapping[g] = 'g'

		return mapping
	}

	map_signal := func(mapping map[rune]rune, signal string) string {
		runeArray := []rune(signal)
		for idx, s := range runeArray {
			runeArray[idx] = mapping[s]
		}
		return aoc_fun.SortString(string(runeArray))
	}

	get_number_from_display := func(signal string) int {
		switch signal {
		case "abcefg":
			return 0
		case "cf":
			return 1
		case "acdeg":
			return 2
		case "acdfg":
			return 3
		case "bcdf":
			return 4
		case "abdfg":
			return 5
		case "abdefg":
			return 6
		case "acf":
			return 7
		case "abcdefg":
			return 8
		case "abcdfg":
			return 9
		default:
			return -1
		}
	}

	char_c_from_one := func(signal string, f rune) rune {
		if len(signal) != 2 {
			log.Panic("Wrong length for one")
		}
		for _, r := range signal {
			if r != f {
				return r
			}
		}
		return 'P'
	}

	char_a_from_seven := func(signal string, c, f rune) rune {
		if len(signal) != 3 {
			log.Panic("Wrong length for one")
		}
		for _, r := range signal {
			if r != c && r != f {
				return r
			}
		}
		return 'P'
	}

	char_d_from_four := func(signal string, b, c, f rune) rune {
		if len(signal) != 4 {
			log.Panic("Wrong length for one")
		}
		for _, r := range signal {
			if r != b && r != c && r != f {
				return r
			}
		}
		return 'P'
	}

	char_g_from_eight := func(signal string, a, b, c, d, e, f rune) rune {
		if len(signal) != 7 {
			log.Panic("Wrong length for one")
		}
		for _, r := range signal {
			if r != a && r != b && r != c && r != d && r != e && r != f {
				return r
			}
		}
		return 'P'
	}

	sum := 0
	for _, record := range data.states {
		mapping := make(map[rune]int)
		s1, s4, s7, s8 := "", "", "", ""
		for _, signal := range record.signal {
			switch len(signal) {
			case 2:
				s1 = signal
			case 4:
				s4 = signal
			case 3:
				s7 = signal
			case 7:
				s8 = signal
			}
			for _, d := range signal {
				mapping[d]++
			}
		}

		var a, b, c, d, e, f, g rune
		for key, val := range mapping {
			switch val {
			case 9:
				f = key
			case 4:
				e = key
			case 6:
				b = key
			default:
				continue
			}
		}

		c = char_c_from_one(s1, f)
		a = char_a_from_seven(s7, c, f)
		d = char_d_from_four(s4, b, c, f)
		g = char_g_from_eight(s8, a, b, c, d, e, f)

		final_mapping := make_mapping(a, b, c, d, e, f, g)
		result := get_number_from_display(map_signal(final_mapping, record.expect[3]))
		result += get_number_from_display(map_signal(final_mapping, record.expect[2])) * 10
		result += get_number_from_display(map_signal(final_mapping, record.expect[1])) * 100
		result += get_number_from_display(map_signal(final_mapping, record.expect[0])) * 1000
		sum += result
	}

	return sum
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()

	log.Printf("01: %d", d08_1(data))
	log.Printf("02: %d", d08_2(data))
}
