package imageaction

import (
	"fmt"
	"imagectl/pkg/harborapi"
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
		defer file.Close()
	}

}
