package components

import (
    "log"
    "net/http"
    "html/template"
    "path"
)

// RenderPageHandler рендерит страницы на основе имени файла
func RenderPageHandler(w http.ResponseWriter, r *http.Request) {
    page := path.Base(r.URL.Path)
    if page == "/" || page == "" {
        page = "hello.html"
    } else {
        page += ".html"
    }
    log.Printf("Rendering page: %s", page)
    tmpl, err := template.ParseFiles("templates/" + page)
    if err != nil {
        log.Printf("Error parsing template: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, nil)
    if err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

