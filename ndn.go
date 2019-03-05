package ndn

import "context"

// Kinds of input.
var (
	InputKindHTTPPush = "http.push"
)

// Kinds of output.
var (
	OutputKindHTTPPull = "http.pull"
)

// Notification holds the value to represent a notification.
type Notification struct {
	Entity  string
	ID      string
	Version string
}

// NotificationPusher represents the struct that gonna do the notification pusher.
type NotificationPusher interface {
	Push(ctx context.Context, n Notification) error
}

// NotificationPusherError represet all kind of errors that can be returned from the NotificationPusher.
type NotificationPusherError interface {
	Error() string
	AlreadyExists() bool
}
