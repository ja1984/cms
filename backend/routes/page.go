package routes

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ja1984/cogCMS/backend/database"
)

const pagePath = "page:%s"

type PageRequest struct {
	Data Page `json:"Data"`
}
type Page struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Fields        []Field `json:"fields"`
	Created       string  `json:"created"`
	Updated       string  `json:"updated"`
	Template      string  `json:"template"`
	Published     bool    `json:"published"`
	EnvironmentID int     `json:"environment_id"`
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
	var data PageRequest

	c.BindJSON(&data)

	err := savePage(data.Data)

	if err != nil {
		c.AbortWithError(500, err)
	}

	c.Status(201)
}

func GetPages(c *gin.Context) {
	var cursor uint64
	keysToGet := []string{}
	for {
		var keys []string
		var err error
		keys, cursor, err = database.REDIS.Scan(0, "page:*", 0).Result()
		if err != nil {
			c.AbortWithError(500, err)
		}

		for _, key := range keys {
			keysToGet = append(keysToGet, key)
		}

		if cursor == 0 {
			break
		}
	}

	pages := []Page{}

	for _, key := range keysToGet {
		page, err := getPage(key)
		if err != nil {
			c.AbortWithError(500, err)
		}
		pages = append(pages, *page)

	}

	c.JSON(200, pages)
}

func getPage(key string) (*Page, error) {
	value, err := database.REDIS.Get(key).Result()
	if err != nil {
		return nil, err
	}

	page := &Page{}

	err = json.Unmarshal([]byte(value), page)
	if err != nil {
		return nil, err
	}

	return page, nil
}

func savePage(page Page) error {
	pageJSON, err := json.Marshal(page)

	if err != nil {
		return err
	}

	return database.REDIS.Set(fmt.Sprintf(pagePath, page.ID), pageJSON, 0).Err()
}
