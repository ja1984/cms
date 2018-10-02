package routes

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID         int    `json:"id"`
	ExternalID string `json:"external_id"`
	Email      string `json:"email"`
}

func RegisterUser(c *gin.Context) {

	//token := c.MustGet("token").(*auth.Token)

	// user := User{
	// 	ExternalID: token.UID,
	// 	Email:      ,
	// }

	// database.SQL.
	c.Status(201)
}
