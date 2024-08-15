package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetCli() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("输入你的命令：")
	scanner.Scan()
	return scanner.Text()
}
