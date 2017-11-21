FROM golang:1.8

WORKDIR /go/src/ifconfig-service
COPY . .

WORKDIR main

RUN go-wrapper download 
RUN go-wrapper install

#EXPOSE 55555

CMD ["go-wrapper", "run"]  ["main.go"]

