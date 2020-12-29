# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:1.15

ENV GO111MODULE=on

LABEL maintainer="Massimiliano Ceriani <massimilianoceriani08@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

ADD /cmd ./cmd
ADD /internal ./internal
RUN ls

RUN go build -o main ./cmd/main.go

WORKDIR /dist

RUN cp /app/main .

CMD ["/dist/main"]