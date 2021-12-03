FROM golang:1.15-alpine

WORKDIR /app

COPY . .

# Of course faster is to have one command with following "&&",
# but I prefer to have fast re-build of the last puzzle
RUN go install d01/d01.go
RUN go install d02/d02.go
RUN go install d03/d03.go

CMD ["./start.sh"]
