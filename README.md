# ColorPrint

ColorPrinter 是一个简单而灵活的 Go 语言包，用于在终端输出彩色文本。该包支持多种前景色、背景色和文字样式，旨在增强命令行应用程序的可读性和美观性。

## 特性

- **多种颜色支持**：包括常规颜色和亮色，满足不同的需求。
- **样式支持**：支持加粗、斜体和下划线等文本样式。
- **简单易用**：易于集成到你的 Go 项目中，无需复杂的配置。

## 安装

使用 Go Modules 进行安装：

```bash
go get github.com/iwen-conf/colorprint
```

## 使用示例

以下是如何使用 ColorPrinter 包的示例：

```go
package main

import (
    "fmt"
    "github.com/iwen-conf/colorprint/clr"
)

func main() {

    // 打印带有前景色的文本
    fmt.Println(clr.FGColor("This is red text", clr.Red))
    fmt.Println(clr.FGColor("This is green text", clr.Green))

    // 打印带有背景色的文本
    fmt.Println(clr.BGColor("This has a yellow background", clr.BgYellow))
    fmt.Println(clr.BGColor("This has a blue background", clr.BgBlue))

    // 打印带有样式的文本
    fmt.Println(clr.Bold("This is bold text"))
    fmt.Println(clr.Italic("This is italic text"))
    fmt.Println(clr.Underline("This text is underlined"))
}
```

## API 参考

### 颜色常量

- `Red`
- `Green`
- `Yellow`
- `Blue`
- `Magenta`
- `Cyan`
- `White`
- `BrightBlack`
- `BrightRed`
- `BrightGreen`
- `BrightYellow`
- `BrightBlue`
- `BrightMagenta`
- `BrightCyan`
- `BrightWhite`

### 背景颜色常量

- `BgBlack`
- `BgRed`
- `BgGreen`
- `BgYellow`
- `BgBlue`
- `BgMagenta`
- `BgCyan`
- `BgWhite`
- `BgBrightBlack`
- `BgBrightRed`
- `BgBrightGreen`
- `BgBrightYellow`
- `BgBrightBlue`
- `BgBrightMagenta`
- `BgBrightCyan`
- `BgBrightWhite`

### 样式函数

- `Bold(text string) string`：返回加粗文本。
- `Italic(text string) string`：返回斜体文本。
- `Underline(text string) string`：返回带下划线文本。

## 贡献

欢迎任何形式的贡献!

## 许可证

该项目使用 MIT 许可证。有关更多详细信息，请查看 [LICENSE](https://raw.githubusercontent.com/iwen-conf/colorprinter/refs/heads/main/LICENSE) 文件。
