package ndn

import (
	"net/http"

	"github.com/go-chi/chi"
)

type ClientConfig struct{}

type Client struct{}

func (Client) Start() error {
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
