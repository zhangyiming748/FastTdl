package main

import (
	"FastTdl/core"
	"FastTdl/rotate"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhangyiming748/GracefullyExit"
	"log"
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

	// 创建根命令
	var rootCmd = &cobra.Command{
		Use:   "my-tdl",
		Short: "FastTdl CLI工具",
		Long:  "一个用于文件下载和归档的命令行工具",
	}

	// 创建下载命令
	var tdlCmd = &cobra.Command{
		Use:   "tdl",
		Short: "下载文件",
		Long:  "根据指定的根目录和链接文件下载文件",
		Run: func(cmd *cobra.Command, args []string) {
			root, _ := cmd.Flags().GetString("dir")
			postlink, _ := cmd.Flags().GetString("input")
			proxy, _ := cmd.Flags().GetString("proxy")
			fmt.Printf("开始执行下载任务...\n根目录:%s\n链接文件: %s\n", root, postlink)
			core.Tdl(root, postlink, proxy)
		},
	}

	// 为 tdl 命令添加标志
	tdlCmd.Flags().StringP("dir", "d", "./", "根目录路径 (必需)")
	tdlCmd.Flags().StringP("input", "i", "./post.link", "链接文件路径 (必需)")
	tdlCmd.Flags().StringP("proxy", "p", "http://127.0.0.1:8889", "代理地址")

	// 设置必选标志
	tdlCmd.MarkFlagRequired("dir")
	tdlCmd.MarkFlagRequired("input")
	tdlCmd.MarkFlagRequired("proxy")

	// 创建归档命令
	var archiveCmd = &cobra.Command{
		Use:   "archive",
		Short: "归档文件",
		Long:  "归档指定目录下的所有文件",
		Run: func(cmd *cobra.Command, args []string) {
			dir, _ := cmd.Flags().GetString("dir")
			fhd, _ := cmd.Flags().GetBool("fhd")

			fmt.Printf("开始执行归档任务...\n目录:%s\nFHD 模式：%v\n", dir, fhd)
			core.ArchiveAllFiles(dir, fhd)
		},
	}

	// 为 archive 命令添加标志
	archiveCmd.Flags().StringP("dir", "d", "./", "要归档的目录路径 (必需)")
	archiveCmd.Flags().BoolP("fhd", "f", false, "是否裁剪到1080p")

	// 设置必填标志
	archiveCmd.MarkFlagRequired("dir")

	// 创建旋转命令
	var rotateCmd = &cobra.Command{
		Use:   "rotate",
		Short: "旋转视频文件",
		Long:  "旋转指定目录下的所有视频文件",
		Run: func(cmd *cobra.Command, args []string) {
			dir, _ := cmd.Flags().GetString("dir")
			rotateDirection, _ := cmd.Flags().GetString("rotate")
			fmt.Printf("开始执行视频旋转任务...\n目录:%s\n方向:%s度\n", dir, rotateDirection)
			rotate.RotateVideos(dir, rotateDirection)
		},
	}

	// 为 rotate 命令添加标志
	rotateCmd.Flags().StringP("dir", "d", "./", "要旋转视频的目录路径 (必需)")
	rotateCmd.Flags().StringP("rotate", "r", "90", "旋转方向:90,270")

	// 设置必填标志
	rotateCmd.MarkFlagRequired("dir")

	// 将子命令添加到根命令
	rootCmd.AddCommand(tdlCmd)
	rootCmd.AddCommand(archiveCmd)
	rootCmd.AddCommand(rotateCmd)
	go GracefullyExit.StartReceivedExit()
	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("执行命令出现致命错误:%v\n", err)
	}
}
