FROM golang:1.10 AS build
WORKDIR /go/src

ENV CGO_ENABLED=0
RUN go get -d -v ./...

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD








RUN go install -v ./...
COPY --from=build /go/bin/ /go/bin/
=======
>>>>>>> 65dbe9c (first push on new branch)
=======
RUN go build -a -installsuffix cgo -o swagger .
ZIHBDIHZBIZBBIZDIZ
dDZ
DZZ
ZIHBDIHZBIZBBIZDIZZD
ZD

ZD
ZIHBDIHZBIZBBIZDIZZD
ZD

ZD
ZD
Z
D
ZIHBDIHZBIZBBIZDIZDZ
ZD
Z
FROM scratch AS runtime
COPY --from=build /go/src/swagger ./
EXPOSE 8080/tcp
ENTRYPOINT ["./swagger"]
>>>>>>> 55d0c3d (test)
=======
>>>>>>> 65dbe9c (first push on new branch)
