FROM golang:1.10.3
EXPOSE 8080

COPY . /go/src/github.com/bwynn/goalmanac
WORKDIR /go/src/github.com/bwynn/goalmanac

RUN go get -d -v
RUN go install -v

RUN cd $GOPATH/src/github.com/bwynn/goalmanac

CMD ["goalmanac"]

