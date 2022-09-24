package broker

import (
	"context"
	"encoding/json"
	"github.com/rafimuhammad01/learn-go-graphql/internal/core"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Kafka struct {
	conn   *kafka.Conn
	reader *kafka.Reader
}

func (k *Kafka) Consume(ctx context.Context) (*core.Message, *core.Error) {
	for {
		var msgParsed core.Message

		m, err := k.reader.ReadMessage(context.Background())
		if err != nil {
			return nil, core.NewError("unable to read message from kafka", "internal server error", err)
		}
		logrus.Infof("[kafka] message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		err = json.Unmarshal(m.Value, &msgParsed)
		if err != nil {
			return nil, core.NewError("unable to parse message to json", "internal server error", err)
		}

		return &msgParsed, nil
	}
}

func (k *Kafka) Produce(ctx context.Context, msg interface{}) *core.Error {
	data, err := json.Marshal(msg)
	if err != nil {
		return core.NewError("failed to parse msg to byte", "internal server error", nil)
	}

	_, err = k.conn.WriteMessages(
		kafka.Message{Value: data},
	)
	if err != nil {
		return core.NewError("failed to write message", "internal server error", err)
	}

	return nil
}

func NewKafka(address, topic string, partition ...int) *Kafka {
	var fixedPartition int
	if len(partition) == 0 {
		fixedPartition = 0
	} else {
		fixedPartition = partition[0]
	}

	// producer
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, fixedPartition)
	if err != nil {
		logrus.Fatal("failed to dial leader:", err)
	}

	// consumer
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{address},
		Topic:     topic,
		Partition: fixedPartition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	return &Kafka{conn: conn, reader: r}
}
