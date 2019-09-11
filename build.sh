#!/usr/bin/env bash

make build
docker build -t promsaramaadapter:1.0.0 .
# docker run -d --net=host -p 8080:8080 --name promsaramaadapter promsaramaadapter:1.0.0 --brokers=10.21.6.148:9092





