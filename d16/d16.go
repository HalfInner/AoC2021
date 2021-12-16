package main

import (
	"AoC2021/aoc_fun"
	"io/ioutil"
	"log"
	"math"
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
	bit_cnt, bit_len, bit_cnt_tmp int
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
	hex_transmission                        string
	bit_cnt, bit_len, vers_cnt, bit_cnt_tmp int
	stack                                   Stack
	debug_mode                              bool
}

func (br *BitReader) init(hex_transmission string) {
	br.hex_transmission = hex_transmission
	br.bit_cnt = 0
	br.bit_len = len(br.hex_transmission) * 4
	br.vers_cnt = 0
	br.bit_cnt_tmp = 0
	br.debug_mode = false
	br.stack.Push(BitReaderStack{br.bit_cnt, br.bit_len, br.bit_cnt_tmp})
}

func (br *BitReader) Printf(format string, a ...interface{}) {
	if br.debug_mode {
		log.Printf(format, a...)
	}
}

func (br *BitReader) read_hex(bits int) int {
	val := 0
	read_bits := 0
	for read_bits < bits && br.bit_cnt < br.bit_len {
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
		for ; left_bits_in_hex > 0 && read_bits < bits && br.bit_cnt < br.bit_len; left_bits_in_hex-- {
			b_n += b & (1 << (left_bits_in_hex - 1))
			read_bits++
			br.advance_cnt(1)
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
	return val
}

func (br *BitReader) nothing_to_read() bool {
	return br.bit_cnt >= br.bit_len
}

func (br *BitReader) call_bit_reader(len int) {
	br.stack.Push(BitReaderStack{br.bit_cnt, br.bit_len, br.bit_cnt_tmp})
	br.bit_len = br.bit_cnt + len
	br.bit_cnt_tmp = 0
}

func (br *BitReader) return_bit_reader() int {
	cnt := br.bit_cnt_tmp
	br.bit_cnt, br.bit_len, br.bit_cnt_tmp = br.stack.Top().bit_cnt, br.stack.Top().bit_len, br.stack.Top().bit_cnt_tmp
	br.stack.Pop()
	br.Printf("Return cnt=%d", cnt)
	return cnt
}

func (br *BitReader) read_operator_sick() {
	switch br.read_hex(1) {
	case 0:
		br.Printf("Operator bits")
		br.read_hex(15)
		br.read_packets()
	case 1:
		br.Printf("Operator packets")
		number_of_packets := br.read_hex(11)
		for packet := 0; packet < number_of_packets; packet++ {
			br.read_packets()
		}
	}
}

func (br *BitReader) advance_cnt(cnt int) {
	br.bit_cnt += cnt
	br.bit_cnt_tmp += cnt
}

func (br *BitReader) read_operator(packet_type int) (ret int) {
	number_of_bits := -1
	defer func() {
		if number_of_bits >= 0 {
			br.advance_cnt(br.return_bit_reader())
		}
	}()
	number_of_packets := math.MaxInt
	switch br.read_hex(1) {
	case 0:
		number_of_bits = br.read_hex(15)
		br.Printf("Bit mode b=%d", number_of_bits)
		br.call_bit_reader(number_of_bits)
	case 1:
		number_of_packets = br.read_hex(11)
		br.Printf("Packet Mode p=%d", number_of_packets)
	}

	switch packet_type {
	case 0: // sum
		br.Printf("Operator sum")
		sum := 0
		for !br.nothing_to_read() && number_of_packets > 0 {
			number_of_packets--
			sum += br.read_packets()
		}
		return sum
	case 1:
		br.Printf("Operator product")
		number_of_packets--
		mul := br.read_packets()
		for !br.nothing_to_read() && number_of_packets > 0 {
			number_of_packets--
			mul *= br.read_packets()
		}
		return mul
	case 2:
		br.Printf("Operator minimum")

		min := math.MaxInt
		for !br.nothing_to_read() && number_of_packets > 0 {
			number_of_packets--
			val := br.read_packets()
			if val < min {
				min = val
			}
		}
		return min
	case 3:
		max := math.MinInt
		for !br.nothing_to_read() && number_of_packets > 0 {
			number_of_packets--
			val := br.read_packets()
			if val > max {
				max = val
			}
		}
		return max
	case 5:
		br.Printf("Operator greater than")
		val1 := br.read_packets()
		val2 := br.read_packets()
		if val1 > val2 {
			return 1
		}
		return 0
	case 6:
		br.Printf("Operator less then")
		val1 := br.read_packets()
		val2 := br.read_packets()
		if val1 < val2 {
			return 1
		}
		return 0
	case 7:
		br.Printf("Operator equal to")
		val1 := br.read_packets()
		val2 := br.read_packets()
		if val1 == val2 {
			return 1
		}
		return 0
	default:
		log.Panic("Operator not recognised")
	}
	return ret
}

func (br *BitReader) read_packets() int {
	for !br.nothing_to_read() {
		ver := br.read_version()
		br.Printf("Read Version: %d", ver)
		packet_type := br.read_type()
		switch packet_type {
		case 4:
			val := br.read_value()
			br.Printf("Read Value: %d", val)
			return val
		default:
			return br.read_operator(packet_type)
		}
	}
	return 0
}

func (br *BitReader) start() int {
	return br.read_packets()
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
	var br BitReader
	br.init(data.transmission)
	return br.start()
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d16_1(data))
	log.Printf("02: %d", d16_2(data))
}
