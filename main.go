package main

import (
    "dum/components"
    "dum/middlewares"
    "dum/router"
    "dum/util"
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "sync"
    "time"
)

var (
    mu        sync.Mutex
    lastMod   time.Time
    clients   = make(map[chan bool]struct{})
)

var spinnerChars = []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'}
var colors = []string{
    "\033[36m", // Cyan
    "\033[35m", // Magenta
    "\033[33m", // Yellow
    "\033[32m", // Green
    "\033[34m", // Blue
    "\033[31m", // Red
    "\033[0m",  // Reset
}

// watchFiles следит за изменениями в указанных директориях и обновляет CSS.
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

// getLastModificationTime возвращает время последней модификации файлов в директориях.
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

// fileWatcherHandler обрабатывает запросы на обновление файлов.
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
    // Приветственное сообщение в квадрате с URL
    fmt.Println("\033[35m" + `
+-------------------------------------+
|                                     |
|  DUM фреймворк приветствует вас     |
|          версия 1.0                 |
|                                     |
|     http://localhost:8080           |
|                                     |
+-------------------------------------+
` + "\033[0m")

    // Канал для управления спиннером
    s := make(chan bool)
    // Запуск спиннера
    go func() {
        i := 0
        for {
            for _, color := range colors {
                select {
                case <-s:
                    return
                default:
                    fmt.Printf("\r%s%s Initializing...\033[0m", color, string(spinnerChars[i%len(spinnerChars)]))
                    time.Sleep(100 * time.Millisecond)
                    i++
                }
            }
        }
    }()

    // Генерация CSS файлов
    util.GenerateCSS()

    // Создание маршрутизатора
    r := router.NewRouter()
    r.AddRoute("/", "GET", components.HelloHandler)
    r.AddRoute("/about", "GET", components.AboutHandler)
    r.AddRoute("/contact", "GET", components.ContactHandler)
    r.AddRoute("/watch", "GET", fileWatcherHandler)

    // Обработка статических файлов
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Настройка обработчиков с middleware
    http.Handle("/", middlewares.Logging(r))

    // Запуск наблюдения за файлами
    go watchFiles()

    // Запуск сервера в отдельной горутине
    go func() {
        server := &http.Server{
            Addr: ":8080",
        }

        log.Print("")
        err := server.ListenAndServe()
        if err != nil {
            log.Fatal("\033[31mError starting server: \033[0m", err)
        }
    }()

    // Остановка спиннера и вывод сообщения об успешном запуске
    time.Sleep(3 * time.Second)
    s <- true
    fmt.Println("")

    // Мониторинг команд в терминале
    monitorCommands()
}

// monitorCommands следит за командами в терминале для управления сервером.
func monitorCommands() {
    for {
        var command string
        fmt.Print("Enter command: ")
        _, err := fmt.Scanln(&command)
        if err != nil {
            fmt.Println("\033[31mError reading command: \033[0m", err)
            continue
        }

        switch command {
        case "quit":
            fmt.Println("\033[32mShutting down server...\033[0m")
            os.Exit(0)
        default:
            fmt.Println("\033[31mUnknown command: \033[0m", command)
        }
    }
}











