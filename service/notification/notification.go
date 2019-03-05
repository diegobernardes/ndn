package notification

import (
	"context"

	"github.com/diegobernardes/ndn"
)

type Pusher struct {
}

func (p Pusher) Push(ctx context.Context, n ndn.Notification) error {
	return nil
}
