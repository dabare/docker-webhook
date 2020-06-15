package webhooks

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type PushData struct {
	Pushed_at int      `json:"pushed_at"`
	images    []string `json:"images"`
	Tag       string   `json:"tag"`
	Pusher    string   `json:"pusher"`
}
type Repository struct {
	Status          string `json:"status"`
	Description     string `json:"description"`
	IsTrusted       bool   `json:"is_trusted"`
	FullDescription string `json:"full_description"`
	RepoUrl         string `json:"repo_url"`
	Owner           string `json:"owner"`
	IsOfficial      bool   `json:"is_official"`
	IsPrivate       bool   `json:"is_private"`
	Name            string `json:"name"`
	Namespace       string `json:"namespace"`
	StarCount       int    `json:"star_count"`
	CommentCount    int    `json:"comment_count"`
	DateCreated     int    `json:"date_created"`
	RepoName        string `json:"repo_name"`
}
type DockerWebHook struct {
	PushData    PushData   `json:"push_data"`
	CallbackUrl string     `json:"callback_url"`
	Repository  Repository `json:"repository"`
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/web-hook", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.POST("/web-hook", func(c *gin.Context) {

		body := c.Request.Body
		x, _ := ioutil.ReadAll(body)

		log.Println(string(x))
		c.String(http.StatusOK, "ok")
	})
	return r
}
