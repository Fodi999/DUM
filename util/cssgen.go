package util

import (
    "fmt"
    "os"
    "path/filepath"
    "regexp"
    "strings"
)

func GenerateCSS() {
    err := os.MkdirAll("static/css", os.ModePerm)
    if err != nil {
        fmt.Println("Error creating static/css directory:", err)
        return
    }

    htmlFiles, err := filepath.Glob("templates/*.html")
    if err != nil {
        fmt.Println("Error reading html files", err)
        return
    }
    classSet := make(map[string]struct{})
    classRegex := regexp.MustCompile(`class="([^"]+)"`)
    for _, file := range htmlFiles {
        content, err := os.ReadFile(file)
        if err != nil {
            fmt.Println("Error reading file:", err)
            return
        }
        matches := classRegex.FindAllSubmatch(content, -1)
        for _, match := range matches {
            classes := strings.Split(string(match[1]), " ")
            for _, class := range classes {
                classSet[class] = struct{}{}
            }
        }
    }

    cssClasses := map[string]string{
        // Fonts
        "font-sans": `font-family: ui-sans-serif, system-ui, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";`,
        "font-serif": `font-family: ui-serif, Georgia, Cambria, "Times New Roman", Times, serif;`,
        "font-mono": `font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;`,
        
        // Text colors
        "text-gray-700": "color: #4a5568;",
        "text-red-500": "color: #f56565;",
        "text-center": "text-align: center;",
        "text-orange-500": "color: rgb(249 115 22);",
        "text-sky-800": "color: rgb(7 89 133);",

        // Padding
        "p-2": "padding: 8px;",

        // Background colors
        "bg-black": "background-color: rgb(0 0 0);",
        "bg-transparent": "background-color: transparent;",
        "bg-current": "background-color: currentColor;",
        "bg-zinc-500": "background-color: rgb(113 113 122);",
        "bg-red-500": "background-color: rgb(239 68 68);",
        "bg-orange-500": "background-color: rgb(249 115 22);",
        "bg-yellow-500": "background-color: rgb(234 179 8);",
        "bg-cyan-500": "background-color: rgb(6 182 212);",
        "bg-indigo-500": "background-color: rgb(99 102 241);",
        "bg-rose-500": "background-color: rgb(244 63 94);",
    }

    file, err := os.Create("static/css/style.css")
    if err != nil {
        fmt.Println("Error creating CSS file:", err)
        return
    }
    defer file.Close()

    for class := range classSet {
        if style, exists := cssClasses[class]; exists {
            _, err := file.WriteString(fmt.Sprintf(".%s {%s}\n", class, style))
            if err != nil {
                fmt.Println("Error writing to CSS file:", err)
                return
            }
        }
    }
    globalStyle := "body, html {margin: 0; padding: 0; font-family: 'Arial', sans-serif; height: 100%;}\n\n@keyframes slideIn {\n\tfrom { transform: translateX(-100%); opacity: 0; }\n\tto { transform: translateX(0); opacity: 1; }\n}\n"
    _, err = file.WriteString(globalStyle)
    if err != nil {
        fmt.Println("Error writing to CSS file:", err)
        return
    }
}


