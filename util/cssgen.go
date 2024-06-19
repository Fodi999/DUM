package util

import (
    "fmt"
    "os"
    "path/filepath"
    "regexp"
    "strings"
    "dum/util/library"
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

    cssClasses := map[string]string{}

    // Add styles from separate files
    library.AddTextColors(cssClasses)
    library.AddTypography(cssClasses) // Предполагается, что функция принимает аргумент cssClasses
    library.AddBackgroundColors(cssClasses)
    library.AddBorderStyles(cssClasses)
    library.AddEffects(cssClasses)
    library.AddLayout(cssClasses)
    library.AddBackgrounds(cssClasses)
    library.AddGradientColorStops(cssClasses)
    library. AddFlexboxGrid(cssClasses)
    library.AddBorder(cssClasses)
    library.AddSizing(cssClasses)
    library.AddSpacing(cssClasses)
    library.AddBorderColorStyles(cssClasses)
    library.AddEffectsColors(cssClasses)
    library. AddFilters(cssClasses)
    library.AddTables(cssClasses)
    library.AddTransitionsAndAnimations(cssClasses)
    library.AddTransforms(cssClasses)
    library.AddInteractivity(cssClasses)


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
    globalStyle := ""
    _, err = file.WriteString(globalStyle)
    if err != nil {
        fmt.Println("Error writing to CSS file:", err)
        return
    }
}
