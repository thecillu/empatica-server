FROM golang:latest
RUN mkdir -p $GOPATH/src/empatica-server
ADD . $GOPATH/src/empatica-server/
WORKDIR $GOPATH/src/empatica-server
RUN go get ./...
RUN go build -o main .
EXPOSE 80
CMD ["./main"]