package main

import (
	"FastTdl/archive"
	"FastTdl/core"
	"FastTdl/rotate"
	"fmt"
	"log"

	"FastTdl/tdl"
	"FastTdl/util"

	"github.com/spf13/cobra"
	a "github.com/zhangyiming748/archive"
)

// 版本信息，通过 ldflags 在构建时注入
var (
	version   = "dev"     // 默认版本号
	buildTime = "unknown" // 构建时间
	gitCommit = "unknown" // Git 提交哈希
)

func main() {
	// 统一初始化
	Initialize()

	// 创建根命令
	var rootCmd = &cobra.Command{
		Use:   "my-tdl",
		Short: "FastTdl CLI Tool",
		Long:  "A command-line tool for file download and archiving",
	}

	// 添加版本标志
	rootCmd.Version = fmt.Sprintf("%s (built at %s, commit %s)", version, buildTime, gitCommit)
	rootCmd.SetVersionTemplate("FastTdl version {{.Version}}\n")
	// 创建下载命令
	var tdlCmd = &cobra.Command{
		Use:   "tdl",
		Short: "Download files",
		Long:  "Download files based on specified root directory and link file",
		Run: func(cmd *cobra.Command, args []string) {
			root, _ := cmd.Flags().GetString("dir")
			postlink, _ := cmd.Flags().GetString("input")
			proxy, _ := cmd.Flags().GetString("proxy")
			fmt.Printf("Starting download task...\nRoot directory: %s\nLink file: %s\n", root, postlink)
			core.Tdl(root, postlink, proxy)
		},
	}

	// 为 tdl 命令添加标志
	tdlCmd.Flags().StringP("dir", "d", "./", "Root directory path (default: current directory)")
	tdlCmd.Flags().StringP("input", "i", "./post.link", "Link file path (default: post.link in current directory)")
	tdlCmd.Flags().StringP("proxy", "p", "http://127.0.0.1:8889", "Proxy address")

	// 设置必选标志
	// tdlCmd.MarkFlagRequired("input")

	// 创建归档命令
	var archiveCmd = &cobra.Command{
		Use:   "archive",
		Short: "Archive files",
		Long:  "Archive all files in specified directory",
		Run: func(cmd *cobra.Command, args []string) {
			dir, _ := cmd.Flags().GetString("dir")
			fhd, _ := cmd.Flags().GetBool("fhd")
			fmt.Printf("Starting archive task...\nDirectory: %s\nFHD mode: %v\n", dir, fhd)
			a.InitSqlte()
			core.ArchiveAllFiles(dir, fhd)
		},
	}

	// 为 archive 命令添加标志
	archiveCmd.Flags().StringP("dir", "d", "./", "Directory path to archive (required)")
	archiveCmd.Flags().BoolP("fhd", "f", false, "Crop to 1080p")

	// 设置必填标志
	archiveCmd.MarkFlagRequired("dir")

	// 创建旋转命令
	var rotateCmd = &cobra.Command{
		Use:   "rotate",
		Short: "Rotate video files",
		Long:  "Rotate all video files in specified directory",
		Run: func(cmd *cobra.Command, args []string) {
			dir, _ := cmd.Flags().GetString("dir")
			rotateDirection, _ := cmd.Flags().GetString("rotate")
			fmt.Printf("Starting video rotation task...\nDirectory: %s\nDirection: %s degrees\n", dir, rotateDirection)
			rotate.RotateVideos(dir, rotateDirection)
		},
	}

	// 为 rotate 命令添加标志
	rotateCmd.Flags().StringP("dir", "d", "./", "Directory path for video rotation (required)")
	rotateCmd.Flags().StringP("rotate", "r", "90", "Rotation direction: 90, 270")

	// 设置必填标志
	rotateCmd.MarkFlagRequired("dir")

	// 创建 DJI 处理命令
	var djiCmd = &cobra.Command{
		Use:   "dji",
		Short: "Process DJI video files",
		Long:  "Convert DJI drone videos to H265 format",
		Run: func(cmd *cobra.Command, args []string) {
			src, _ := cmd.Flags().GetString("src")
			dst, _ := cmd.Flags().GetString("dst")
			fmt.Printf("Starting DJI video processing task...\nSource directory: %s\nTarget directory: %s\n", src, dst)
			archive.Dji(src, dst)
		},
	}

	// 为 dji 命令添加标志
	djiCmd.Flags().StringP("src", "i", "./", "Source directory path (required)")
	djiCmd.Flags().StringP("dst", "o", "./output", "Target directory path (required)")

	// 设置必填标志
	djiCmd.MarkFlagRequired("src")
	djiCmd.MarkFlagRequired("dst")

	// 将子命令添加到根命令
	rootCmd.AddCommand(tdlCmd)
	rootCmd.AddCommand(archiveCmd)
	rootCmd.AddCommand(rotateCmd)
	rootCmd.AddCommand(djiCmd)
	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("执行命令出现致命错误:%v\n", err)
	}
}

// Initialize 统一初始化函数，在 main 函数开始时调用
func Initialize() {
	log.Println("开始初始化 FastTdl...")

	// 1. 初始化时区
	util.InitTimeZone()
	log.Println("时区初始化完成")

	// 2. 初始化日志系统
	util.SetLog("tdl.log")
	log.Println("日志系统初始化完成")

	// 3. 加载中英文映射表
	tdl.LoadTranslationMap()
	log.Println("中英文映射表加载完成")

	log.Println("FastTdl 初始化完成")
}
