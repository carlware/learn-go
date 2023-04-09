package redis

import (
	"errors"
	"sync"
)

var (
	ErrKeyNotFound = errors.New("key does not exist")
)

type Cache interface {
	Set(key string, untyped interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) (interface{}, error)
}

type server struct {
	sync.Mutex
	store map[string]interface{}
}

func New() *server {
	return &server{
		store: map[string]interface{}{},
	}
}

func (s *server) Set(key string, untyped interface{}) error  {
	s.Lock()
	defer s.Unlock()
	s.store[key] = untyped
	return nil
}

func (s *server) Get(key string) (interface{}, error) {
	s.Lock()
	defer s.Unlock()
	untyped, ok := s.store[key]
	if !ok {
		return nil, ErrKeyNotFound
	}
	return untyped, nil
}

func (s server) Delete(key string) (interface{}, error) {
	s.Lock()
	defer s.Unlock()
	untyped, ok := s.store[key]
	if !ok {
		return nil, ErrKeyNotFound
	}
	delete(s.store, key)
	return untyped, nil
}