package main

import (
	"errors"
	"net/http"
)

/*
 * In Go there are no optional/default function arguments.
 * There are a couple of different ways of how to handle optionality.
 */

// 1. Args + struct
// Mandatory arguments are regular arguments, optional arguments are kept within a struct

type ArgsStructOptionalConfig struct {
	Port *int // pointers because otherwise optional values would get initialized with zero values
}

func NewArgsStructServer(address string, config *ArgsStructOptionalConfig) (*http.Server, error) {
	var port int
	if config == nil {
		port = 3000
	} else {
		if config.Port == nil {
			port = 3000
		}
		// and so on
	}
	_ = port
	return &http.Server{}, nil
}

func ArgsStructOptionalConfigExamples() {
	explicitPort := 3000 // minor inconvenience, cannot inline pointer arguments
	_, _ = NewArgsStructServer("http://localhost", &ArgsStructOptionalConfig{
		Port: &explicitPort,
	})
	_, _ = NewArgsStructServer("http://localhost", &ArgsStructOptionalConfig{})
	_, _ = NewArgsStructServer("http://localhost", nil)
}

// 2. Builder
// Optional arguments are exposed as methods on Builder, defaults building is decopupled from the constructor function

type Config struct {
	Port int
}

type ConfigBuilder struct {
	port *int
}

func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	if b.port == nil {
		cfg.Port = 3000
	} else {
		if *b.port == 0 {
			cfg.Port = 1234
		} else if *b.port < 0 {
			return Config{}, errors.New("wrong port number")
		} else {
			cfg.Port = *b.port
		}
	}

	return cfg, nil
}

func NewBulderServer(address string, config Config) (*http.Server, error) {
	return &http.Server{}, nil
}

func ConfigBuilderExamples() {
	builder := ConfigBuilder{}
	config, _ := builder.Port(8080).Build()
	_, _ = NewBulderServer("http://localhost", config)
}

// this pattern is a classic in OOP languages, but some of its implementations are not well suited for Go
// especially since Go doesn't have thrown exceptions, so a method for an option can't return errors
// and be chainable at the same time

// 3. Functional options
// expose a set of functions to manipulate optional values
// this way is most idiomatic to Go

// `options` is a private type so that options handling is under control of the package that exposes the options
type options struct {
	port *int
}

type Option func(options *options) error

func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("wrong port number")
		}
		options.port = &port
		return nil
	}
}

func NewFunctionalOptionsServer(address string, opts ...Option) (*http.Server, error) {
	var options options
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}

	// at this point, options are initialized and can be further processed
	var port int
	if options.port == nil {
		port = 3000
	} else {
		if *options.port == 0 {
			port = 1234
		} else {
			port = *options.port
		}
	}
	_ = port
	return &http.Server{}, nil
}

func OptionsExamples() {
	_, _ = NewFunctionalOptionsServer("http://localhost", WithPort(3030))
	_, _ = NewFunctionalOptionsServer("http://localhost") // variadic arguments are truly optional!
}
