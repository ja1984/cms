package routes

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ja1984/cogCMS/backend/database"
)

const pagePath = "page:%s:%s"

type Request struct {
	Data Page `json:"Data"`
}
type Page struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Slug     string `json:"slug"`
	Value    string `json:"value"`
	Required bool   `json:"required"`
	Tooltip  string `json:"tooltip"`
}

func CreatePage(c *gin.Context) {
	var data Request

	c.BindJSON(&data)

	err := savePage(data.Data)

	if err != nil {
		c.AbortWithError(500, err)
	}

	c.Status(201)
}

func GetPages(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func savePage(page Page) error {
	pageJSON, err := json.Marshal(page)

	if err != nil {
		return err
	}

	return database.REDIS.Set(fmt.Sprintf(pagePath, page.ID, page.Name), pageJSON, 0).Err()
}
