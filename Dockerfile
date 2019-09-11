#FROM golang:1.11-alpine as build
#
#WORKDIR /src/prometheus-sarama-adapter
#ADD . /src/prometheus-sarama-adapter
#ENV GO111MODULE=on
#ENV GOPROXY=https://goproxy.io
#COPY go.mod .
#COPY go.sum .
#RUN go mod download
#COPY . .
#RUN make build
#FROM alpine:3.8
#MAINTAINER mqiqe@163.com
#COPY --from=build /prometheus-sarama-adapter/promsaramaadapter /
#
#ENTRYPOINT ["/promsaramaadapter"]

FROM alpine:3.8
MAINTAINER mqiqe@163.com
COPY ./promsaramaadapter /
ENTRYPOINT ["/promsaramaadapter"]
