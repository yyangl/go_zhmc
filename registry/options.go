package registry

import (
	"context"
	"time"
)

type Options struct {
	Addrs   []string
	Timeout time.Duration
	Context context.Context
}

type RegisterOptions struct {
	TTL     time.Duration
	Context context.Context
}

type WatchOptions struct {
	Service string
	Context context.Context
}

func Addrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}
func RegisterTTL(t time.Duration) RegisterOption {
	return func(o *RegisterOptions) {
		o.TTL = t
	}
}

// Watch a service
func WatchService(name string) WatchOption {
	return func(o *WatchOptions) {
		o.Service = name
	}
}
