# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:1.15.6

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

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