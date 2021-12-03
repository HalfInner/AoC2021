package aoc_fun

import (
	"log"
	"time"
)

func Runningtime() time.Time {
	return time.Now()
}

func Track(startTime time.Time) {
	endTime := time.Now()
	log.Println("Took", endTime.Sub(startTime))
}
