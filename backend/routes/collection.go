package routes

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ja1984/cms/backend/database"
)

const collectionPath = "collections:%s"

func CreateCollection(c *gin.Context) {

	collection := Collection{}

	c.BindJSON(&collection)

	err := saveCollection(collection)

	if err != nil {
		c.AbortWithError(500, err)
	}

	c.Status(201)
}

func ListCollections(c *gin.Context) {

	var cursor uint64
	keysToGet := []string{}
	for {
		var keys []string
		var err error
		keys, cursor, err = database.REDIS.Scan(0, "collections:*", 0).Result()
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

	collections := []Collection{}

	for _, key := range keysToGet {
		collection, err := getCollection(key)
		if err != nil {
			c.AbortWithError(500, err)
		}
		collections = append(collections, *collection)

	}

	c.JSON(200, collections)
}

func GetCollection(c *gin.Context) {

	key := c.Param("key")

	collection, err := getCollection(fmt.Sprintf(collectionPath, key))
	if err != nil {
		c.AbortWithError(500, err)
	}

	c.JSON(200, collection)
}

func getCollection(key string) (*Collection, error) {
	value, err := database.REDIS.Get(key).Result()
	if err != nil {
		return nil, err
	}

	collection := &Collection{}

	err = json.Unmarshal([]byte(value), collection)
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func saveCollection(collection Collection) error {
	collectionJson, err := json.Marshal(collection)

	if err != nil {
		return err
	}

	return database.REDIS.Set(fmt.Sprintf(collectionPath, collection.Key), collectionJson, 0).Err()
}

type Collection struct {
	Key        string     `json:"key"`
	Properties []Property `json:"properties"`
}

type Property struct {
	Key   string            `json:"key"`
	Value map[string]string `json:"value"`
}
