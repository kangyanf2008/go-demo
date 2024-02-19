package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"io/ioutil"
	"path/filepath"
	"time"
)

type KafkaStorage struct {
	addrs    []string
	config   *sarama.Config
	topic    string
	producer sarama.SyncProducer
}

func NewKafkaStorage(addrs []string, topic, user, pass, cert string) (*KafkaStorage, error) {
	config := sarama.NewConfig()

	config.Version = sarama.V0_10_2_1
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Timeout = time.Second * 5

	//如果启用了SASL认证
	if user != "" && pass != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = user
		config.Net.SASL.Password = pass
		config.Net.SASL.Handshake = true
	}
	//如果启用了SSL加密
	if cert != "" {
		certPath, err := filepath.Abs(cert)
		if err != nil {
			return nil, err
		}
		certBytes, err := ioutil.ReadFile(certPath)
		clientCertPool := x509.NewCertPool()
		ok := clientCertPool.AppendCertsFromPEM(certBytes)
		if !ok {
			return nil, errors.New("kafka producer failed to parse root certificate")
		}

		config.Net.TLS.Config = &tls.Config{
			RootCAs:            clientCertPool,
			InsecureSkipVerify: true,
		}

		config.Net.TLS.Enable = true
		config.Producer.Return.Successes = true
	}
	//校验配置是否合法
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("kafka producer config invalidate. config: %v. err: %v", *config, err)
	}

	client, err := sarama.NewSyncProducer(addrs, config)

	if err != nil {
		return nil, err
	}
	kafkaClient := &KafkaStorage{
		addrs:    addrs,
		config:   config,
		producer: client,
		topic:    topic,
	}

	return kafkaClient, nil
}

func (c *KafkaStorage) Write(content []byte) error {

	message := &sarama.ProducerMessage{
		Topic: c.topic,
	}
	message.Value = sarama.ByteEncoder(content)

	//使用通道发送
	_, _, err := c.producer.SendMessage(message)
	return err

}

func (c *KafkaStorage) createTopic(topic string) error {
	if len(c.addrs) <= 0 {
		return errors.New("kafka 地址错误")
	}
	broker := sarama.NewBroker(c.addrs[0])
	err := broker.Open(c.config)
	if err != nil {
		println("初始化客户端失败 -> ", "error:", err, "Topic:", topic)
		return err
	}
	req := &sarama.CreateTopicsRequest{
		TopicDetails: map[string]*sarama.TopicDetail{
			"blah25s": {
				NumPartitions:     10,
				ReplicationFactor: 1,
				ConfigEntries:     make(map[string]*string),
			},
		},
		Timeout: time.Second * 15,
	}
	if _, err := broker.CreateTopics(req); err != nil {
		println("创建 Topic 失败 -> ", "error:", err, "Topic:", topic)
		return err
	}
	return nil
}

func (c *KafkaStorage) Close() error {
	if c.producer != nil {
		return c.producer.Close()
	}
	return nil
}
