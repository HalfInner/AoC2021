# AoC2021
Solutions of AoC2021. Written in Go language.

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