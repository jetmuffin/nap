FROM golang:1.8.1-alpine

COPY . /workspace
WORKDIR /workspace

RUN make clean && make

ENTRYPOINT ["./build/nap"]
