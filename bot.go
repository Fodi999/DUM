// bot.go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "net/url"
    "strconv"
    "sync"
    "time"
)

type Update struct {
    UpdateID int `json:"update_id"`
    Message  struct {
        MessageID int `json:"message_id"`
        From      struct {
            ID        int    `json:"id"`
            FirstName string `json:"first_name"`
            LastName  string `json:"last_name"`
            Username  string `json:"username"`
        } `json:"from"`
        Chat struct {
            ID int64 `json:"id"`
        } `json:"chat"`
        Text string `json:"text"`
    } `json:"message"`
}

func StartBot(botToken string, wg *sync.WaitGroup, broadcast chan Message) {
    defer wg.Done()
    telegramAPI := "https://api.telegram.org/bot" + botToken
    updatesChan := make(chan Update)

    go func() {
        offset := 0
        for {
            updates, err := getUpdates(telegramAPI, offset)
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

    go func() {
        for update := range updatesChan {
            if update.Message.Text != "" {
                log.Printf("[%s] %s", update.Message.From.Username, update.Message.Text)
                broadcast <- Message{Username: update.Message.From.Username, Message: update.Message.Text}
                err := sendMessage(telegramAPI, update.Message.Chat.ID, "Hello, "+update.Message.From.FirstName)
                if err != nil {
                    log.Println("Error sending message:", err)
                }
            }
        }
    }()
}

func getUpdates(telegramAPI string, offset int) ([]Update, error) {
    resp, err := http.Get(telegramAPI + "/getUpdates?offset=" + strconv.Itoa(offset))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        OK     bool     `json:"ok"`
        Result []Update `json:"result"`
    }

    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return nil, err
    }

    return result.Result, nil
}

func sendMessage(telegramAPI string, chatID int64, text string) error {
    values := url.Values{}
    values.Set("chat_id", strconv.FormatInt(chatID, 10))
    values.Set("text", text)

    resp, err := http.PostForm(telegramAPI+"/sendMessage", values)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    var result struct {
        OK bool `json:"ok"`
    }

    return json.NewDecoder(resp.Body).Decode(&result)
}



