package imageaction

import (
	"fmt"
	"github.com/fatih/color"
	"imagectl/pkg/harborapi"
	"imagectl/pkg/migrate"
	"log"
	"os"
	"strings"
)

func GetProjectsUrl(url string, username string, password string) (client *harborapi.Client) {
	// 去掉url最后/
	url = strings.TrimSuffix(url, "/")

	client = harborapi.NewClient(url, username, password)

	return client

}

// GetProjectsName 获取项目名列表
func GetProjectsName(url string, username string, password string) {

	client := GetProjectsUrl(url, username, password)
	projects, err := client.GetProjects()
	if err != nil {
		log.Fatalf("Error fetching projects: %v", err)
	}

	fmt.Printf("harbor所有项目列表: \n")
	for _, project := range projects {
		fmt.Printf("%s\n", project.Name)
	}
}

func containsString(slice []harborapi.Project, target string) bool {
	for _, item := range slice {
		if strings.Contains(item.Name, target) {
			return true
		}
	}
	return false
}

func Getrepoonly(url string, username string, password string, catprojectname string) {
	url = strings.TrimSuffix(url, "/")
	client := GetProjectsUrl(url, username, password)
	projects, err := client.GetProjects()
	if err != nil {
		log.Fatalf("Error fetching projects: %v", err)
	}

	if !containsString(projects, catprojectname) {
		log.Fatal("查询的项目，不存在，请检查")
	}
	for _, project := range projects {
		if project.Name == catprojectname {
			allnumbers := Getrepoonlynumber(client, project.ProjectID)
			fmt.Printf("%v项目的仓库总数是%v\n", catprojectname, allnumbers)
			repositories, err := client.GetRepositories(project.ProjectID)
			if err != nil {
				log.Fatalf("Error fetching repositories for project ID  %v", err)
			}

			fmt.Printf("%v项目的镜像列表:\n", catprojectname)
			for _, repo := range repositories {
				repositoriestag, err := client.GetRepositoriesTag(repo.Name)
				if err != nil {
					log.Fatal("%s存在", repositoriestag)
				}

				for _, tag := range repositoriestag {
					tagdesc := fmt.Sprintf("%v/%s:%s\n", url, repo.Name, tag.Name)
					fmt.Printf(tagdesc)
				}

			}
		}
	}
}

func Getrepoall(url string, username string, password string, imagesavefile string) {
	url = strings.TrimSuffix(url, "/")
	client := GetProjectsUrl(url, username, password)
	projects, err := client.GetProjects()
	if err != nil {
		log.Fatalf("Error fetching projects: %v", err)
	}

	file, err := os.OpenFile(imagesavefile, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	for _, project := range projects {
		repositories, err := client.GetRepositories(project.ProjectID)
		if err != nil {
			log.Fatalf("Error fetching repositories for project ID  %v", err)
		}

		for _, repo := range repositories {
			repositoriestag, err := client.GetRepositoriesTag(repo.Name)
			if err != nil {
				log.Fatal("%s存在", repositoriestag)
			}

			for _, tag := range repositoriestag {
				tagdesc := fmt.Sprintf("%v/%s:%s\n", url, repo.Name, tag.Name)
				fmt.Printf(tagdesc)
				_, err = file.WriteString(tagdesc)
				if err != nil {
					fmt.Printf("文件写入失败: %s\n", err)
				}
			}

		}

	}
	defer file.Close()
}

func Migrepoonly(url, username, password, catprojectname, dsturl, dstusername, dstpassword string) {
	url = strings.TrimSuffix(url, "/")
	client := GetProjectsUrl(url, username, password)
	projects, err := client.GetProjects()
	if err != nil {
		log.Fatalf("Error fetching projects: %v", err)
	}

	if !containsString(projects, catprojectname) {
		log.Fatal("查询的项目，不存在，请检查")
	}
	for _, project := range projects {
		if project.Name == catprojectname {
			allnumbers := Getrepoonlynumber(client, project.ProjectID)
			fmt.Println("需要迁移的仓库总数是", allnumbers)
			repositories, err := client.GetRepositories(project.ProjectID)
			//fmt.Printf("仓库名:%v, 仓库版本:%T", repositories)
			if err != nil {
				log.Fatalf("Error fetching repositories for project ID  %v", err)
			}

			counts := 0
			//fmt.Printf("%v项目的镜像列表:\n", catprojectname)
			for _, repo := range repositories {
				repositoriestag, err := client.GetRepositoriesTag(repo.Name)
				//fmt.Printf("仓库名:%v, 仓库版本:%T", repositories)
				if err != nil {
					log.Fatal("%s存在", repositoriestag)
				}
				for _, tag := range repositoriestag {
					changeurl := migrate.ExtractIP(url)
					changedsturl := migrate.ExtractIP(dsturl)
					changeurl = fmt.Sprintf("%v/%s:%s", changeurl, repo.Name, tag.Name)

					parts := strings.Split(repo.Name, "/")
					changedsturl = fmt.Sprintf("%v/%s:%s", changedsturl, parts[1], tag.Name)
					cmdshell := fmt.Sprintf("skopeo copy --policy=policy.json --src-creds='%s:%s' --dest-creds='%s:%s' --src-tls-verify=%t --dest-tls-verify=%t docker://%s docker://%s",
						username, password, dstusername, dstpassword, false, false, changeurl, changedsturl)

					//fmt.Printf("%s\n", cmdshell)
					//migrate.Command(cmdshell)
					//fmt.Printf(tagdesc)
					//fmt.Printf("%v\n\n", cmdshell)
					counts += 1
					fmt.Printf("开启迁移%v个仓库   ", counts)
					blue := color.New(color.FgBlue)
					_, _ = blue.Print(changeurl)
					fmt.Printf("   目标地址")
					color.Blue(changedsturl)
					migrate.Command(cmdshell)
					fmt.Println()
				}

			}
		}
	}
}

func Getrepoonlynumber(client *harborapi.Client, ProjectID int) int {
	repositories, err := client.GetRepositories(ProjectID)
	//fmt.Printf("仓库名:%v, 仓库版本:%T", repositories)
	if err != nil {
		log.Fatalf("Error fetching repositories for project ID  %v", err)
	}

	number := 0
	//fmt.Printf("%v项目的镜像列表:\n", catprojectname)
	for _, repo := range repositories {

		repositoriestag, err := client.GetRepositoriesTag(repo.Name)
		//fmt.Printf("仓库名:%v, 仓库版本:%T", repositories)
		if err != nil {
			log.Fatal("%s存在", repositoriestag)
		}
		number += len(repositoriestag)
	}

	return number
}
