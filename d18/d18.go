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
	sfls []SnailFishList
}

func parse(line string, data *Data) {
	var sfl SnailFishList
	cur_nest := 0
	for _, r := range line {
		switch r {
		case ']':
			cur_nest--
		case '[':
			cur_nest++
		case ',':
			continue
		default:
			val, _ := strconv.Atoi(string(r))
			sfl.sfns = append(sfl.sfns, SnailFishNumber{val, cur_nest})
		}
	}
	data.sfls = append(data.sfls, sfl)
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
	var data Data
	for _, line := range strings.Split(input, "\n") {
		parse(line, &data)
	}

	return data
}

type SnailFishNumber struct {
	value, deep int
}

type SnailFishList struct {
	sfns []SnailFishNumber
}

func (sfl *SnailFishList) reduce() {
	reduction_happen := true
	for reduction_happen {
		reduction_happen = false
		for idx := 0; idx < len(sfl.sfns); idx++ {
			if sfl.sfns[idx].deep > 4 {
				if idx > 0 {
					sfl.sfns[idx-1].value += sfl.sfns[idx].value
				}
				if idx+2 < len(sfl.sfns) {
					sfl.sfns[idx+2].value += sfl.sfns[idx+1].value
				}
				sfl.sfns[idx+1].value = 0
				sfl.sfns[idx+1].deep--

				*sfl = sfl.remove(idx)
				reduction_happen = true
				continue
			}
		}
		for idx := 0; idx < len(sfl.sfns); idx++ {
			if sfl.sfns[idx].value >= 10 {
				a := sfl.sfns[idx].value / 2
				b := sfl.sfns[idx].value/2 + sfl.sfns[idx].value%2

				sfl.sfns[idx].value = a
				sfl.sfns[idx].deep++

				*sfl = sfl.insert(idx+1, SnailFishNumber{b, sfl.sfns[idx].deep})

				reduction_happen = true
				break
			}
		}
	}
}

func (sfl *SnailFishList) insert(index int, sfn SnailFishNumber) SnailFishList {
	var ret_sfl SnailFishList
	if len(sfl.sfns) < index || index < 0 {
		log.Panic("Adding out the range")
	}
	if len(sfl.sfns) == index {
		ret_sfl.sfns = append(sfl.sfns, sfn)
		return ret_sfl
	}
	ret_sfl.sfns = append(sfl.sfns[:index+1], sfl.sfns[index:]...)
	ret_sfl.sfns[index] = sfn
	return ret_sfl
}

func (sfl *SnailFishList) remove(index int) SnailFishList {
	var ret_sfl SnailFishList
	ret_sfl.sfns = append(sfl.sfns[:index], sfl.sfns[index+1:]...)
	return ret_sfl
}

func (sfl *SnailFishList) add(sfl2 *SnailFishList) SnailFishList {
	var sfl_ret SnailFishList
	total_size := len(sfl.sfns) + len(sfl2.sfns)
	for idx := 0; idx < total_size; idx++ {
		var sfn SnailFishNumber
		if idx < len(sfl.sfns) {
			sfn = sfl.sfns[idx]
		} else {
			sfn = sfl2.sfns[idx-len(sfl.sfns)]
		}
		if len(sfl.sfns) > 0 {
			sfn.deep++
		}
		sfl_ret.sfns = append(sfl_ret.sfns, sfn)
	}
	return sfl_ret
}

func (sfl *SnailFishList) magnitude() int {
	_magnitude := func(l, r int) int { return l*3 + r*2 }

	var sfl_mag SnailFishList
	sfl_mag.sfns = append(sfl_mag.sfns, sfl.sfns...)

	for deep := 4; deep > 0; {
		for idx := 0; idx < len(sfl.sfns)-1; idx++ {
			if sfl_mag.sfns[idx].deep == deep {
				sfl_mag.sfns[idx].value = _magnitude(sfl_mag.sfns[idx].value, sfl_mag.sfns[idx+1].value)
				sfl_mag.sfns[idx].deep--
				sfl_mag.remove(idx + 1)
				continue
			}
		}
		deep--
	}

	return sfl_mag.sfns[0].value
}

func d18_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	var res SnailFishList
	for _, sfl := range data.sfls {
		res = res.add(&sfl)
		res.reduce()
	}

	return res.magnitude()
}

func d18_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	max_magnitude := -1
	for idx := 0; idx < len(data.sfls); idx++ {
		for jdx := idx + 1; jdx < len(data.sfls); jdx++ {
			res := data.sfls[idx].add(&data.sfls[jdx])
			res.reduce()
			mag := res.magnitude()
			if mag > max_magnitude {
				max_magnitude = mag
			}

			res = data.sfls[jdx].add(&data.sfls[idx])
			res.reduce()
			mag = res.magnitude()
			if mag > max_magnitude {
				max_magnitude = mag
			}
		}
	}

	return max_magnitude
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d18_1(data))
	log.Printf("02: %d", d18_2(data))
}
