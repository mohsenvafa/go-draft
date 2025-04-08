// common/kafka.go
package common

import "os"

func GetKafkaBrokers() string {
	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		return "kafka:9092"
	}
	return brokers
}
