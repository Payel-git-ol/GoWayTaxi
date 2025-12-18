package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	KafkaMessagesOut = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "kafka_messages_out_total",
		Help: "Total number of Kafka messages produced",
	})
	KafkaMessagesIn = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "kafka_messages_in_total",
		Help: "Total number of Kafka messages consumed",
	})
)

func Init() {
	prometheus.MustRegister(KafkaMessagesOut, KafkaMessagesIn)
}
