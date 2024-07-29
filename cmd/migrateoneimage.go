/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"imagectl/pkg/migrate"
	"imagectl/pkg/settings"
)

// migrateoneimageCmd represents the migrateoneimage command
var migrateoneimageCmd = &cobra.Command{
	Use:   "migrateoneimage",
	Short: "迁移单个仓库",
	Long: `迁移单个仓库：
./imagectl migrateoneimage -l http://192.168.153.11/library/centos:latest -L http://192.168.153.11/test01/centos:v3`,
	Run: func(cmd *cobra.Command, args []string) {
		dst_url, _ := cmd.Flags().GetString("dst-url")
		dst_username, _ := cmd.Flags().GetString("dst-username")
		dst_password, _ := cmd.Flags().GetString("dst-password")
		settings.Srcurl = migrate.ExtractIP(settings.Srcurl)
		dst_url = migrate.ExtractIP(dst_url)
		cmdshell := fmt.Sprintf("skopeo copy --policy=policy.json --src-creds='%s:%s' --dest-creds='%s:%s' --src-tls-verify=%t --dest-tls-verify=%t docker://%s docker://%s",
			settings.Srcname, settings.Srcpassword, dst_username, dst_password, false, false, settings.Srcurl, dst_url)

		//fmt.Printf("%s\n", cmdshell)
		migrate.Command(cmdshell)
	},
}

func init() {
	rootCmd.AddCommand(migrateoneimageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateoneimageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateoneimageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	migrateoneimageCmd.Flags().StringP("dst-url", "L", "192.168.153.11/library/centos:latest", "输入目标仓库镜像地址,不要加http或https")
	migrateoneimageCmd.Flags().StringP("dst-username", "U", "admin", "输入目标地址的用户名")
	migrateoneimageCmd.Flags().StringP("dst-password", "P", "Harbor12345", "输入目标地址的密码")
}
