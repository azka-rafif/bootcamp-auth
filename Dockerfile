# Builder
FROM golang:1.14.2-alpine3.11 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

RUN go get -u github.com/swaggo/swag/cmd/swag

COPY . .

RUN swag init

RUN make engine

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && apk --no-cache add ca-certificates && \
    apk --update --no-cache add tzdata

ENV TZ=Asia/Jakarta

WORKDIR /app 

EXPOSE 9090

COPY --from=builder /app/engine /app

# enable if running kube
RUN ln -s /env/.env .

CMD /app/engine