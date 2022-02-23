FROM golang:1.16.5-alpine3.13 AS BUILD

ENV ENV 'prod'

RUN apk upgrade --update && \
    apk --no-cache add git

WORKDIR /app
COPY . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/sls

FROM alpine:3.13
COPY --from=BUILD /bin/sls /bin/
CMD ["/bin/sls"]
