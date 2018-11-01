package routes

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ja1984/cogCMS/backend/database"
)

const templatePath = "template:%s"

type TemplateRequest struct {
	Data Template `json:"Data"`
}
type Template struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Fields []TemplateField `json:"fields"`
}

type TemplateField struct {
	ID   string    `json:"id"`
	Type string    `json:"type"`
	Data FieldData `json:"data"`
}

type FieldData struct {
	ChildFields []TemplateField `json:"childFields"`
	Name        string          `json:"name"`
	Options     []FieldOption   `json:"options"`
	Required    bool            `json:"required"`
	Slug        string          `json:"slug"`
	Tooltip     string          `json:"tooltip"`
}

type FieldOption struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

func CreateTemplate(c *gin.Context) {
	var data TemplateRequest

	c.BindJSON(&data)

	err := saveTemplate(data.Data)

	if err != nil {
		c.AbortWithError(500, err)
	}

	c.Status(201)
}

func GetTemplates(c *gin.Context) {
	var cursor uint64
	keysToGet := []string{}
	for {
		var keys []string
		var err error
		keys, cursor, err = database.REDIS.Scan(0, "template:*", 0).Result()
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

	templates := []Template{}

	for _, key := range keysToGet {
		template, err := getTemplate(key)
		if err != nil {
			c.AbortWithError(500, err)
		}
		templates = append(templates, *template)

	}

	c.JSON(200, templates)
}

func getTemplate(key string) (*Template, error) {
	value, err := database.REDIS.Get(key).Result()
	if err != nil {
		return nil, err
	}

	template := &Template{}

	err = json.Unmarshal([]byte(value), template)
	if err != nil {
		return nil, err
	}

	return template, nil
}

func saveTemplate(template Template) error {
	templateJSON, err := json.Marshal(template)

	if err != nil {
		return err
	}

	return database.REDIS.Set(fmt.Sprintf(templatePath, template.ID), templateJSON, 0).Err()
}
