package routes

import (
	// "github.com/gin-gonic"
	"github.com/gin-gonic/gin"
)

func setupRoutes() {
	router := gin.Default()

	// create a project
	router.POST("/projects", createProjectHandler)

	// get all projects
	router.GET("/projects", getProjectsHandler)

	// get a project by ID
	router.GET("/projects/:id", getProjectByIDHandler)

	// update project by ID
	router.PUT("/projects/:id", updateProjectHandler)

	// delete project
	router.DELETE("/projects/:id", deleteProjectHandler)

	// ...others

	router.Run(":8080")

}
