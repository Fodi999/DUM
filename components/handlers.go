//handlers.go содержит обработчики запросов для веб-страниц
package components

import (
    "log"
    "net/http"
    "html/template"
    "path"
    "path/filepath"
)

// RenderPageHandler рендерит страницы на основе имени файла
func RenderPageHandler(w http.ResponseWriter, r *http.Request) {
    page := path.Clean(r.URL.Path)
    if page == "/" || page == "" {
        page = "hello.html"
    } else {
        page = page[1:] + ".html"
    }
    log.Printf("Rendering page: %s", page)
    tmpl, err := template.ParseFiles(filepath.Join("templates", page))
    if err != nil {
        log.Printf("Error parsing template: %v", err)
        http.Error(w, "Error rendering page", http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, nil)
    if err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, "Error executing template", http.StatusInternalServerError)
    }
}


