FROM golang:1.21.4-alpine3.17

WORKDIR /app

#RUN go mod tidy
RUN apk add --no-cache gcc
RUN apk add --no-cache musl-dev

#CMD ["go", "run", "exsample.go"]
