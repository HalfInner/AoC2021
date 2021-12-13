package main

import (
	"AoC2021/aoc_fun"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type Point struct {
	x, y int
}

type Fold struct {
	axis rune
	dist int
}

type Data struct {
	points []Point
	folds  []Fold
}

func parse_point(line string, data *Data) {
	var x, y int
	fmt.Sscanf(line, "%d,%d", &x, &y)
	data.points = append(data.points, Point{x, y})
}

func parse_fold(line string, data *Data) {
	var axis rune
	var dist int
	var dumb string
	fmt.Sscanf(line, "%s %s %c=%d", &dumb, &dumb, &axis, &dist)
	data.folds = append(data.folds, Fold{axis, dist})
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
		parse_point(line, &data)
	}
	for _, line := range strings.Split(input_split[1], "\n") {
		parse_fold(line, &data)
	}

	return data
}

func fold_paper(fold Fold, paper map[Point]int) {
	for p, _ := range paper {
		tmp_p := p
		check := false
		switch fold.axis {
		case 'x':
			if p.x > fold.dist {
				local_dist := p.x - fold.dist
				p.x = p.x - 2*local_dist
				check = true
			}
		case 'y':
			if p.y > fold.dist {
				local_dist := p.y - fold.dist
				p.y = p.y - 2*local_dist
				check = true
			}

		default:
			log.Panic("No axis")
		}

		if check {
			if _, ok := paper[tmp_p]; ok {
				delete(paper, tmp_p)
			}
		}
		paper[p]++
	}
}

func d13_1(data Data) int {
	defer aoc_fun.Track(aoc_fun.Runningtime())

	paper := make(map[Point]int)
	for _, p := range data.points {
		paper[p]++
	}
	fold := data.folds[0]
	fold_paper(fold, paper)

	return len(paper)
}

func d13_2(data Data) string {
	defer aoc_fun.Track(aoc_fun.Runningtime())
	paper := make(map[Point]int)

	for _, p := range data.points {
		paper[p]++
	}

	for _, fold := range data.folds {
		fold_paper(fold, paper)
	}

	var short_points []Point
	for key, _ := range paper {
		short_points = append(short_points, key)
	}

	sort.Slice(short_points, func(i, j int) bool {
		if short_points[i].y == short_points[j].y {
			return short_points[i].x < short_points[j].x
		}
		return short_points[i].y < short_points[j].y
	})

	var result strings.Builder
	last := Point{0, 0}
	for idx := 0; idx < len(short_points); idx++ {
		p := short_points[idx]
		for jdx := last.y; jdx < p.y; jdx++ {
			result.WriteRune('\n')
		}

		for kdx := last.x; kdx < p.x-1; kdx++ {
			result.WriteRune('.')
		}
		result.WriteRune('#')
		last = p
	}
	result.WriteRune('\n')

	return result.String()
}

func main() {
	defer aoc_fun.Unprofile(aoc_fun.ProfileCPU())

	data := read_data()
	log.Printf("01: %d", d13_1(data))
	log.Printf("02:\n%s", d13_2(data))
}
