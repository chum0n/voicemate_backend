FROM golang:1.16-alpine3.13

WORKDIR /go/src

# Install Go modules.
COPY ./voicemate_api/go.mod ./
RUN go mod download
