/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"imagectl/pkg/settings"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "imagectl",
	Short: "image tools",
	Long:  `image tools`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.imagectl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.PersistentFlags().StringVarP(&settings.Srcname, "src-username", "u", "admin", "输入源地址的用户名")
	rootCmd.PersistentFlags().StringVarP(&settings.Srcpassword, "src-password", "p", "Harbor12345", "输入源地址的密码")
	rootCmd.PersistentFlags().StringVarP(&settings.Srcurl, "src-url", "l", "192.168.153.11", "输入源地址的URL")
	rootCmd.PersistentFlags().StringVarP(&settings.File, "file", "f", "imageList.txt", "输入你要保存到哪个文件")
}
