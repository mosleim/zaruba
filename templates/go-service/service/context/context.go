package context

import (
	"registry.com/user/servicename/communication"
	"registry.com/user/servicename/config"
)

// NewContext initiate new Context
func NewContext() (context *Context) {
	return &Context{
		Config: config.NewConfig(),
		Status: &Status{
			IsAlive: true,
			IsReady: true,
		},
		LocalCache: communication.Message{},
	}
}

// Context is a general service context
type Context struct {
	Config     *config.Config
	Status     *Status
	LocalCache communication.Message
}

// InitLocalCache set LocalCache
func (context *Context) InitLocalCache(key string, val interface{}) {
	if _, exists := context.LocalCache[key]; !exists {
		context.LocalCache[key] = val
	}
}

// Status represent liveness and readiness of service
type Status struct {
	IsAlive bool
	IsReady bool
}
