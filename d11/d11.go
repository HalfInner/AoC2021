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
	octopuses []int
	flashed   []bool
	rows      int
	columns   int
}

func parse(line string, data *Data) {
	for _, r := range line {
		r_int := int(r - '0')
		data.octopuses = append(data.octopuses, r_int)
		data.flashed = append(data.flashed, false)
	}
	if data.columns == 0 {
		data.columns = len(data.octopuses)
	}
	data.rows++
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
	for _, input := range strings.Split(input_n, "\n") {
		parse(input, &data)
	}
	return data
}

func print2d(data []int, columns int) (str string) {
	for idx, el := range data {
		str += "\t" + strconv.Itoa(el)
		if (idx)%columns == columns-1 {
			str += "\n"
		}
	}
	return str
}

func d11_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	hash_key := func(x int, y int) int {
		return int(x<<32 | y)
	}

	y_last := data.rows - 1
	x_last := data.columns - 1
	marks := make(map[int]int)
	flash := func(x, y int) int {
		l_avail := x > 0
		r_avail := x < x_last
		u_avail := y > 0
		d_avail := y < y_last

		_flash := func(cond bool, x, y int) {
			if !cond {
				return
			}
			marks[hash_key(x, y)]++
		}
		_flash(true, x, y)
		_flash(l_avail, x-1, y)
		_flash(r_avail, x+1, y)
		_flash(u_avail, x, y-1)
		_flash(d_avail, x, y+1)
		_flash(l_avail && u_avail, x-1, y-1)
		_flash(l_avail && d_avail, x-1, y+1)
		_flash(r_avail && u_avail, x+1, y-1)
		_flash(r_avail && d_avail, x+1, y+1)
		return 0
	}

	flashes := 0
	for day := 0; day < 100; day++ {
		for y := 0; y < data.rows; y++ {
			for x := 0; x < data.columns; x++ {
				data.octopuses[data.columns*y+x]++
			}
		}

		for {
			to_continue := false
			for y := 0; y < data.rows; y++ {
				for x := 0; x < data.columns; x++ {
					if data.flashed[data.columns*y+x] {
						continue
					}
					if data.octopuses[data.columns*y+x] >= 10 {
						flash(x, y)
						data.flashed[data.columns*y+x] = true
						to_continue = true
					}
				}
			}
			if !to_continue {
				break
			}

			for key, val := range marks {
				x, y := key>>32, key&0xffffffff
				data.octopuses[data.columns*y+x] += val
				delete(marks, key)
			}
		}
		for y := 0; y < data.rows; y++ {
			for x := 0; x < data.columns; x++ {
				if data.octopuses[data.columns*y+x] > 9 {
					data.octopuses[data.columns*y+x] = 0
					data.flashed[data.columns*y+x] = false
					flashes++
				}
			}
		}

	}

	return flashes
}

func d11_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	hash_key := func(x int, y int) int {
		return int(x<<32 | y)
	}
	y_last := data.rows - 1
	x_last := data.columns - 1
	marks := make(map[int]int)
	flash := func(x, y int) int {
		l_avail := x > 0
		r_avail := x < x_last
		u_avail := y > 0
		d_avail := y < y_last

		_flash := func(cond bool, x, y int) {
			if !cond {
				return
			}
			marks[hash_key(x, y)]++
		}
		_flash(true, x, y)
		_flash(l_avail, x-1, y)
		_flash(r_avail, x+1, y)
		_flash(u_avail, x, y-1)
		_flash(d_avail, x, y+1)
		_flash(l_avail && u_avail, x-1, y-1)
		_flash(l_avail && d_avail, x-1, y+1)
		_flash(r_avail && u_avail, x+1, y-1)
		_flash(r_avail && d_avail, x+1, y+1)
		return 0
	}

	flashes_last := 0
	flashes := 0
	for day := 0; ; day++ {
		is_synchronised := true
		for y := 0; y < data.rows && is_synchronised; y++ {
			for x := 0; x < data.columns && is_synchronised; x++ {
				if data.octopuses[data.columns*y+x] > 0 {
					is_synchronised = false
				}
			}
		}
		if is_synchronised {
			return day
		}

		for y := 0; y < data.rows; y++ {
			for x := 0; x < data.columns; x++ {
				data.octopuses[data.columns*y+x]++
			}
		}

		flashes_last = 0
		for {
			flashes_last++
			to_continue := false
			for y := 0; y < data.rows; y++ {
				for x := 0; x < data.columns; x++ {
					if data.flashed[data.columns*y+x] {
						continue
					}
					if data.octopuses[data.columns*y+x] >= 10 {
						flash(x, y)
						data.flashed[data.columns*y+x] = true
						to_continue = true
					}
				}
			}
			if !to_continue {
				break
			}

			for key, val := range marks {
				x, y := key>>32, key&0xffffffff
				data.octopuses[data.columns*y+x] += val
				delete(marks, key)
			}
		}
		for y := 0; y < data.rows; y++ {
			for x := 0; x < data.columns; x++ {
				if data.octopuses[data.columns*y+x] > 9 {
					data.octopuses[data.columns*y+x] = 0
					data.flashed[data.columns*y+x] = false
					flashes++
				}

			}
		}

	}
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d11_1(data))
	data = read_data() // deep copy is Golang's cryptonite
	log.Printf("02: %d", d11_2(data))
}
