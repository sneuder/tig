FROM golang:1.26.5-alpine

WORKDIR /app

RUN apk update && apk upgrade
CMD apk add tar && apk add git

ENV VERSION="3.0.0"

CMD ["sleep", "infinity"]

