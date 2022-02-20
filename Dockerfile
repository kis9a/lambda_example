FROM golang:1.16.5-alpine3.13 AS BUILD

ENV ENDPOINT 'lambda'
ENV AWS_REGION 'ap-northeast-1'
ENV AWS_S3_BUCKET 'kis9a-lambda-sls'

WORKDIR /app

RUN apk add gcc build-base

ADD /go.mod /app/
ADD /go.sum /app/

ADD /main.go /app/
ADD /handlers /app/handlers

RUN go mod download

RUN apk upgrade --update && \
    apk --no-cache add git

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/sls

FROM alpine:3.13
COPY --from=BUILD /bin/sls /bin/
CMD ["/bin/sls"]
