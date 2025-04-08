// consumer/main.go
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	messagesConsumed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "messages_consumed_total",
		Help: "Total number of messages consumed",
	})
)

func init() {
	prometheus.MustRegister(messagesConsumed)
}

type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Consumed message: %s\n", string(msg.Value))
		messagesConsumed.Inc()
		session.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	log.Println("Consumer Started....")

	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	brokers := []string{"kafka:9092"}
	var group sarama.ConsumerGroup
	var err error

	// Retry loop: wait for Kafka to become available
	for retries := 0; retries < 10; retries++ {
		group, err = sarama.NewConsumerGroup(brokers, "my-group", config)
		if err == nil {
			break
		}
		log.Printf("Kafka not ready yet (attempt %d): %v\n", retries+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to Kafka after retries: %v", err)
	}

	defer group.Close()

	go func() {
		for {
			err := group.Consume(context.Background(), []string{"test-topic"}, ConsumerGroupHandler{})
			if err != nil {
				log.Println("Error consuming:", err)
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8081", nil))
}
