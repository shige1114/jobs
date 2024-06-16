FROM golang:1.22-alpine
RUN apk update && \
    apk add git openssh-client socat && \
    apk add --upgrade grep
WORKDIR /app/02_company

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
