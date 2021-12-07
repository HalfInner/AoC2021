package aoc_fun

import (
	"log"
	"os"
	"path"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
)

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Runningtime() time.Time {
	return time.Now()
}

func Track(startTime time.Time) {
	endTime := time.Now()
	log.Println("Took", endTime.Sub(startTime))
}

func GetDefaultInputFilePath() string {
	_, filename, _, _ := runtime.Caller(2)
	filename = string(path.Dir(filename)) + "/input.txt"
	return filename
}

func ProfileCPU() *os.File {
	// Example run:
	// ```ENABLE_PROFILING=TRUE go run d04/d04.go && go tool pprof -ignore 'syscall' -ignore 'aoc_fun' -dot cpu.prof | dot -Tpng  -o call_profile_graph.png```
	do_profile := strings.ToUpper(os.Getenv("ENABLE_PROFILING")) == "TRUE"
	if !do_profile {
		return nil
	}
	log.Println("*** PROFILING ENABLED ***")
	cpu_profile_file_handler, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}

	if err := pprof.StartCPUProfile(cpu_profile_file_handler); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}

	return cpu_profile_file_handler
}

func Unprofile(cpu_profile_file *os.File) {
	if cpu_profile_file != nil {
		pprof.StopCPUProfile()
		cpu_profile_file.Close()
	}
}
