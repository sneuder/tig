FROM golang:1.23.1-alpine

WORKDIR /app

RUN apk update && apk upgrade
CMD apk add tar

ENV VERSION="1.0.0"

CMD ["sleep", "infinity"]

