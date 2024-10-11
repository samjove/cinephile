package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/samjove/cinephile/pkg/discovery"
)

type serviceName string
type instanceID string

// Registry defines an in-memory service registry.
type Registry struct {
	sync.RWMutex
	serviceAddrs map[serviceName]map[instanceID]*serviceInstance
}

type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

// NewRegistry creates a new in-memory service registry instance.
func NewRegistry() *Registry {
	return &Registry{serviceAddrs: map[serviceName]map[instanceID]*serviceInstance{}}
}

// Register creates a service record in the registry.
func (r *Registry) Register(ctx context.Context, instance instanceID, serviceName serviceName, hostPort string) error {
    r.Lock()
    defer r.Unlock()
    if _, ok := r.serviceAddrs[serviceName]; !ok {
        r.serviceAddrs[serviceName] = map[instanceID]*serviceInstance{}
    }
    r.serviceAddrs[serviceName][instance] = &serviceInstance{hostPort: hostPort, lastActive: time.Now()}
    return nil
}

// Deregister removes a service record from the registry.
func (r *Registry) Deregister(ctx context.Context, instance instanceID, service serviceName) error {
    r.Lock()
    defer r.Unlock()
    if _, ok := r.serviceAddrs[service]; !ok {
        return nil
    }
    delete(r.serviceAddrs[service], instance)
    return nil
}

// ReportHealthyState is a push mechanism for reporting healthy state to the registry.
func (r *Registry) ReportHealthyState(instance instanceID, service serviceName) error {
    r.Lock()
    defer r.Unlock()
    if _, ok := r.serviceAddrs[service]; !ok {
        return errors.New("service is not registered yet")
    }

    if _, ok := r.serviceAddrs[service][instance]; !ok {
        return errors.New("service instance is not registered yet")
    }
    r.serviceAddrs[service][instance].lastActive = time.Now()
    return nil
}

// ServiceAddresses returns the list of addresses of active instances of the given service.
func (r *Registry) ServiceAddresses(ctx context.Context, service serviceName) ([]string, error) {
    r.RLock()
    defer r.RUnlock()
    if len(r.serviceAddrs[service]) == 0 {
        return nil, discovery.ErrNotFound
	}
    var res []string
    for _, i := range r.serviceAddrs[service] {
        if i.lastActive.Before(time.Now().Add(-5 * time.Second)) {
            continue
        }
        res = append(res, i.hostPort)
    }
    return res, nil
}