# imagectl 工具


```shell
[root@localhost ~]# imagectl  -h
image tools

Usage:
  imagectl [flags]
  imagectl [command]

Available Commands:
  completion     Generate the autocompletion script for the specified shell
  getprojectname 获取harbor项目列表
  getrepoall     获取所有镜像列表
  getrepoonly    查询某个项目的镜像列表
  help           Help about any command

Flags:
  -f, --file string           输入你要保存到哪个文件 (default "imageList.txt")
  -h, --help                  help for imagectl
  -p, --src-password string   输入源地址的密码 (default "Harbor12345")
  -l, --src-url string        输入源地址的URL (default "http://192.168.153.11")
  -u, --src-username string   输入源地址的用户名 (default "admin")

Use "imagectl [command] --help" for more information about a command.

```
