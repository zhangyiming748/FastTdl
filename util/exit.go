package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var exit bool

// SetExit 从控制台读取输入，当输入 "q" 时提示确认
// 确认后有 10 秒等待时间，超时自动退出；输入 "n" 取消退出
// 此函数应该在独立的 goroutine 中调用
func SetExit() {
	lines := make(chan string, 1)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "读取输入时发生错误: %v\n", err)
		}
		close(lines)
	}()

	fmt.Println("请输入 'q' 退出程序...")
	for line := range lines {
		input := strings.TrimSpace(line)
		if strings.ToLower(input) == "q" {
			fmt.Print("确认退出？(y/n) 10秒内未确认将自动退出: ")
			select {
			case answer, ok := <-lines:
				if !ok || strings.TrimSpace(strings.ToLower(answer)) == "y" {
					exit = true
					fmt.Println("\n收到退出信号，程序即将退出...")
					return
				}
				fmt.Println("已取消退出，继续运行...")
			case <-time.After(10 * time.Second):
				exit = true
				fmt.Println("\n等待超时，程序即将退出...")
				return
			}
			fmt.Println("请输入 'q' 退出程序...")
		}
	}
}

// GetExit 返回 exit 状态
func GetExit() bool {
	return exit
}
