package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "sync"
    "time"
    "dum/middlewares"
    "dum/components"
    "dum/util"
)

var (
    mu        sync.Mutex
    lastMod   time.Time
    clients   = make(map[chan bool]struct{})
)

func watchFiles() {
    for {
        time.Sleep(1 * time.Second)
        modTime := getLastModificationTime("components", "static/css")
        mu.Lock()
        if modTime.After(lastMod) {
            lastMod = modTime
            for ch := range clients {
                select {
                case ch <- true:
                default:
                }
                close(ch)
                delete(clients, ch)
            }
            util.GenerateCSS() // Вызов генерации CSS после изменения
        }
        mu.Unlock()
    }
}

func getLastModificationTime(dirs ...string) time.Time {
    var latestMod time.Time
    for _, dir := range dirs {
        _ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if info.ModTime().After(latestMod) {
                latestMod = info.ModTime()
            }
            return nil
        })
    }
    return latestMod
}

func fileWatcherHandler(w http.ResponseWriter, r *http.Request) {
    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
        return
    }

    notify := make(chan bool)
    mu.Lock()
    clients[notify] = struct{}{}
    mu.Unlock()

    for {
        select {
        case <-notify:
            w.Header().Set("Content-Type", "text/event-stream")
            w.Header().Set("Cache-Control", "no-cache")
            w.Header().Set("Connection", "keep-alive")
            _, _ = w.Write([]byte("data: reload\n\n"))
            flusher.Flush()
            return
        case <-r.Context().Done():
            mu.Lock()
            delete(clients, notify)
            mu.Unlock()
            return
        }
    }
}

func main() {
    util.GenerateCSS()

    http.Handle("/", middlewares.Logging(http.HandlerFunc(components.HelloHandler)))
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/watch", fileWatcherHandler)

    go watchFiles()

    // Запуск сервера в отдельной горутине
    go func() {
        server := &http.Server{
            Addr: ":8080",
        }

        log.Printf("Server started at http://localhost:8080")
        err := server.ListenAndServe()
        if err != nil {
            log.Fatal(err)
        }
    }()

    // Мониторинг команд в терминале
    monitorCommands()
}

func monitorCommands() {
    for {
        var command string
        _, err := fmt.Scanln(&command)
        if err != nil {
            log.Println("Error reading command:", err)
            continue
        }

        switch command {
        case "quit":
            log.Println("Shutting down server...")
            os.Exit(0)
        default:
            log.Println("Unknown command:", command)
        }
    }
}


