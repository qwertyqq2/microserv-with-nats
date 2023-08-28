package broker

import (
	"github.com/nats-io/nats.go"
)

type Config struct {
	ClusterID string
	URL       string
	Subject   string
}

var conn *nats.Conn

func Connect(url string) error {
	newConn, err := nats.Connect(url)
	if err != nil {
		return err
	}

	conn = newConn

	return nil
}

func Close() { conn.Close() }
