package library

func AddLayout(cssClasses map[string]string) {
    layout := map[string]string{
        "aspect-auto": "aspect-ratio: auto;", // Автоматическое соотношение сторон
        "aspect-square": "aspect-ratio: 1 / 1;", // Соотношение сторон 1:1 (квадрат)
        "aspect-video": "aspect-ratio: 16 / 9;", // Соотношение сторон 16:9 (видео)

        // Ширина контейнера
        "container": "width: 100%;",
        "sm": "max-width: 640px;",
        "md": "max-width: 768px;",
        "lg": "max-width: 1024px;",
        "xl": "max-width: 1280px;",
        "2xl": "max-width: 1536px;",

        // Колонки
        "columns-1": "columns: 1;",
        "columns-2": "columns: 2;",
        "columns-3": "columns: 3;",
        "columns-4": "columns: 4;",
        "columns-5": "columns: 5;",
        "columns-6": "columns: 6;",
        "columns-7": "columns: 7;",
        "columns-8": "columns: 8;",
        "columns-9": "columns: 9;",
        "columns-10": "columns: 10;",
        "columns-11": "columns: 11;",
        "columns-12": "columns: 12;",
        "columns-auto": "columns: auto;",
        "columns-3xs": "columns: 16rem;", // 256px
        "columns-2xs": "columns: 18rem;", // 288px
        "columns-xs": "columns: 20rem;", // 320px
        "columns-sm": "columns: 24rem;", // 384px
        "columns-md": "columns: 28rem;", // 448px
        "columns-lg": "columns: 32rem;", // 512px
        "columns-xl": "columns: 36rem;", // 576px
        "columns-2xl": "columns: 42rem;", // 672px
        "columns-3xl": "columns: 48rem;", // 768px
        "columns-4xl": "columns: 56rem;", // 896px
        "columns-5xl": "columns: 64rem;", // 1024px
        "columns-6xl": "columns: 72rem;", // 1152px
        "columns-7xl": "columns: 80rem;", // 1280px

        // Переносы
        "break-after-auto": "break-after: auto;",
        "break-after-avoid": "break-after: avoid;",
        "break-after-all": "break-after: all;",
        "break-after-avoid-page": "break-after: avoid-page;",
        "break-after-page": "break-after: page;",
        "break-after-left": "break-after: left;",
        "break-after-right": "break-after: right;",
        "break-after-column": "break-after: column;",

        "break-before-auto": "break-before: auto;",
        "break-before-avoid": "break-before: avoid;",
        "break-before-all": "break-before: all;",
        "break-before-avoid-page": "break-before: avoid-page;",
        "break-before-page": "break-before: page;",
        "break-before-left": "break-before: left;",
        "break-before-right": "break-before: right;",
        "break-before-column": "break-before: column;",

        "break-inside-auto": "break-inside: auto;",
        "break-inside-avoid": "break-inside: avoid;",
        "break-inside-avoid-page": "break-inside: avoid-page;",
        "break-inside-avoid-column": "break-inside: avoid-column;",

        // Декорирование блоков
        "box-decoration-clone": "box-decoration-break: clone;",
        "box-decoration-slice": "box-decoration-break: slice;",

        // Размеры боксов
        "box-border": "box-sizing: border-box;",
        "box-content": "box-sizing: content-box;",

        // Отображение
        "block": "display: block;",
        "inline-block": "display: inline-block;",
        "inline": "display: inline;",
        "flex": "display: flex;",
        "inline-flex": "display: inline-flex;",
        "table": "display: table;",
        "inline-table": "display: inline-table;",
        "table-caption": "display: table-caption;",
        "table-cell": "display: table-cell;",
        "table-column": "display: table-column;",
        "table-column-group": "display: table-column-group;",
        "table-footer-group": "display: table-footer-group;",
        "table-header-group": "display: table-header-group;",
        "table-row-group": "display: table-row-group;",
        "table-row": "display: table-row;",
        "flow-root": "display: flow-root;",
        "grid": "display: grid;",
        "inline-grid": "display: inline-grid;",
        "contents": "display: contents;",
        "list-item": "display: list-item;",
        "hidden": "display: none;",

        // Обтекание
        "float-start": "float: inline-start;",
        "float-end": "float: inline-end;",
        "float-right": "float: right;",
        "float-left": "float: left;",
        "float-none": "float: none;",

        // Очистка обтекания
        "clear-start": "clear: inline-start;",
        "clear-end": "clear: inline-end;",
        "clear-left": "clear: left;",
        "clear-right": "clear: right;",
        "clear-both": "clear: both;",
        "clear-none": "clear: none;",

        // Изоляция
        "isolate": "isolation: isolate;",
        "isolation-auto": "isolation: auto;",

        // Помещение объектов
        "object-contain": "object-fit: contain;",
        "object-cover": "object-fit: cover;",
        "object-fill": "object-fit: fill;",
        "object-none": "object-fit: none;",
        "object-scale-down": "object-fit: scale-down;",

        // Позиционирование объектов
        "object-bottom": "object-position: bottom;",
        "object-center": "object-position: center;",
        "object-left": "object-position: left;",
        "object-left-bottom": "object-position: left bottom;",
        "object-left-top": "object-position: left top;",
        "object-right": "object-position: right;",
        "object-right-bottom": "object-position: right bottom;",
        "object-right-top": "object-position: right top;",
        "object-top": "object-position: top;",

        // Переполнение
        "overflow-auto": "overflow: auto;",
        "overflow-hidden": "overflow: hidden;",
        "overflow-clip": "overflow: clip;",
        "overflow-visible": "overflow: visible;",
        "overflow-scroll": "overflow: scroll;",
        "overflow-x-auto": "overflow-x: auto;",
        "overflow-y-auto": "overflow-y: auto;",
        "overflow-x-hidden": "overflow-x: hidden;",
        "overflow-y-hidden": "overflow-y: hidden;",
        "overflow-x-clip": "overflow-x: clip;",
        "overflow-y-clip": "overflow-y: clip;",
        "overflow-x-visible": "overflow-x: visible;",
        "overflow-y-visible": "overflow-y: visible;",
        "overflow-x-scroll": "overflow-x: scroll;",
        "overflow-y-scroll": "overflow-y: scroll;",

        // Поведение при скролле
        "overscroll-auto": "overscroll-behavior: auto;",
        "overscroll-contain": "overscroll-behavior: contain;",
        "overscroll-none": "overscroll-behavior: none;",
        "overscroll-y-auto": "overscroll-behavior-y: auto;",
        "overscroll-y-contain": "overscroll-behavior-y: contain;",
        "overscroll-y-none": "overscroll-behavior-y: none;",
        "overscroll-x-auto": "overscroll-behavior-x: auto;",
        "overscroll-x-contain": "overscroll-behavior-x: contain;",
        "overscroll-x-none": "overscroll-behavior-x: none;",

        // Позиционирование
        "static": "position: static;",
        "fixed": "position: fixed;",
        "absolute": "position: absolute;",
        "relative": "position: relative;",
        "sticky": "position: sticky;",

        // Смещения
        "inset-0": "inset: 0px;",
        "inset-x-0": "left: 0px; right: 0px;",
        "inset-y-0": "top: 0px; bottom: 0px;",
        "start-0": "inset-inline-start: 0px;",
        "end-0": "inset-inline-end: 0px;",
        "top-0": "top: 0px;",
        "right-0": "right: 0px;",

        // Видимость
        "visible": "visibility: visible;",
        "invisible": "visibility: hidden;",
        "collapse": "visibility: collapse;",

        // Z-индексы
        "z-0": "z-index: 0;",
        "z-10": "z-index: 10;",
        "z-20": "z-index: 20;",
        "z-30": "z-index: 30;",
        "z-40": "z-index: 40;",
        "z-50": "z-index: 50;",
        "z-auto": "z-index: auto;",
    }
    for k, v := range layout {
        cssClasses[k] = v
    }
}
