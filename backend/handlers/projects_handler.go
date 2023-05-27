package handlers

import (
	"log"
	"net/http"

	"github.com/ReiterAdam/pREject/backend/database"
	"github.com/ReiterAdam/pREject/backend/models"
	"github.com/gin-gonic/gin"
)

func CreateProjectHandler(c *gin.Context) {
	// Parse JSON request body

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// save project to database
	db := database.SetupDB()
	defer db.Close()
	res := database.CheckDB(db)

	log.Printf("Result of databaseCheck is: %v", res)
	// Insert the project to the database

	// prepare  query.
	stmt, err := db.Prepare("INSERT INTO projects (name, description) VALUES (?, ?)")
	if err != nil {
		log.Print("Statement preparation failed!")
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer stmt.Close()
	// execute statement
	result, err := stmt.Exec(project.Name, project.Description)

	if err != nil {
		log.Print("Statement execution failed!")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get ID of the freshly added project
	projectID, _ := result.LastInsertId()

	// add additional properties of the project
	for _, prop := range project.CustomProperties {
		stmt, err = db.Prepare("INSERT INTO project_properties (project_id, property_key, property_value) VALUES (?, ?, ?)")
		if err != nil {
			log.Print("Statement preparation failed!")
			log.Print(err)
			// log.Printf("Values: \nProjectID: %v\nProperty: %v\nValue: %v", projectID, prop.Key, prop.Value)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		_, err = stmt.Exec(projectID, prop.Key, prop.Value)
		if err != nil {
			log.Printf("Statement execution failed! Values: \nProjectID: %v\nProperty: %v\nValue: %v", projectID, prop.Key, prop.Value)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project created succesfully"})

}

func getProjectsHandler(c *gin.Context) {
	// logic here to get all projects
}

func getProjectByIDHandler(c *gin.Context) {
	// logic to retrieve a project by ID
}

func updateProjectHandler(c *gin.Context) {
	// logic to update a project
}

func deleteProjectHandler(c *gin.Context) {
	// logic to delete a project
}

//  other handlers for positions, etc.
// ...
