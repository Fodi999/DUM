//robotgen.go содержит функции для генерации кода бота Telegram.
package util

import (
    "fmt"
    "os"
)

// CreateRobotFolder создает новую папку для бота с заданным именем.
func CreateRobotFolder(name string) error {
    path := "robot/" + name
    if _, err := os.Stat(path); !os.IsNotExist(err) {
        return fmt.Errorf("folder %s already exists", path)
    }
    if err := os.MkdirAll(path, os.ModePerm); err != nil {
        return fmt.Errorf("error creating folder: %v", err)
    }
    return nil
}

// GenerateBotCode генерирует код бота Telegram.
func GenerateBotCode(name string) error {
    path := "robot/" + name + "/bot.go"
    botCode := `package main

import (
    "encoding/json"
    "log"
    "net/http"
    "net/url"
    "os"
    "strconv"
    "sync"
    "time"
    "strings"
)

var telegramAPI = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN_" + strings.ToUpper("` + name + `"))

type Update struct {
    UpdateID int ` + "`json:\"update_id\"`" + `
    Message  struct {
        MessageID int ` + "`json:\"message_id\"`" + `
        From      struct {
            ID        int    ` + "`json:\"id\"`" + `
            FirstName string ` + "`json:\"first_name\"`" + `
            LastName  string ` + "`json:\"last_name\"`" + `
            Username  string ` + "`json:\"username\"`" + `
        } ` + "`json:\"from\"`" + `
        Chat struct {
            ID int64 ` + "`json:\"id\"`" + `
        } ` + "`json:\"chat\"`" + `
        Text string ` + "`json:\"text\"`" + `
    } ` + "`json:\"message\"`" + `
}

func StartBot(wg *sync.WaitGroup) {
    defer wg.Done()
    updatesChan := make(chan Update)

    // Горутина для получения обновлений
    go func() {
        offset := 0
        for {
            updates, err := getUpdates(offset)
            if err != nil {
                log.Println("Error getting updates:", err)
                time.Sleep(1 * time.Second)
                continue
            }

            for _, update := range updates {
                updatesChan <- update
                offset = update.UpdateID + 1
            }
        }
    }()

    // Горутина для обработки обновлений
    go func() {
        for update := range updatesChan {
            if update.Message.Text != "" {
                log.Printf("[%s] %s", update.Message.From.Username, update.Message.Text)
                wg.Add(1)
                go func(update Update) {
                    defer wg.Done()
                    err := sendMessage(update.Message.Chat.ID, "Hello, "+update.Message.From.FirstName)
                    if err != nil {
                        log.Println("Error sending message:", err)
                    }
                }(update)
            }
        }
    }()

    // Ожидание завершения всех горутин
    wg.Wait()
}

func getUpdates(offset int) ([]Update, error) {
    resp, err := http.Get(telegramAPI + "/getUpdates?offset=" + strconv.Itoa(offset))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        OK     bool     ` + "`json:\"ok\"`" + `
        Result []Update ` + "`json:\"result\"`" + `
    }

    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return nil, err
    }

    return result.Result, nil
}

func sendMessage(chatID int64, text string) error {
    values := url.Values{}
    values.Set("chat_id", strconv.FormatInt(chatID, 10))
    values.Set("text", text)

    resp, err := http.PostForm(telegramAPI+"/sendMessage", values)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    var result struct {
        OK bool ` + "`json:\"ok\"`" + `
    }

    return json.NewDecoder(resp.Body).Decode(&result)
}
`
    if err := os.WriteFile(path, []byte(botCode), 0644); err != nil {
        return fmt.Errorf("error writing bot code: %v", err)
    }
    return nil
}


