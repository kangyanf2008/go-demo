package kafka

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"strings"
	"sync/atomic"
	"time"
)

type KafkaConsumer struct {
	config  *sarama.Config
	addrs   []string
	group   sarama.ConsumerGroup
	groupId string
	done    int32
}

func NewKafkaConsumer(addr string, groupId string, config *sarama.Config) (*KafkaConsumer, error) {
	addrs := strings.Split(addr, ",")
	consume := &KafkaConsumer{
		config:  config,
		addrs:   addrs,
		groupId: groupId,
	}
	group, err := sarama.NewConsumerGroup(addrs, groupId, config)
	if err != nil {
		log.Printf("初始化消费者组失败 -> %v %s %s", err, addr, groupId)
		return nil, err
	}
	consume.group = group

	return consume, nil
}

func (c *KafkaConsumer) Run(ctx context.Context, topics []string) error {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				println("Kafka 消费者出现 panic ->", "error", err)
			}
		}()
		consume := &consumer{}
		for {
			if err := c.group.Consume(ctx, topics, consume); err != nil {
				println("创建消费者失败 ->", "topics:", strings.Join(topics, ","), "error:", err)
			}
		}
	}()
	log.Println("kafka 消费者服务已启动...")
	return nil
}

func (c *KafkaConsumer) Close() error {
	if atomic.LoadInt32(&c.done) == 1 {
		return nil
	}
	if c.done == 0 {
		atomic.StoreInt32(&c.done, 1)
		if c.group != nil {
			_ = c.group.Close()
		}
	}
	return nil
}
func CreateKafkaConfig(version, user, password string, certBytes []byte) (*sarama.Config, error) {
	config := sarama.NewConfig()

	if ver, err := sarama.ParseKafkaVersion(version); err == nil {
		config.Version = ver
	} else {
		config.Version = sarama.V0_10_2_0
	}
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Timeout = time.Second * 5
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Return.Errors = true

	if user != "" && password != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = user
		config.Net.SASL.Password = password
		config.Net.SASL.Handshake = true
	}
	if certBytes != nil && len(certBytes) > 0 {
		clientCertPool := x509.NewCertPool()
		ok := clientCertPool.AppendCertsFromPEM(certBytes)
		if !ok {
			return nil, errors.New("kafka producer failed to parse root certificate")
		}

		config.Net.TLS.Config = &tls.Config{
			//Certificates:       []tls.Certificate{},
			RootCAs:            clientCertPool,
			InsecureSkipVerify: true,
		}

		config.Net.TLS.Enable = true
		config.Producer.Return.Successes = true
	}

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("kafka producer config invalidate. config: %v. err: %v", *config, err)
	}
	return config, nil
}
