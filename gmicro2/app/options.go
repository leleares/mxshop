package app

import (
	registry "mxshop/gmicro2/registry"
	"net/url"
	"os"
	"time"
)

type Options struct {
	Endpoints       []url.URL
	ID              string
	Name            string
	Signals         []os.Signal
	Register        registry.RegisterInterface // 允许用户传入自己实现的 struct
	RegisterTimeout time.Duration
	StopTimeout     time.Duration
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

func WithSignal(sigs []os.Signal) Option {
	return func(o *Options) {
		o.Signals = sigs
	}
}

func WithRegisterTimeduration(t time.Duration) Option {
	return func(o *Options) {
		o.RegisterTimeout = t
	}
}

func WithStopTimeduration(t time.Duration) Option {
	return func(o *Options) {
		o.StopTimeout = t
	}
}
