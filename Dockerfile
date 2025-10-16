#STAGE 1
FROM golang:1.25.1-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

RUN go install github.com/air-verse/air@latest
ENV PATH="$PATH:$(go env GOPATH)/bin"

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# BUILD MAIN BINARY
RUN go build -o user-service main.go


#STAGE 2

FROM alpine:3.19

RUN apk update && apk add --no-cache ca-certificates && \
    apk add --no-cache wget \
    && addgroup -g 1000 usergo && adduser -u 1000 -G usergo -s /bin/sh -D usergo

WORKDIR /app/

USER usergo

COPY --from=builder /app/user-service .

EXPOSE 5002

CMD [ "./user-service" ]
