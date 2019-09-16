// Copyright 2018 Telefónica
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/Shopify/sarama"
	"github.com/containous/traefik/log"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func main() {
	log.Info("creating kafka producer")

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	client, err := sarama.NewClient(strings.Split(brokers, ","), config)
	if err != nil {
		log.Fatalf("unable to create kafka client: %q", err)
		panic(err)
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		log.Fatalf("unable to create kafka producer: %q", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err != nil {
		logrus.WithError(err).Fatal("couldn't create kafka producer")
	}

	r := gin.New()

	r.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true), gin.Recovery())

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	if basicAuth {
		authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
			basicAuthUsername: basicAuthPassword,
		}))
		authorized.POST("/receive", receiveHandler(producer))
	} else {
		r.POST("/receive", receiveHandler(producer))
	}
	err = r.Run()
	logrus.Error(err)
}
