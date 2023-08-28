package broker

import (
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
)

type Publisher struct {
	conn stan.Conn
	subj string
}

func NewPublisher(clientID string, subj string, conf Config) (*Publisher, error) {
	if conn == nil {
		return nil, fmt.Errorf("not conn")
	}

	if !conn.IsConnected() {
		return nil, fmt.Errorf("not conn")
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

	return &Publisher{
		conn: sc,
		subj: subj,
	}, nil
}

func (p *Publisher) Publish(data []byte) error {
	return p.conn.Publish(p.subj, data)
}
