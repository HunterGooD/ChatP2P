package input

import (
	"bufio"
	"os"
	"strings"
)

// IUser пользовательский ввод
func IUser() string {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	return strings.TrimSpace(str)
}

// IFile читает даные из файла и возвращает пользователю
func IFile(fileName string) []string {

	return []string{""}
}
