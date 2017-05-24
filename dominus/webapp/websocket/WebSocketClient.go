package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
	"github.com/gorilla/websocket"
	"github.com/Shopify/sarama"
	"strconv"
)

var addr = flag.String("addr", "45.63.51.192:9090", "http service address")

const (
	PRODUCER_URL string = "10.139.50.148:9092"
	KAFKA_TOPIC string = "KP-SPIDER-170523"
)

//read remote ws stream from websocketd and write it to kafka.
func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	producer := kafkaProducer()
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			if len(message) == 0 {
				continue
			}
			log.Printf("recv: %s", message)

			//EE: send to kafka
			strTime := strconv.Itoa(int(time.Now().Unix()))
			msg := &sarama.ProducerMessage{
				Topic: KAFKA_TOPIC,
				Key:   sarama.StringEncoder(strTime),
				Value: sarama.StringEncoder(message),
			}
			producer.Input() <- msg
		}
	}()

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		//read from producer errors channel in case of deadlock.
		case errors := <-producer.Errors():
		//MUST read from this channel or the Producer will deadlock when the channel is full
		//panic: kafka: Failed to deliver 1 messages.
			println(errors)
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
		// To cleanly close a connection, a client should send a close
		// frame and wait for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
				select {
				case <-done:
				case <-time.After(time.Second):
				}
			c.Close()
			return
		}
	}
}

func kafkaProducer() sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	brokers := []string{PRODUCER_URL}
	producer, err := sarama.NewAsyncProducer(brokers, config); if (err != nil) {
		panic(err)
	}
	return producer
}