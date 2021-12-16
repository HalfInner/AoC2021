package main

import (
	"AoC2021/aoc_fun"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Data struct {
	transmission string
}

func parse_template(line string, data *Data) {
	data.transmission = line
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
	input_split := strings.Split(input, "\n")

	var data Data
	parse_template(input_split[0], &data)

	return data
}

func BitReverse32(x uint32) uint32 {
	x = (x&0x55555555)<<1 | (x&0xAAAAAAAA)>>1
	x = (x&0x33333333)<<2 | (x&0xCCCCCCCC)>>2
	x = (x&0x0F0F0F0F)<<4 | (x&0xF0F0F0F0)>>4
	x = (x&0x00FF00FF)<<8 | (x&0xFF00FF00)>>8
	return (x&0x0000FFFF)<<16 | (x&0xFFFF0000)>>16
}

type BitReader struct {
	hex_transmission string
	bit_cnt, bit_len int
}

func (br *BitReader) init(hex_transmission string) {
	br.hex_transmission = hex_transmission
	br.bit_cnt = 0
	br.bit_len = len(br.hex_transmission) * 4
}

func (br *BitReader) read_hex(bits int) int {
	val := 0
	read_bits := 0
	for read_bits < bits {
		val <<= 4
		hex_idx, bit_idx := br.bit_cnt/4, br.bit_cnt%4
		var b byte
		switch {
		case br.hex_transmission[hex_idx] >= '0' && br.hex_transmission[hex_idx] <= '9':
			b = br.hex_transmission[hex_idx] - '0'
		case br.hex_transmission[hex_idx] >= 'A' && br.hex_transmission[hex_idx] <= 'F':
			b = br.hex_transmission[hex_idx] - 'A' + 10
		default:
			log.Panic("?")
		}
		if b > 0xf {
			log.Panic("Read byte is greather than 0xf")
		}

		b_n := byte(0)
		left_bits_in_hex := 4 - bit_idx
		for ; left_bits_in_hex > 0 && read_bits < bits; left_bits_in_hex-- {
			b_n += b & (1 << (left_bits_in_hex - 1))
			read_bits++
			br.bit_cnt++
		}

		val += int(b_n)
		val >>= left_bits_in_hex
	}

	return val
}

func (br *BitReader) read_version() int {
	return br.read_hex(3)
}

func (br *BitReader) read_type() int {
	return br.read_hex(3)
}

func (br *BitReader) read_value() int {
	val := 0
	for {
		skip := br.read_hex(1) == 0
		val += br.read_hex(4)
		if skip {
			break
		}
		val <<= 4
	}
	br.align()
	return val
}

func (br *BitReader) align() {
	bit_idx := br.bit_cnt % 4
	if bit_idx > 0 {
		br.bit_cnt += 4 - bit_idx
	}
}

func (br *BitReader) nothing_to_read() bool {
	return br.bit_len == br.bit_cnt
}

func d16_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	var br BitReader
	br.init(data.transmission)
	br.read_version()
	type_ := br.read_type()
	literal_value_id := 4
	if type_ == literal_value_id {
		return br.read_value()
	}

	return -1
}

func d16_2(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	return -1
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d16_1(data))
	log.Printf("02: %d", d16_2(data))
}
