package client

import (
	"time"
)

type optionApplyFunc func(client *baseClient) error

type Option interface {
	applyOption(client *baseClient) error
}

func (f optionApplyFunc) applyOption(p *baseClient) error {
	return f(p)
}

func WithName(name string) Option {
	return optionApplyFunc(func(client *baseClient) error {
		client.name = name
		return nil
	})
}

func SetTimeout(seconds int64) Option {
	return optionApplyFunc(func(client *baseClient) error {
		client.timeout = time.Duration(seconds) * time.Second
		return nil
	})
}
