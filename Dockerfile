FROM golang:1.25

RUN apt-get update && apt-get install -y gcc libc6-dev

ENV CGO_ENABLED=1

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

CMD ["./main"]