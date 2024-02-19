FROM golang:1.21.6
LABEL authors="vexsx"


WORKDIR /app


COPY go.* ./


RUN go mod download


COPY . .


RUN go build -o ./app


CMD ["./app"]