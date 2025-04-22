package routes

import (
	"module-name/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes(server *gin.Engine) {
	server.GET("/get-events", getEvents)
	server.GET("/get-event/:id", getEventById)
	server.POST("/signup", saveUser)
	server.POST("/login", checkUser)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.AuthMiddleware)
	authenticated.POST("/add-event", addEvent)
	authenticated.PUT("/update-event/:id", updateEventById)
	authenticated.DELETE("/delete-event/:id", deleteEventById)
	authenticated.POST("/events/:id/register", registerUserForEvent)
	authenticated.DELETE("/events/:id/unregister", unregisterUserForEvent)
}
