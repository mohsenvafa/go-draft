// producer/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	messagesProduced = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "messages_produced_total",
		Help: "Total number of messages produced",
	})
)

func init() {
	prometheus.MustRegister(messagesProduced)
}

func main() {
	log.Println("Producer Started....")
	brokers := []string{"kafka:9092"}
	var producer sarama.SyncProducer
	var err error

	for retries := 0; retries < 20; retries++ {
		producer, err = sarama.NewSyncProducer(brokers, nil)
		if err == nil {
			log.Println("✅ Connected to Kafka as producer")
			break
		}
		log.Printf("❌ Kafka not ready yet for producer (attempt %d): %v\n", retries+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("❌ Failed to connect to Kafka after retries: %v", err)
	}

	log.Println("Producer to kafa is passed")

	defer producer.Close()

	go func() {
		for {

			msg := &sarama.ProducerMessage{
				Topic: "test-topic",
				Value: sarama.StringEncoder(fmt.Sprintf("Hello Kafka at %v", time.Now())),
			}
			_, _, err := producer.SendMessage(msg)
			if err == nil {
				messagesProduced.Inc()
				log.Println("Message produced")
			}

			if err != nil {
				log.Println("Failed to send message to Kafka producer: %v", err)
			}

			time.Sleep(5 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
