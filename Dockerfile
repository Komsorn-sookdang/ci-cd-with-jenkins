FROM golang:1.14

ENV GO111MODULE=on
WORKDIR /c2fit-hw1

COPY . .

RUN go mod download
RUN go mod vendor

RUN go build -o main ./cmd/web
CMD ["./main"]