//components/hello.go
package components

import (
    "net/http"
    "html/template"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("components/html/hello.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("components/html/about.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
func ContactHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("components/html/contact.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}