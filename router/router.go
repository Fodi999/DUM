package router

import (
    "net/http"
)

// Route - структура для хранения маршрутов
type Route struct {
    Path    string
    Method  string
    Handler http.HandlerFunc
}

// Router - структура для маршрутизатора
type Router struct {
    routes []Route
}

// NewRouter - создает новый экземпляр маршрутизатора
func NewRouter() *Router {
    return &Router{routes: []Route{}}
}

// AddRoute - добавляет маршрут в маршрутизатор
func (r *Router) AddRoute(path, method string, handler http.HandlerFunc) {
    r.routes = append(r.routes, Route{Path: path, Method: method, Handler: handler})
}

// ServeHTTP - обрабатывает HTTP-запросы
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    for _, route := range r.routes {
        if route.Path == req.URL.Path && route.Method == req.Method {
            route.Handler(w, req)
            return
        }
    }
    http.NotFound(w, req)
}
