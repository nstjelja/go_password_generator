FROM golang
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go test
RUN go install -v ./...
CMD ["app"]