FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d github.com/gorilla/mux
RUN go install

CMD ["main"]