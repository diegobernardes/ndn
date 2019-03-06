package http

import "net/http"

type Input struct{}

func (Input) Index(w http.ResponseWriter, r *http.Request) {
	// lista todos os inputs
}

func (Input) Create(w http.ResponseWriter, r *http.Request) {
	// cria um novo input
}

func (Input) Delete(w http.ResponseWriter, r *http.Request) {
	// deleta um input
}

func NewInput() Input {
	return Input{}
}
