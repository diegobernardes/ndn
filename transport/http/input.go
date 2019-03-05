package http

import "net/http"

type Input struct{}

func (Input) Index(w http.ResponseWriter, r *http.Request) {}

func (Input) Create(w http.ResponseWriter, r *http.Request) {}

func (Input) Delete(w http.ResponseWriter, r *http.Request) {}

func NewInput() Input {
	return Input{}
}
