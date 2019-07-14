# Creates webhookd daemon image

FROM golang:1.11.2-alpine
ADD . /go/src/whosonfirst/go-webhookd
WORKDIR /go/src/whosonfirst/go-webhookd
RUN apk add --no-cache git make
RUN make bin

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/whosonfirst/go-webhookd/bin/webhookd /webhookd
EXPOSE 8080
ENTRYPOINT ["/webhookd"]
