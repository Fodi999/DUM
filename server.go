package main

import (
    "dum/components"
    "dum/middlewares"
    "dum/router"
    "dum/util"
    "fmt"
    "log"
    "net/http"
    "path/filepath"
    "os"
    "sync"
    "time"
)

var (
    mu        sync.Mutex
    lastMod   time.Time
    clients   = make(map[chan bool]struct{})
    r         = router.NewRouter()
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
        modTime := getLastModificationTime("templates", "static/css")
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

func addRoute(path string) {
    r.AddRoute(path, "GET", components.RenderPageHandler)
}

func startServer() {
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

    // Создание стандартного HTML файла hello.html
    err := util.CreateDefaultHTMLFile()
    if err != nil {
        log.Fatalf("Error creating default HTML file: %v", err)
    }

    // Создание стандартного JS файла script.js
    err = util.CreateDefaultJSFile()
    if err != nil {
        log.Fatalf("Error creating default JS file: %v", err)
    }

    // Генерация CSS файлов
    util.GenerateCSS()

    // Добавление маршрутов для существующих файлов
    addRoute("/")
    addRoute("/about")
    addRoute("/contact")
    addRoute("/user")

    // Обработка статических файлов
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Настройка обработчиков с middleware
    chain := middlewares.CORS(middlewares.Logging(r))
    http.Handle("/", chain)

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
}

