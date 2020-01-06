package options

import "sync"

// Options is used for initialisation
type Options interface {
	// Initialise options
	Init(...Option) error
	// Options returns the current options
	Values() *Values
	// The name for who these options exist
	String() string
}

// Values holds the set of option values and protects them
type Values struct {
	sync.RWMutex
	values map[interface{}]interface{}
}

// Option gives access to options
type Option func(o *Values) error
