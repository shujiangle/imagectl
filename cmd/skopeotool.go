/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"imagectl/pkg/migrate"
	"os"
)

func createPolicyFile() {
	file, err := os.Create("policy.json")
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("文件已存在，报错退出")
			os.Exit(1)
		} else {
			fmt.Println("创建文件时出错:", err)
			os.Exit(1)
		}
	}
	defer file.Close()

	// 准备要写入的内容
	data := []byte(`{
    "default": [
        {
            "type": "insecureAcceptAnything"
        }
    ],
    "transports":
        {
            "docker-daemon":
                {
                    "": [{"type":"insecureAcceptAnything"}]
                }
        }
}`)

	// 写入文件
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("写入文件时出错:", err)
		os.Exit(1)
	}

	fmt.Println("policy.json文件创建完成")
	migrate.Command("chmod +x skopeo")
	migrate.Command("mv -f skopeo /usr/local/bin")
	fmt.Println("skopeo文件创建完成,可以通过skopeo -h查看")
}

// skopeotoolCmd represents the skopeotool command
var skopeotoolCmd = &cobra.Command{
	Use:   "skopeotool",
	Short: "下载skopeo命令和创建policy.json文件",
	Long:  "下载skopeo命令和创建policy.json文件",
	Run: func(cmd *cobra.Command, args []string) {
		cmdshell := fmt.Sprintf("curl -fL 'https://pixiupkg-generic.pkg.coding.net/pixiu/k8stoolsdefault/skopeo?version=latest' -o skopeo")
		fmt.Println("开始下载skopeo")
		migrate.Command(cmdshell)
		createPolicyFile()
	},
}

func init() {
	rootCmd.AddCommand(skopeotoolCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// skopeotoolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// skopeotoolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
