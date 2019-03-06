package postgres

import (
	"time"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"
)

// ClientConfig contains all the options used to by client struct.
type ClientConfig struct {
	Host           string
	Port           uint16
	Database       string
	User           string
	Password       string
	MaxConnections int

	// Max wait time when all connections are busy (0 means no timeout).
	AcquireTimeout time.Duration
}

// Client is used to connect to postgres.
type Client struct {
	Pool *pgx.ConnPool
}

func (c *Client) init(cfg ClientConfig) error {
	pgcfg := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     cfg.Host,
			Port:     cfg.Port,
			Database: cfg.Database,
			User:     cfg.User,
			Password: cfg.Password,
		},
		AcquireTimeout: cfg.AcquireTimeout,
		MaxConnections: cfg.MaxConnections,
	}
	cpool, err := pgx.NewConnPool(pgcfg)
	if err != nil {
		return errors.Wrap(err, "connection pool initialization error")
	}
	c.Pool = cpool
	return nil
}

// Close the connection pool.
func (c *Client) Close() {
	c.Pool.Close()
}

// NewClient return a initialized client struct.
func NewClient(cfg ClientConfig) (*Client, error) {
	var c Client
	if err := c.init(cfg); err != nil {
		return nil, errors.Wrap(err, "initialization error")
	}
	return &c, nil
}
