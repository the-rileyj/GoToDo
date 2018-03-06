package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type credentials struct {
	Nasa string `json:"nasa"`
}

type user struct {
	UID   string `json:"uid"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Slack string `json:"slack"`
}

type todo struct {
	Description string `json:"description"`
	Title       string `json:"title"`

	Date      string `json:"date"`
	EndTime   string `json:"endTime"`
	StartTime string `json:"startTime"`

	Email bool `json:"email"`
	Phone bool `json:"phone"`
	Slack bool `json:"slack"`

	UID string `json:"uid"`
}

func getCredentials(file string) (credentials, error) {
	c := credentials{}
	if raw, err := ioutil.ReadFile(file); err == nil {
		json.Unmarshal(raw, &c)
		return c, nil
	} else {
		return credentials{}, fmt.Errorf("Could not read JSON from %s: %v", file, err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) { http.ServeFile(c.Writer, c.Request, "./index.html") })
	r.GET("/test", func(c *gin.Context) { http.ServeFile(c.Writer, c.Request, "./test.html") })
	r.GET("/static/css/:fi", static.Serve("/static/css", static.LocalFile("static/css/", true)))
	r.GET("/static/img/:fi", static.Serve("/static/img", static.LocalFile("static/img/", true)))
	r.GET("/static/js/:fi", static.Serve("/static/js", static.LocalFile("static/js/", true)))
	r.GET("/static/custom/:fi", static.Serve("/static/custom", static.LocalFile("static/custom/", true)))

	r.GET("/creds/nasa", func(c *gin.Context) {
		c.JSON(200, gin.H{"key": "value"})
	})

	r.Run(":5566")
}
