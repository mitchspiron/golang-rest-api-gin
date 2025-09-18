package main

import (
	"golang-rest-api-gin/internal/database"

	"github.com/gin-gonic/gin"
)

func (app *application) GetUserFromContext(c *gin.Context) *database.User {
	// Retrieve the user from the context
	contextUser, exists := c.Get("user")

	// If the user does not exist in the context, return an empty user
	if !exists {
		return &database.User{}
	}

	// Type assert the user to *database.User
	user, ok := contextUser.(*database.User)

	if !ok {
		return &database.User{}
	}

	return user
}
