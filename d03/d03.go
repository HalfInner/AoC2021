package main

import (
	"AoC2021/aoc_fun"
	"bufio"
	"log"
	"os"
	"strconv"
)

type Record struct {
	power   string
	power_i int
}

func parse_move(line string) Record {
	v := line
	power_i, ok := strconv.ParseUint(v, 2, 64)
	if ok != nil {
		log.Fatal(ok)
	}
	return Record{v, int(power_i)}
}

func read_data() []Record {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []Record
	for scanner.Scan() {
		mv := parse_move(scanner.Text())
		data = append(data, mv)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func d03_1(data []Record) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	var values []int
	for _, bit := range data {
		for i, char := range bit.power {
			if len(values) <= i {
				values = append(values, 0)
			}
			switch char {
			case '0':
				values[i] -= 1
			case '1':
				values[i] += 1
			}

		}
	}

	var gamma_rate_rune []rune
	var epsilon_rate_rune []rune
	for _, val := range values {
		switch {
		case val == 0:
			log.Fatal("should not be")
		case val > 0:
			gamma_rate_rune = append(gamma_rate_rune, '1')
			epsilon_rate_rune = append(epsilon_rate_rune, '0')
		case val < 0:
			gamma_rate_rune = append(gamma_rate_rune, '0')
			epsilon_rate_rune = append(epsilon_rate_rune, '1')
		}
	}

	gamma_rate, ok := strconv.ParseUint(string(gamma_rate_rune), 2, 64)
	if ok != nil {
		log.Fatal(ok)
	}
	epsilon_rate, ok := strconv.ParseUint(string(epsilon_rate_rune), 2, 64)
	if ok != nil {
		log.Fatal(ok)
	}

	return int(gamma_rate) * int(epsilon_rate)
}

func d03_2(data []Record) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	oxygen_idxs := make(map[int]bool)
	scrubber_idxs := make(map[int]bool)

	for idx := 0; idx < len(data); idx++ {
		oxygen_idxs[idx] = true
		scrubber_idxs[idx] = true
	}

	for idx := 0; idx < len(data[0].power); idx++ {
		comparision := 0
		for oxygen_idx, _ := range oxygen_idxs {
			char := data[oxygen_idx].power[idx]
			switch char {
			case '0':
				comparision--
			case '1':
				comparision++
			}
		}
		for oxygen_idx, _ := range oxygen_idxs {
			if len(oxygen_idxs) == 1 {
				break
			}
			char := data[oxygen_idx].power[idx]
			if comparision < 0 && char == '1' || comparision >= 0 && char == '0' {
				delete(oxygen_idxs, oxygen_idx)
			}
		}

		comparision = 0
		for scrubber_idx, _ := range scrubber_idxs {
			char := data[scrubber_idx].power[idx]
			switch char {
			case '0':
				comparision--
			case '1':
				comparision++
			}
		}

		for scrubber_idx, _ := range scrubber_idxs {
			if len(scrubber_idxs) == 1 {
				break
			}
			char := data[scrubber_idx].power[idx]
			if comparision < 0 && char == '0' || comparision >= 0 && char == '1' {
				delete(scrubber_idxs, scrubber_idx)
			}
		}

	}

	oxygen_generator_rating := 0
	for oxygen_key, _ := range oxygen_idxs {
		oxygen_generator_rating = data[oxygen_key].power_i
	}

	scrubber_generator_rating := 0
	for scrubber_key, _ := range scrubber_idxs {
		scrubber_generator_rating = data[scrubber_key].power_i
	}

	return int(oxygen_generator_rating) * int(scrubber_generator_rating)
}

func main() {
	data := read_data()

	log.Printf("01: %d", d03_1(data))
	log.Printf("02: %d", d03_2(data))
}
