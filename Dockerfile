FROM golang:1.16.2-alpine as builder


RUN apk update && apk add --no-cache make gcc libc-dev protobuf-dev
RUN go get github.com/golang/protobuf/protoc-gen-go

ENV GO11MODULE=on

ENV PORT=18080
ENV TOPIC=kafka_accident
ENV TOPIC_ACCIDENT=kafka_accident
ENV TOPIC_DROWSINESS=kafka_dds
ENV BROKER_1_ADDRESS=localhost:9092
ENV BROKER_2_ADDRESS=localhost:9093
ENV BROKER_3_ADDRESS=localhost:9094
ENV GROUP_ID=my-group
ENV DATA_MANAGEMENT_CONNECTION=127.0.0.1:8082
ENV DATABASE_URI=mongodb+srv://admin:x2vg5rs@cluster0.8ulm4.mongodb.net
ENV DATABASE_URI=5gv2x
ENV DATABASE_URI=
ENV DATABASE_URI=

EXPOSE 18080

