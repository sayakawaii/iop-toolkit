/*
# ------------------------------------------------------------
# -- kafka.go
# --
# -- Huang Minghe
# -- 2022-8-16
# ------------------------------------------------------------
*/

package global

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

var (
	KafkaClient sarama.Client
)

func InitKafka() error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Consumer.Return.Errors = false
	config.Version = sarama.V4_0_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	var err error
	KafkaClient, err = sarama.NewClient([]string{AppConf.Kafka.Addr}, config)
	if err != nil {
		panic(err)
	}
	return nil
}

func KafkaSyncProducer(key string, msg []byte, topic string) error {
	producer, err := sarama.NewSyncProducerFromClient(KafkaClient)
	if err != nil {
		return err
	}
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{Topic: topic, Key: sarama.StringEncoder(key), Value: sarama.ByteEncoder(msg)})
	if err != nil {
		log.Fatalf("unable to produce message: %q", err)
	}
	return nil
}

func KafkaAsyncProducer(key string, msg []byte, topic string) error {
	producer, err := sarama.NewSyncProducerFromClient(KafkaClient)
	if err != nil {
		return err
	}
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{Topic: topic, Key: sarama.StringEncoder(key), Value: sarama.ByteEncoder(msg)})
	if err != nil {
		log.Fatalf("unable to produce message: %q", err)
	}
	return nil
}

type MsgHandler func(*sarama.ConsumerMessage)

type groupHandler struct {
    handler MsgHandler
}

func (groupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (groupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h groupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
    for msg := range claim.Messages() {
        h.handler(msg)
        session.MarkMessage(msg, "")
    }
    return nil
}

func ConsumeTopic(topic, group string, handler MsgHandler) {
    groupClient, err := sarama.NewConsumerGroupFromClient(group, KafkaClient)
    if err != nil {
        panic(err)
    }

    go func() {
        defer groupClient.Close()
        ctx := context.Background()
        for {
            err := groupClient.Consume(ctx, []string{topic}, &groupHandler{
                handler: handler,
            })
            if err != nil {
                log.Println("consume error:", err)
            }
        }
    }()
}
