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
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func receiveHandler(producer sarama.SyncProducer, serializer Serializer) func(c *gin.Context) {
	return func(c *gin.Context) {
		httpRequestsTotal.Add(float64(1))
		compressed, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			logrus.WithError(err).Error("couldn't read body")
			return
		}
		// 发送信息到kafka
		msg := &sarama.ProducerMessage{Topic: topics, Value: sarama.StringEncoder(compressed)}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			logrus.Printf("FAILED to send message: %s\n", err)
		} else {
			logrus.Printf("> message sent to partition %d at offset %d\n", partition, offset)
		}
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			logrus.WithError(err).Error("couldn't produce message in kafka")
			return
		}

		//
		//reqBuf, err := snappy.Decode(nil, compressed)
		//if err != nil {
		//	c.AbortWithStatus(http.StatusBadRequest)
		//	logrus.WithError(err).Error("couldn't decompress body")
		//	return
		//}
		//
		//var req prompb.WriteRequest
		//if err := proto.Unmarshal(reqBuf, &req); err != nil {
		//	c.AbortWithStatus(http.StatusBadRequest)
		//	logrus.WithError(err).Error("couldn't unmarshal body")
		//	return
		//}

		//metrics, err := processWriteRequest(&req)
		//if err != nil {
		//	c.AbortWithStatus(http.StatusInternalServerError)
		//	logrus.WithError(err).Error("couldn't process write request")
		//	return
		//}

		//for _, metric := range metrics {
		//	err := producer.Produce(&kafka.Message{
		//		TopicPartition: kafkaPartition,
		//		Value:          metric,
		//	}, nil)
		//
		//	if err != nil {
		//		c.AbortWithStatus(http.StatusInternalServerError)
		//		logrus.WithError(err).Error("couldn't produce message in kafka")
		//		return
		//	}
		//}

	}
}
