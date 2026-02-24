package main

import (
	"FastTdl/core"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	/*
		这里也使用cobra改成cli命令行工具
		go build -o main main.go编译二进制文件之后‘
		main   tdl --root "/Users/zen/gitea/FastTdl" --postlink "/Users/zen/gitea/FastTdl/des.link"
		实际上要运行的函数是FastTdl.Tdl(root, postlink)
		main archive --root "/Users/zen/gitea/FastTdl"
		实际上要运行的函数是
		FastTdl.ArchiveAllFiles(mainFolder)
	*/
	var rootCmd = &cobra.Command{
		Use:   "fasttdl",
		Short: "FastTdl CLI工具",
		Long:  "一个用于文件下载和归档的命令行工具",
	}
	var tdlCmd = &cobra.Command{
		Use:   "tdl",
		Short: "下载文件",
		Long:  "根据指定的根目录和链接文件下载文件",
		Run: func(cmd *cobra.Command, args []string) {
			root, _ := cmd.Flags().GetString("root")
			postlink, _ := cmd.Flags().GetString("postlink")
			proxy, _ := cmd.Flags().GetString("proxy")
			if root == "" {
				fmt.Println("错误: 必须指定 --root 参数")
				return
			}
			if postlink == "" {
				fmt.Println("错误: 必须指定 --postlink 参数")
				return
			}
			fmt.Printf("开始执行下载任务...\n根目录: %s\n链接文件: %s\n", root, postlink)
			core.Tdl(root, postlink,proxy)
		},
	}
	var archiveCmd = &cobra.Command{
		Use:   "archive",
		Short: "归档文件",
		Long:  "归档指定目录下的所有文件",
		Run: func(cmd *cobra.Command, args []string) {
			root, _ := cmd.Flags().GetString("root")
			if root == "" {
				fmt.Println("错误: 必须指定 --root 参数")
				return
			}
			fmt.Printf("开始执行归档任务...\n目录: %s\n", root)
			core.ArchiveAllFiles(root)
		},
	}
	// 为tdl命令添加标志
	tdlCmd.Flags().String("root", "./", "根目录路径 (必需)")
	tdlCmd.Flags().String("postlink", "./post.link", "链接文件路径 (必需)")
	tdlCmd.Flags().String("proxy", "http://127.0.0.1:8889", "代理地址 (必需)")
	// 为archive命令添加标志
	archiveCmd.Flags().String("root", "./", "要归档的目录路径 (必需)")
	// 将子命令添加到根命令
	rootCmd.AddCommand(tdlCmd)
	rootCmd.AddCommand(archiveCmd)
	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("执行命令出现致命错误:%v\n", err)
	}
}
