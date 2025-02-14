package pkg

import (
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// ServiceRegistry manages cleanup functions for resources.
type ServiceRegistry struct {
	mu       sync.Mutex
	cleanups []func() error
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		cleanups: make([]func() error, 0),
	}
}

func (sr *ServiceRegistry) Register(cleanup func() error) {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	sr.cleanups = append(sr.cleanups, cleanup)
}

// Close invokes all registered cleanup functions in reverse order.
func (sr *ServiceRegistry) close() error {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	var cleanupErrors []error
	for i := len(sr.cleanups) - 1; i >= 0; i-- {
		if err := sr.cleanups[i](); err != nil {
			cleanupErrors = append(cleanupErrors, err)
		}
	}
	return errors.Join(cleanupErrors...)
}

func (sr *ServiceRegistry) WaitForSignal() error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig
	return sr.close()
}
