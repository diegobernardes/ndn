package postgres

import "ndn"

type Input struct{}

// List all inputs.
func (Input) List() ([]ndn.Input, error) {
	return nil, nil
}

func NewInput() Input {
	return Input{}
}
