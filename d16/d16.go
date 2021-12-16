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

type BitReaderStack struct {
	bit_cnt, bit_len int
}

type Stack struct {
	capacity int
	vec      []BitReaderStack
}

func (s *Stack) Size() int {
	return s.capacity
}
func (s *Stack) Empty() {
	s.capacity = 0
}

func (s *Stack) Push(value BitReaderStack) {
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

func (s *Stack) Top() BitReaderStack {
	if s.capacity == 0 {
		log.Panic("Trying to read from empty stack")
	}
	return s.vec[s.capacity-1]
}

type BitReader struct {
	hex_transmission           string
	bit_cnt, bit_len, vers_cnt int
	stack                      Stack
}

func (br *BitReader) init(hex_transmission string) {
	br.hex_transmission = hex_transmission
	br.bit_cnt = 0
	br.bit_len = len(br.hex_transmission) * 4
	br.vers_cnt = 0
	br.stack.Push(BitReaderStack{br.bit_cnt, br.bit_len})
}

func (br *BitReader) read_hex(bits int) int {
	if br.bit_cnt+bits >= br.bit_len {
		br.bit_cnt += bits
		return 0
	}
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

func (br *BitReader) get_vers_cnt() int {
	return br.vers_cnt
}
func (br *BitReader) read_version() int {
	ver := br.read_hex(3)
	br.vers_cnt += ver
	return ver
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
	// br.align()
	return val
}

func (br *BitReader) align() {
	bit_idx := br.bit_cnt % 4
	if bit_idx > 0 {
		br.bit_cnt += 4 - bit_idx
	}
}

func (br *BitReader) nothing_to_read() bool {
	return br.bit_cnt >= br.bit_len
}

func (br *BitReader) call_bit_reader_def() {
	br.call_bit_reader(br.bit_cnt)
}

func (br *BitReader) call_bit_reader(len int) {
	br.stack.Push(BitReaderStack{br.bit_cnt, br.bit_len})
	br.bit_len = br.bit_cnt + len
}

func (br *BitReader) return_bit_reader() int {
	cnt := br.bit_cnt
	br.bit_cnt, br.bit_len = br.stack.Top().bit_cnt, br.stack.Top().bit_len
	br.stack.Pop()
	return cnt
}

func (br *BitReader) read_operator() {
	switch br.read_hex(1) {
	case 0: // sum
		log.Print("Operator bits")
		length_in_bits := br.read_hex(15)
		br.call_bit_reader(length_in_bits)
		br.read_packets()
		br.bit_cnt = br.return_bit_reader()
	case 1:
		log.Print("Operator packets")
		number_of_packets := br.read_hex(11)
		for packet := 0; packet < number_of_packets; packet++ {
			br.read_packets()
		}
	case 2:
		log.Print("Operator packets")
		number_of_packets := br.read_hex(11)
		for packet := 0; packet < number_of_packets; packet++ {
			br.read_packets()
		}
	case 3:
		log.Print("Operator packets")
		number_of_packets := br.read_hex(11)
		for packet := 0; packet < number_of_packets; packet++ {
			br.read_packets()
		}
	case 4:
		log.Print("Operator packets")
		number_of_packets := br.read_hex(11)
		for packet := 0; packet < number_of_packets; packet++ {
			br.read_packets()
		}
	case 5:
		log.Print("Operator packets")
		number_of_packets := br.read_hex(11)
		for packet := 0; packet < number_of_packets; packet++ {
			br.read_packets()
		}
	case 6:
		log.Print("Operator packets")
		number_of_packets := br.read_hex(11)
		for packet := 0; packet < number_of_packets; packet++ {
			br.read_packets()
		}
	case 7:
		log.Print("Operator packets")
		number_of_packets := br.read_hex(11)
		for packet := 0; packet < number_of_packets; packet++ {
			br.read_packets()
		}
	}
}

func (br *BitReader) read_packets() {
	for !br.nothing_to_read() {
		ver := br.read_version()
		log.Printf("Read Version: %d", ver)
		packet_type := br.read_type()
		switch packet_type {
		case 4: // value
			val := br.read_value()
			log.Printf("Read Value: %d", val)
		default: // operator
			br.read_operator()
		}
	}
}

func (br *BitReader) start() {
	br.read_packets()
}

func d16_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	var br BitReader
	br.init(data.transmission)
	br.start()
	return br.get_vers_cnt()
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
