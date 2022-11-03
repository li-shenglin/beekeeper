package socket

import (
	"net"
	"sync"
)

type Application struct {
	instances   []*Instance
	lock        *sync.Mutex
	dispatchMap map[int]Handler
}

var app *Application
var once = sync.Once{}

func GetApplication() *Application {
	once.Do(func() {
		app = &Application{
			lock:        &sync.Mutex{},
			instances:   make([]*Instance, 0),
			dispatchMap: make(map[int]Handler, 10),
		}
	})

	return app
}

func (application *Application) Remove(p *Instance) {
	application.lock.Lock()
	defer application.lock.Unlock()
	for i := range application.instances {
		if application.instances[i] == p {
			x := application.instances
			application.instances = x[0:i]
			if i <= len(x) {
				application.instances = append(application.instances, x[i+1:]...)
			}
			return
		}
	}
}

func (application *Application) Add(conn net.Conn) {
	application.lock.Lock()
	defer application.lock.Unlock()
	instance := NewServerInstance(conn, application.dispatch)
	application.instances = append(application.instances, instance)
	instance.Start()
}

func (application *Application) dispatch(key int, message Param) (any, error) {
	return application.dispatchMap[key].Accept(message)
}

func (application *Application) Router(key int, handler Handler) {
	application.dispatchMap[key] = handler
}
