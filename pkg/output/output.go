// Package output Для различного вывода информации
package output

import (
	"fmt"
	"math/rand"
	"time"
)

// начало конец
const (
	BEGIN = "\033["
	END   = "m"
)

// Формат текст
const (
	Reset = iota
	Bold
	Faint
	Italic
	Underline
	BackgroudReverse = iota + 2
	Concealed
)

// цвет текста
const (
	Black = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	LBlack = iota + 82
	LRed
	LGreen
	LYellow
	LBlue
	LMagenta
	LCyan
	LWhite
)

var colorArray = []int{
	Black, Red, Green,
	Yellow, Blue, Magenta,
	Cyan, White, LBlack,
	LRed, LGreen, LYellow,
	LBlue, LMagenta,
	LCyan, LWhite,
}

// задний фон
const (
	BgBlack = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Color структура отвечающая за вывод цвета
type Color struct {
	BgColor   string
	ColorText string
}

// Reset сброс цвета консоли
func (c Color) Reset() string {
	return fmt.Sprintf("%s%d%s", BEGIN, Reset, END)
}

// RandText рандомный цвет
func (c *Color) RandText() {
	rand.Seed(time.Now().Unix())
	c.BgColor = fmt.Sprintf("%s%d%s", BEGIN, BgBlack, END)
	a := rand.Intn(14)
	c.ColorText = fmt.Sprintf("%s%d%s", BEGIN, colorArray[a], END)
}

// OPrintln цветовой вывод
func OPrintln(user, data string, c *Color) {
	fmt.Printf("%s[%s]:%s%s\n", c.ColorText, user, c.Reset(), data)
}
