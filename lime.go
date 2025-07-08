package lime

import "net/http"

// Iroutes define all router handle interface.
type Iroutes interface {
	GET(path string, handler http.HandlerFunc) Iroutes
	POST(path string, handler http.HandlerFunc) Iroutes
	DELETE(path string, handler http.HandlerFunc) Iroutes
	PUT(path string, handler http.HandlerFunc) Iroutes
	OPTIONS(path string, handler http.HandlerFunc) Iroutes
	HEAD(path string, handler http.HandlerFunc) Iroutes
}

type OptionFunc func(*Engine)

// Engine is the framework's instance
type Engine struct {
	router map[string]http.HandlerFunc
}

var _ Iroutes = (*Engine)(nil)

// New creates a new Engine instance.
func New(opts ...OptionFunc) *Engine {
	engine := &Engine{
		router: make(map[string]http.HandlerFunc),
	}

	return engine.With(opts...)
}

// With returns a new Engine instance with the given options.
func (engine *Engine) With(opts ...OptionFunc) *Engine {
	for _, opt := range opts {
		opt(engine)
	}
	return engine
}

// ServeHTTP implements the http.Handler interface.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (engine *Engine) addRoute(method, path string, handlers ...http.HandlerFunc) {
	key := method + "-" + path
	for _, handler := range handlers {
		engine.router[key] = handler
	}
}

// Run attatches the router to a http server and starts listening for connections.
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) GET(path string, handler http.HandlerFunc) Iroutes {
	engine.addRoute(http.MethodGet, path, handler)
	return engine
}

func (engine *Engine) POST(path string, handler http.HandlerFunc) Iroutes {
	engine.addRoute(http.MethodPost, path, handler)
	return engine
}

func (engine *Engine) DELETE(path string, handler http.HandlerFunc) Iroutes {
	engine.addRoute(http.MethodDelete, path, handler)
	return engine
}

func (engine *Engine) PUT(path string, handler http.HandlerFunc) Iroutes {
	engine.addRoute(http.MethodPut, path, handler)
	return engine
}

func (engine *Engine) OPTIONS(path string, handler http.HandlerFunc) Iroutes {
	engine.addRoute(http.MethodOptions, path, handler)
	return engine
}

func (engine *Engine) HEAD(path string, handler http.HandlerFunc) Iroutes {
	engine.addRoute(http.MethodHead, path, handler)
	return engine
}
