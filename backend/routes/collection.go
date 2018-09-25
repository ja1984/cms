package routes

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ja1984/cms/backend/database"
)

func CreateCollection(c *gin.Context) {

	collection := Collection{}

	c.BindJSON(&collection)

	err := saveCollection(collection)

	if err != nil {
		c.AbortWithError(500, err)
	}

	c.Status(201)
}

func saveCollection(collection Collection) error {
	collectionJson, err := json.Marshal(collection)

	if err != nil {
		return err
	}

	return database.REDIS.Set(fmt.Sprintf("collections:%s", collection.Key), collectionJson, 0).Err()
}

type Collection struct {
	Key        string     `json:"key"`
	Properties []Property `json:"properties"`
}

type Property struct {
	Key   string            `json:"key"`
	Value map[string]string `json:"value"`
}
