package utils

import (
	"os"
	"bufio"
	"fmt"
)
func GetCli() string{
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("输入你的命令：")
	scanner.Scan()
	return scanner.Text()
}