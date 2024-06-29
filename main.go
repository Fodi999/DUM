// main.go
package main

import (
    "bufio"
    "dum/util"
    "fmt"
    "os"
    "os/signal"
    "path/filepath"
    "strings"
    "sync"
    "syscall"
    "os/exec"
    "log"
)

var wg sync.WaitGroup
var quit = make(chan os.Signal, 1)

func main() {
    loadEnvVariables()

    startServer()

    startWebSocketServer()
    startWebSocketServerAbout()

    go monitorCommands()

    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    fmt.Println("Shutting down server...")
    close(broadcast)
    wg.Wait()
}

func monitorCommands() {
    reader := bufio.NewReader(os.Stdin)
    for {
        printPrompt()
        command, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("\033[31mError reading command: \033[0m", err)
            continue
        }
        command = strings.TrimSpace(command)
        args := strings.Split(command, " ")

        switch args[0] {
        case "quit":
            fmt.Println("\033[32mShutting down server...\033[0m")
            quit <- syscall.SIGTERM
            return
        case "reload":
            fmt.Println("\033[32mReloading server...\033[0m")
            go watchFiles()
        case "status":
            fmt.Println("\033[32mServer is running...\033[0m")
        case "create":
            handleCreateFile(reader)
        case "create_bot":
            handleCreateBot(reader)
        case "start_bot":
            if len(args) > 1 {
                handleStartBots(args[1:])
            } else {
                fmt.Println("\033[31mPlease specify bot names to start.\033[0m")
            }
        case "list_bots":
            listBots()
        case "create_site":
            handleCreateSite(reader)
        case "start_site":
            if len(args) > 1 {
                handleStartSite(args[1:])
            } else {
                fmt.Println("\033[31mPlease specify site names to start.\033[0m")
            }
        default:
            fmt.Println("\033[31mUnknown command: \033[0m", command)
        }
    }
}

func printPrompt() {
    fmt.Print("> ")
}

func handleCreateFile(reader *bufio.Reader) {
    fmt.Print("Enter file name (e.g., about.html): ")
    fileName, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("\033[31mError reading file name: \033[0m", err)
        return
    }
    fileName = strings.TrimSpace(fileName)
    if fileName == "" {
        fmt.Println("\033[31mInvalid file name. Please try again.\033[0m")
        return
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
        return
    }

    jsContent := `document.addEventListener("DOMContentLoaded", function() {
    console.log("` + strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ` page loaded");
});`
    err = util.CreateJSFile(strings.TrimSuffix(fileName, filepath.Ext(fileName))+".js", jsContent)
    if err != nil {
        fmt.Println("\033[31mError creating JS file: \033[0m", err)
        return
    }

    addRoute("/" + strings.TrimSuffix(fileName, filepath.Ext(fileName)))
    fmt.Println("\033[32mFiles created successfully: \033[0m", fileName, "and", strings.TrimSuffix(fileName, filepath.Ext(fileName))+".js")
}

func handleCreateBot(reader *bufio.Reader) {
    fmt.Print("Enter bot name: ")
    botName, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("\033[31mError reading bot name: \033[0m", err)
        return
    }
    botName = strings.TrimSpace(botName)
    if botName == "" {
        fmt.Println("\033[31mInvalid bot name. Please try again.\033[0m")
        return
    }

    fmt.Print("Enter bot token: ")
    botToken, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("\033[31mError reading bot token: \033[0m", err)
        return
    }
    botToken = strings.TrimSpace(botToken)
    if botToken == "" {
        fmt.Println("\033[31mInvalid bot token. Please try again.\033[0m")
        return
    }

    err = util.CreateRobotFolder(botName)
    if err != nil {
        fmt.Println("\033[31mError creating bot folder: \033[0m", err)
        return
    }

    err = util.GenerateBotCode(botName)
    if err != nil {
        fmt.Println("\033[31mError generating bot code: \033[0m", err)
        return
    }

    bots[botName] = botToken
    fmt.Println("\033[32mBot created successfully: \033[0m", botName)
}

func handleStartBots(botNames []string) {
    for _, botName := range botNames {
        botName = strings.TrimSpace(botName)
        if botName == "" {
            fmt.Println("\033[31mInvalid bot name. Please try again.\033[0m")
            continue
        }

        botToken, exists := bots[botName]
        if !exists {
            fmt.Printf("\033[31mBot %s not found. Please create the bot first.\033[0m\n", botName)
            continue
        }

        wg.Add(1)
        go StartBot(botToken, &wg, broadcast)
        fmt.Printf("\033[32mBot %s started.\033[0m\n", botName)
    }
}

func listBots() {
    fmt.Println("Created bots:")
    for botName := range bots {
        fmt.Printf(" - %s\n", botName)
    }
}

func handleCreateSite(reader *bufio.Reader) {
    fmt.Print("Enter site name: ")
    siteName, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("\033[31mError reading site name: \033[0m", err)
        return
    }
    siteName = strings.TrimSpace(siteName)
    if siteName == "" {
        fmt.Println("\033[31mInvalid site name. Please try again.\033[0m")
        return
    }

    err = util.CreateSite(siteName)
    if err != nil {
        fmt.Println("\033[31mError creating site: \033[0m", err)
        return
    }

    fmt.Printf("\033[32mSite %s created successfully.\033[0m\n", siteName)
}

func handleStartSite(siteNames []string) {
    for _, siteName := range siteNames {
        siteName = strings.TrimSpace(siteName)
        if siteName == "" {
            fmt.Println("\033[31mInvalid site name. Please try again.\033[0m")
            continue
        }

        port := util.GetSitePort(siteName)
        cmd := exec.Command("go", "run", "sites/"+siteName+"/main.go")
        cmd.Env = append(os.Environ(), fmt.Sprintf("PORT=%d", port))
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        err := cmd.Start()
        if err != nil {
            fmt.Printf("\033[31mError starting site %s: %v\033[0m\n", siteName, err)
            continue
        }

        fmt.Printf("\033[32mSite %s started on http://localhost:%d\033[0m\n", siteName, port)
    }
}

func logAndPrint(message string) {
    log.Println(message)
    printPrompt()
}

func printAndLog(message string) {
    fmt.Println(message)
    printPrompt()
}





