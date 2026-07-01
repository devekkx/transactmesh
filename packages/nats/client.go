package nats

import (
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
