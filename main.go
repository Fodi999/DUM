package main

import (
    "bufio"
    "dum/util"
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    // Запуск сервера
    startServer()

    // Мониторинг команд в терминале
    monitorCommands()
}

// monitorCommands следит за командами в терминале для управления сервером.
func monitorCommands() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter command: ")
        command, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("\033[31mError reading command: \033[0m", err)
            continue
        }
        command = strings.TrimSpace(command)

        switch command {
        case "quit":
            fmt.Println("\033[32mShutting down server...\033[0m")
            os.Exit(0)
        case "reload":
            fmt.Println("\033[32mReloading server...\033[0m")
            go watchFiles()
        case "status":
            fmt.Println("\033[32mServer is running...\033[0m")
        case "create":
            fmt.Print("Enter file name (e.g., about.html): ")
            fileName, err := reader.ReadString('\n')
            if err != nil {
                fmt.Println("\033[31mError reading file name: \033[0m", err)
                continue
            }
            fileName = strings.TrimSpace(fileName)
            if fileName == "" {
                fmt.Println("\033[31mInvalid file name. Please try again.\033[0m")
                continue
            }
            htmlContent := `<!DOCTYPE html>
<html>
    <head>
        <link href="/static/css/style.css" rel="stylesheet">
        <title>` + fileName + `</title>
    </head>
    <body>
        <nav>
            | <a href="/">Home</a> 
            | <a href="/about">About</a> 
            | <a href="/contact">Contact</a> 
            | <a href="/user">User</a>
        </nav>
        <h1>` + strings.TrimSuffix(fileName, filepath.Ext(fileName)) + `</h1>
        <script src="/static/js/` + strings.TrimSuffix(fileName, filepath.Ext(fileName)) + `.js"></script>
    </body>
</html>`
            err = util.CreateHTMLFile(fileName, htmlContent)
            if err != nil {
                fmt.Println("\033[31mError creating HTML file: \033[0m", err)
                continue
            }

            jsContent := `document.addEventListener("DOMContentLoaded", function() {
    console.log("` + strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ` page loaded");
});`
            err = util.CreateJSFile(strings.TrimSuffix(fileName, filepath.Ext(fileName))+".js", jsContent)
            if err != nil {
                fmt.Println("\033[31mError creating JS file: \033[0m", err)
                continue
            }

            addRoute("/" + strings.TrimSuffix(fileName, filepath.Ext(fileName)))
            fmt.Println("\033[32mFiles created successfully: \033[0m", fileName, "and", strings.TrimSuffix(fileName, filepath.Ext(fileName))+".js")
        case "create_bot":
            fmt.Print("Enter bot name: ")
            botName, err := reader.ReadString('\n')
            if err != nil {
                fmt.Println("\033[31mError reading bot name: \033[0m", err)
                continue
            }
            botName = strings.TrimSpace(botName)
            if botName == "" {
                fmt.Println("\033[31mInvalid bot name. Please try again.\033[0m")
                continue
            }

            err = util.CreateRobotFolder(botName)
            if err != nil {
                fmt.Println("\033[31mError creating bot folder: \033[0m", err)
                continue
            }

            err = util.GenerateBotCode(botName)
            if err != nil {
                fmt.Println("\033[31mError generating bot code: \033[0м", err)
                continue
            }

            fmt.Println("\033[32mBot created successfully: \033[0m", botName)
        default:
            fmt.Println("\033[31mUnknown command: \033[0m", command)
        }
    }
}







































































 












































































































 
































































































 






























































































 
















































































 





























































































 






















































































































 




























































































 

























































































 
























