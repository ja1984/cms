package routes

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/ja1984/cms/backend/database"
)

func TestSaveCollection(t *testing.T) {
	database.REDIS = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	testCollection := Collection{
		Key: "test-save",
		Properties: []Property{
			Property{
				Key: "prop1",
				Value: map[string]string{
					"en": "cool",
				},
			},
			Property{
				Key: "prop2withoutLanguage",
				Value: map[string]string{
					"default": "cool",
				},
			},
			Property{
				Key: "prop3Languages",
				Value: map[string]string{
					"en": "horse",
					"sv": "häst",
				},
			},
		},
	}

	err := saveCollection(testCollection)

	if err != nil {
		t.Error(err)
	}
}
