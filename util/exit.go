package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var exit bool

// SetExit 从控制台读取输入，当输入 "q" 时设置 exit 为 true
// 此函数应该在独立的 goroutine 中调用
func SetExit() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if strings.ToLower(input) == "q" {
			exit = true
			fmt.Println("\n收到退出信号，程序即将退出...")
			return
		}
	}
}

// GetExit 返回 exit 状态
func GetExit() bool {
	return exit
}
