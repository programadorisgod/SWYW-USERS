#STAGE 1
FROM golang:1.25.1-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

RUN go install github.com/air-verse/air@latest
ENV PATH="$PATH:$(go env GOPATH)/bin"

COPY go.mod go.sum ./

RUN go mod download

COPY . .



EXPOSE 5002

RUN apk update && apk add --no-cache ca-certificates && \
    apk add --no-cache wget

CMD ["air", "-c", ".air.toml"]
