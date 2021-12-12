# AoC2021
Solutions of the Advent of Code 2021. Written in a Go language. Measuring an execution time.

```txt
    *** Advent of Code 2021 ***
===================================
d01:
2021/12/10 18:09:58 Took 3.9µs
2021/12/10 18:09:58 01: 1393
2021/12/10 18:09:58 Took 8µs
2021/12/10 18:09:58 02: 1359
```

## Run
```sh
go run d01/d01.go d01/input.txt
# or
go run d01/d01.go
```

## Docker
```sh
docker build -t aoc_2021:latest .
docker container run aoc_2021:latest
```
## Profiling
Profiling is available for each solution, however env 'ENABLE_PROFILING' has to be set to true.

```sh
ENABLE_PROFILING=TRUE go run d04/d04.go && go tool pprof -ignore 'syscall' -ignore 'aoc_fun' -dot cpu.prof | dot -Tpng  -o call_profile_graph.png
```
_it requires 'dot' program installed_

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Author
**Half's Inner**