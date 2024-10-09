# build image
FROM alpine:latest

RUN mkdir /app 

RUN mkdir /app/logs && chmod -R 0777 /app/logs

COPY brokerApp /app

RUN ls -lrt

CMD ["/app/brokerApp"]



