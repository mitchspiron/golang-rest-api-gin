package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		// Event routes
		v1.GET("/events", app.getAllEvents)
		v1.GET("/events/:id", app.getEvent)

		// User registration route
		v1.POST("/auth/register", app.registerUser)

		// Attendee routes
		v1.GET("/attendees/:id/events", app.getEventsByAttendee)

		// Login route
		v1.POST("/auth/login", app.login)
	}

	// Protected routes
	authGroup := v1.Group("/")
	authGroup.Use(app.authMiddleware())
	{
		// Event routes
		authGroup.POST("/events", app.createEvent)
		authGroup.PUT("/events/:id", app.updateEvent)
		authGroup.DELETE("/events/:id", app.deleteEvent)
		authGroup.POST("/events/:id/attendees/:userId", app.addAttendeeToEvent)

		// Attendee routes
		authGroup.DELETE("/events/:id/attendees/:userId", app.deleteAttendeeFromEvent)
	}

	// Swagger documentation route
	// Redirect /swagger to /swagger/index.html
	// This ensures that accessing /swagger will show the Swagger UI
	g.GET("/swagger/*any", func(c *gin.Context) {
		if c.Request.RequestURI == "/swagger/" {
			c.Redirect(302, "/swagger/index.html")
			return
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	})

	return g
}
