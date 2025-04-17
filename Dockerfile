FROM golang:1.19 AS build-env


ENV GOOS=linux
ENV GO_PRIVATE=github.com/runetale/client-go

WORKDIR /go/src/github.com/runetale/runetale-handshake-server

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /runetale-handshake-server .

FROM debian:bullseye

COPY --from=build-env /runetale-handshake-server /runetale-handshake-server

RUN chmod u+x /runetale-handshake-server

CMD ["/runetale-handshake-server"]
