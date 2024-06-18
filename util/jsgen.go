package util

import (
    "fmt"
    "os"
)

// CreateJSFile создает JS файл с заданным именем и содержимым.
func CreateJSFile(fileName, content string) error {
    err := os.MkdirAll("static/js", os.ModePerm)
    if err != nil {
        return fmt.Errorf("error creating js directory: %v", err)
    }

    filePath := "static/js/" + fileName
    _, err = os.Stat(filePath)
    if os.IsNotExist(err) {
        file, err := os.Create(filePath)
        if err != nil {
            return fmt.Errorf("error creating JS file: %v", err)
        }
        defer file.Close()

        _, err = file.WriteString(content)
        if err != nil {
            return fmt.Errorf("error writing to JS file: %v", err)
        }
    } else if err == nil {
        // Файл уже существует, не возвращаем ошибку, просто выходим из функции
        return nil
    } else {
        return fmt.Errorf("error checking if file exists: %v", err)
    }
    return nil
}

// CreateDefaultJSFile создает стандартный JS файл script.js, если он не существует.
func CreateDefaultJSFile() error {
    content := `document.addEventListener("DOMContentLoaded", function() {
    console.log("Hello, World!");
});`
    return CreateJSFile("script.js", content)
}


