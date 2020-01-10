FROM golang:1.12.15-alpine

WORKDIR /opt/MQTT-client
COPY . .

RUN go build -mod=vendor

ENTRYPOINT ["./MQTT-client"]
