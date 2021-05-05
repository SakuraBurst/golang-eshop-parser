FROM golang:1.16

WORKDIR /go/parser

COPY . .

RUN go get

RUN go install

RUN go build

CMD ["/go/parser/eshop-parser"]