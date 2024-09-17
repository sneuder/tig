FROM golang:1.23.1-alpine

WORKDIR /app

RUN apk update && apk upgrade

CMD ["sleep", "infinity"]

