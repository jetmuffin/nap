FROM golang:1.8.1-alpine

RUN apk --no-cache add make git gcc

WORKDIR /go/src/github.com/JetMuffin/nap
COPY . .

RUN make clean && make

ENTRYPOINT ["./bin/nap", "master"]
