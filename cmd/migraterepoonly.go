/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"imagectl/pkg/imageaction"
	"imagectl/pkg/settings"
)

// migraterepoonlyCmd represents the migraterepoonly command
var migraterepoonlyCmd = &cobra.Command{
	Use:   "migraterepoonly",
	Short: "迁移单个项目所有仓库",
	Long: `迁移单个项目所有仓库:
  ./imagectl migraterepoonly -c test01 --dst-har-url http://192.168.153.11/test02`,
	Run: func(cmd *cobra.Command, args []string) {
		catprojectname, _ := cmd.Flags().GetString("catprojectname")
		dst_url, _ := cmd.Flags().GetString("dst-har-url")
		dst_username, _ := cmd.Flags().GetString("dst-har-username")
		dst_password, _ := cmd.Flags().GetString("dst-har-password")
		imageaction.Migrepoonly(settings.Srcurl, settings.Srcname, settings.Srcpassword, catprojectname, dst_url, dst_username, dst_password)
	},
}

func init() {
	rootCmd.AddCommand(migraterepoonlyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migraterepoonlyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migraterepoonlyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	migraterepoonlyCmd.Flags().String("dst-har-url", "192.168.153.11/library", "输入目标仓库镜像地址,不要加http或https")
	migraterepoonlyCmd.Flags().String("dst-har-username", "admin", "输入目标地址的用户名")
	migraterepoonlyCmd.Flags().String("dst-har-password", "Harbor12345", "输入目标地址的密码")
}
