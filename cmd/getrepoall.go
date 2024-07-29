/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"imagectl/pkg/imageaction"
	"imagectl/pkg/settings"

	"github.com/spf13/cobra"
)

// getrepoallCmd represents the getrepoall command
var getrepoallCmd = &cobra.Command{
	Use:   "getrepoall",
	Short: "获取所有镜像列表",
	Long:  "获取所有镜像列表",
	Run: func(cmd *cobra.Command, args []string) {

		imageaction.Getrepoall(settings.Srcurl, settings.Srcname, settings.Srcpassword, settings.File)
	},
}

func init() {
	rootCmd.AddCommand(getrepoallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getrepoallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getrepoallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
