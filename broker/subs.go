package broker

import (
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
)

func NewSubscriber(clientID string, subj string, conf Config, hdl stan.MsgHandler) (stan.Subscription, error) {
	if conn == nil {
		return nil, fmt.Errorf("nil conn")
	}

	if conn.IsClosed() {
		return nil, fmt.Errorf("closed conn")
	}

	sc, err := stan.Connect(
		conf.ClusterID,
		clientID,
		stan.NatsConn(conn),
		stan.NatsURL(conf.URL),
		stan.Pings(1, 3),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))

	if err != nil {
		return nil, err
	}

	sub, err := sc.Subscribe(subj, hdl)
	if err != nil {
		return nil, err
	}

	return sub, nil
}
