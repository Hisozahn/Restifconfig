FROM golang:1.8

WORKDIR /go/src/ifconfig-cli
COPY . .

WORKDIR main

RUN go-wrapper download 
RUN go-wrapper install

ENTRYPOINT ["go-wrapper", "run"] ["main.go"]
