/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"imagectl/pkg/imageaction"
	"imagectl/pkg/settings"
)

// getimageCmd represents the getimage command
var getimageCmd = &cobra.Command{
	Use:   "getprojectname",
	Short: "获取harbor项目列表",
	Long:  `获取harbor项目列表`,
	Run: func(cmd *cobra.Command, args []string) {
		imageaction.GetProjectsName(settings.Srcurl, settings.Srcname, settings.Srcpassword)
	},
}

func init() {
	rootCmd.AddCommand(getimageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getimageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getimageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
