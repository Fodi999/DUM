package components

import (
    "net/http"
    "html/template"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("components/hello.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
