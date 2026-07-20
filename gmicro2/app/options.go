package app

import "net/url"

type Options struct {
	Endpoints []url.URL
	ID        string
	Name      string
}

type Option func(o *Options)

func WithEndpoints(endpoints []url.URL) Option {
	return func(o *Options) {
		o.Endpoints = endpoints
	}
}

func WithID(id string) Option {
	return func(o *Options) {
		o.ID = id
	}
}

func WithName(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}
