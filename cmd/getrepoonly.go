/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"imagectl/pkg/imageaction"
	"imagectl/pkg/settings"
)

// getrepoonlyCmd represents the getrepoonly command
var getrepoonlyCmd = &cobra.Command{
	Use:   "getrepoonly",
	Short: "查询某个项目的镜像列表",
	Long:  `查询某个项目的镜像列表`,
	Run: func(cmd *cobra.Command, args []string) {
		catprojectname, _ := cmd.Flags().GetString("catprojectname")
		imageaction.Getrepoonly(settings.Srcurl, settings.Srcname, settings.Srcpassword, catprojectname)
	},
}

func init() {
	rootCmd.AddCommand(getrepoonlyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getrepoonlyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getrepoonlyCmd.Flags().StringP("catprojectname", "c", "library", "输入要查询的项目")
}
