package main

import (
	"bytes"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io/ioutil"
	"net/http"
)

func CreateKafkaConsumer() *kafka.Consumer {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	return consumer
}

func SubscribeTopic(consumer *kafka.Consumer)  {
	consumer.Subscribe("postgres.public.Product",nil)

	 fmt.Println("Subscribed to product topic")
}

func ReadTopicMessages(consumer *kafka.Consumer) string {

	var message string
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			//message = message + string(msg.Value)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	CloseConsumer(consumer)

	return message
}

func CloseConsumer(consumer *kafka.Consumer){
	consumer.Close()
}

func RegisterConnector() *http.Response {
	plan, _ := ioutil.ReadFile("./connectors/debezium-connector.json")
	response, err := http.Post("http://localhost:8083/connectors/","application/json",bytes.NewBuffer(plan))

	if err != nil{
		panic(err)
	}

	return response
}

func CheckConnector() {
	response, err := http.Get("http://localhost:8083/connectors/product_connector")
	defer response.Body.Close()

	if err != nil{
		panic(err)
	}
	if response.StatusCode != 200 {
		RegisterConnector()
	}

	// If you want to show response, uncomment following lines

	//body, _ := ioutil.ReadAll(response.Body)

	//fmt.Println(string(body))
}

func main()  {
	consumer := CreateKafkaConsumer()
	CheckConnector()

	SubscribeTopic(consumer)
	ReadTopicMessages(consumer)

	fmt.Scan()
}