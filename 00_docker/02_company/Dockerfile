FROM golang:1.22-alpine
RUN apk update && \
    apk add git && \
    apk add --upgrade grep
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
