package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"sync"
	"time"
	kafkaConfig "linebot/config"
	"linebot/util/log"
)

func newConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V2_2_0_0
	return config
}

func Push(topic string, data interface{}) (*sarama.ProducerMessage, error) {
	var (
		wg sync.WaitGroup
		mesg *sarama.ProducerMessage
	)

	config := newConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	if producer, err := sarama.NewAsyncProducer([]string{kafkaConfig.ConfigKafka.KafkaIp}, config); err != nil {
		return nil, err
	} else if jsonData, err := json.Marshal(data); err != nil {
		return nil, err
	} else {
		defer producer.AsyncClose()
		msg := &sarama.ProducerMessage{
			Topic: topic,
		}
		msg.Value = sarama.ByteEncoder(jsonData)
		msg.Key = sarama.ByteEncoder(topic)
		producer.Input() <- msg
		isSuccess := true
		wg.Add(1)

		func(p sarama.AsyncProducer) {
			defer wg.Done()
			select {
			case mesg = <-p.Successes():

			case fail := <-p.Errors():
				isSuccess = false
				err = fail.Err
			}
		}(producer)

		wg.Wait()

		if !isSuccess {
			return nil, err
		}
	}
	return mesg, nil
}


func Listen(topic string, handler func(value []byte)(isBreak bool)) {
	var (
		pc sarama.PartitionConsumer
		value []byte
	)

	config := newConfig()
	config.Consumer.Offsets.CommitInterval = 1 * time.Second

	if consumer, err := sarama.NewConsumer([]string{kafkaConfig.ConfigKafka.KafkaIp}, config); err != nil {
		log.Error(err)
	} else if partitionList, err := consumer.Partitions(topic); err != nil {
		log.Error(err)
	} else {
		for partition := range partitionList {
			if pc, err = consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest); err != nil {
				log.Error(err)
			} else {
				for {
					select {
					case msg := <-pc.Messages():
						value = msg.Value
						break
					}
					if handler(value) {
						break
					}
				}
			}
			err = consumer.Close()
		}
		defer closePc(pc)
	}
}

func closePc(pc sarama.PartitionConsumer)  {
	if pc != nil {
		pc.AsyncClose()
	}
}