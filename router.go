package simple

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleware func(httprouter.Handle) httprouter.Handle

// Router is a http.Handler that wraps httprouter.Router with additional features.
type Router struct {
	middlewares []middleware
	path        string
	router      *httprouter.Router
}

// New returns *Router with a new initialized *httprouter.Router embedded.
func New() *Router {
	return &Router{router: httprouter.New()}
}

func (r *Router) joinPath(path string) string {
	if (r.path + path)[0] != '/' {
		panic("path should start with '/' in path '" + path + "'.")
	}

	return r.path + path
}

// Group returns new *Router with given path and middlewares.
// It should be used for handles which have same path prefix or common middlewares.
func (r *Router) Group(path string, m ...middleware) *Router {
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	return &Router{
		middlewares: append(m, r.middlewares...),
		path:        r.joinPath(path),
		router:      r.router,
	}
}

// Use appends new middleware to current Router.
func (r *Router) Use(m ...middleware) *Router {
	r.middlewares = append(m, r.middlewares...)
	return r
}

// Handle registers a new request handle combined with middlewares.
func (r *Router) Handle(method, path string, handle httprouter.Handle) {
	for _, v := range r.middlewares {
		handle = v(handle)
	}
	r.router.Handle(method, r.joinPath(path), handle)
}

// GET is a shortcut for Router.Handle("GET", path, handle)
func (r *Router) GET(path string, handle httprouter.Handle) {
	r.Handle("GET", path, handle)
}

// HEAD is a shortcut for Router.Handle("HEAD", path, handle)
func (r *Router) HEAD(path string, handle httprouter.Handle) {
	r.Handle("HEAD", path, handle)
}

// OPTIONS is a shortcut for Router.Handle("OPTIONS", path, handle)
func (r *Router) OPTIONS(path string, handle httprouter.Handle) {
	r.Handle("OPTIONS", path, handle)
}

// POST is a shortcut for Router.Handle("POST", path, handle)
func (r *Router) POST(path string, handle httprouter.Handle) {
	r.Handle("POST", path, handle)
}

// PUT is a shortcut for Router.Handle("PUT", path, handle)
func (r *Router) PUT(path string, handle httprouter.Handle) {
	r.Handle("PUT", path, handle)
}

// PATCH is a shortcut for Router.Handle("PATCH", path, handle)
func (r *Router) PATCH(path string, handle httprouter.Handle) {
	r.Handle("PATCH", path, handle)
}

// DELETE is a shortcut for Router.Handle("DELETE", path, handle)
func (r *Router) DELETE(path string, handle httprouter.Handle) {
	r.Handle("DELETE", path, handle)
}

// Handler is an adapter for http.Handler.
func (r *Router) Handler(method, path string, handler http.Handler) {
	handle := func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		handler.ServeHTTP(w, req)
	}
	r.Handle(method, path, handle)
}

// HandlerFunc is an adapter for http.HandlerFunc.
func (r *Router) HandlerFunc(method, path string, handler http.HandlerFunc) {
	r.Handler(method, path, handler)
}

// Static serves files from given root directory.
func (r *Router) Static(path, root string) {
	if len(path) < 10 || path[len(path)-10:] != "/*filepath" {
		panic("path should end with '/*filepath' in path '" + path + "'.")
	}

	base := r.joinPath(path[:len(path)-9])
	fileServer := http.StripPrefix(base, http.FileServer(http.Dir(root)))

	r.Handler("GET", path, fileServer)
}

// File serves the named file.
func (r *Router) File(path, name string) {
	r.HandlerFunc("GET", path, func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, name)
	})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
