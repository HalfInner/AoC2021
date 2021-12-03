FROM golang:1.15-alpine

WORKDIR /app

COPY aoc_fun/ aoc_fun/
COPY go.mod go.mod
COPY start.sh start.sh

COPY d01/ d01/
RUN go install d01/d01.go

COPY d02/ d02/
RUN go install d02/d02.go

COPY d03/ d03/
RUN go install d03/d03.go

CMD ["./start.sh"]
