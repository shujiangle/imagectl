package harborapi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseurl    string
	username   string
	password   string
	httpclient *http.Client
}

type Project struct {
	ProjectID    int       `json:"project_id"`
	Name         string    `json:"name"`
	CreationTime time.Time `json:"creation_time"`
}

type Repository struct {
	RepositoryID int       `json:"id"`
	Name         string    `json:"name"`
	CreationTime time.Time `json:"creation_time"`
}

type Repositorytag struct {
	Name string `json:"name"`
}

func NewClient(baseurl, username, password string) *Client {
	// - InsecureSkipVerify: 在 tls.Config 中将 InsecureSkipVerify 设置为 true 会跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &Client{
		baseurl:  baseurl,
		username: username,
		password: password,
		httpclient: &http.Client{
			Transport: tr,
			Timeout:   30 * time.Second,
		},
	}
}

func (c *Client) doRequest(req *http.Request, result interface{}) error {
	req.SetBasicAuth(c.username, c.password)
	resp, err := c.httpclient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed request: status code %d, response: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, result)
}

func (c *Client) GetProjects() ([]Project, error) {
	apiURL := fmt.Sprintf("%s/api/projects", c.baseurl)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	var projects []Project
	if err := c.doRequest(req, &projects); err != nil {
		return nil, err
	}

	//fmt.Printf("projects:%v", projects)
	return projects, nil
}

func (c *Client) GetRepositories(projectID int) ([]Repository, error) {
	apiURL := fmt.Sprintf("%s/api/repositories?project_id=%d", c.baseurl, projectID)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	var repositories []Repository
	if err := c.doRequest(req, &repositories); err != nil {
		return nil, err
	}

	return repositories, nil
}

func (c *Client) GetRepositoriesTag(repositoriestag string) ([]Repositorytag, error) {
	apiURL := fmt.Sprintf("%s/api/repositories/%v/tags", c.baseurl, repositoriestag)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	var tags []Repositorytag
	if err := c.doRequest(req, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}
