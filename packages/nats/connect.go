package nats

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

type Client struct {
	conn *nats.Conn
	js   nats.JetStreamContext
}

type Config struct {
	URL           string
	Name          string
	MaxReconnects int
	ReconnectWait time.Duration
	Timeout       time.Duration
}

func New(cfg Config) (*Client, error) {

	opts := []nats.Option{
		nats.Name(cfg.Name),
		nats.MaxReconnects(cfg.MaxReconnects),
		nats.ReconnectWait(cfg.ReconnectWait),
		nats.Timeout(cfg.Timeout),
		nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
			fmt.Printf("[NATS] disconnected: %v\n", err)
		}),
		nats.ReconnectHandler(func(_ *nats.Conn) {
			fmt.Println("[NATS] reconnected")
		}),
	}

	conn, err := nats.Connect(cfg.URL, opts...)
	if err != nil {
		return nil, err
	}

	js, err := conn.JetStream()
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
		js:   js,
	}, nil
}
