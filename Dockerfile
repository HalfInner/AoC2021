FROM golang:1.17-alpine

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

COPY d04/ d04/
RUN go install d04/d04.go

COPY d05/ d05/
RUN go install d05/d05.go

COPY d06/ d06/
RUN go install d06/d06.go

COPY d07/ d07/
RUN go install d07/d07.go

COPY d08/ d08/
RUN go install d08/d08.go

COPY d09/ d09/
RUN go install d09/d09.go

COPY d10/ d10/
RUN go install d10/d10.go

COPY d11/ d11/
RUN go install d11/d11.go

COPY d12/ d12/
RUN go install d12/d12.go

COPY d13/ d13/
RUN go install d13/d13.go

CMD ["./start.sh"]
