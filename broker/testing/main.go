package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func main() {

	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)

	}

	sconn, err := stan.Connect("test-cluster", "clientID", stan.NatsConn(conn))
	if err != nil {
		log.Fatal(err)
	}

	defer sconn.Close()

	sub, err := sconn.Subscribe("topec", func(msg *stan.Msg) {
		fmt.Println(string(msg.Data))
	})
	defer sub.Unsubscribe()

	var i int
	for {

		err := sconn.Publish("topec", []byte(fmt.Sprintf("data%d", i)))
		if err != nil {
			log.Println(err)
			break
		}
		i++
		time.Sleep(1 * time.Second)
	}
}
