# Base go image
FROM golang:1.22-alpine as gobuilder

RUN mkdir /app

RUN mkdir /app/logs && chmod -R 0777 /app/logs

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o broker ./cmd/api

RUN chmod +x /app/broker

# build image
FROM alpine:latest

RUN mkdir /app 

# RUN mkdir /app/logs && chmod -R 0777 /app/logs

COPY --from=gobuilder . /app

# COPY --from=gobuilder /app/logs /app/logs

CMD ["/app/broker"]



