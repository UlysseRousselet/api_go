FROM golang:1.10 AS build
WORKDIR /go/src

ENV CGO_ENABLED=0
RUN go get -d -v ./...









RUN go install -v ./...
COPY --from=build /go/bin/ /go/bin/
