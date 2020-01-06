package cmd

import "github.com/yyangl/go_zhmc/registry"

type Options struct {
	Name        string
	Version     string
	Description string
	Registry    *registry.Registry

	Registries map[string]func(...registry.Option) registry.Registry
}

func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}
func Description(d string) Option {
	return func(o *Options) {
		o.Description = d
	}
}
func NewRegistry(name string, r func(...registry.Option) registry.Registry) Option {
	return func(o *Options) {
		o.Registries[name] = r
	}
}
