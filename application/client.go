package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"ndn/repository/postgres"
)

type ClientConfig struct {
	Storage ClientConfigStorage
}

type ClientConfigStorage struct {
	Postgres postgres.ClientConfig
}

type Client struct{}

func (Client) Start() error {
	pgcfg := postgres.ClientConfig{
		Host:           "localhost",
		Port:           5432,
		Database:       "ndn",
		User:           "postgres",
		Password:       "postgres",
		MaxConnections: 100,
		AcquireTimeout: time.Second,
	}
	c, err := postgres.NewClient(pgcfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Pool.Stat())

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
	return nil
}

func NewClient(cfg ClientConfig) (Client, error) {
	return Client{}, nil
}
