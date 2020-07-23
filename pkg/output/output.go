// Package output Для различного вывода информации
package output

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
)

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
