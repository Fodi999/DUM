package router

import (
    "net/http"
    "regexp"
)

// Route - структура для хранения маршрутов
type Route struct {
    Path    string
    Method  string
    Handler http.HandlerFunc
    Regex   *regexp.Regexp
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
    re := regexp.MustCompile("^" + path + "$")
    r.routes = append(r.routes, Route{Path: path, Method: method, Handler: handler, Regex: re})
}

// ServeHTTP - обрабатывает HTTP-запросы
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    for _, route := range r.routes {
        if route.Method == req.Method && route.Regex.MatchString(req.URL.Path) {
            route.Handler(w, req)
            return
        }
    }
    http.NotFound(w, req)
}

