package handlers

import (
	"log"
	"net/http"
	"time"

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

	// perform database check, get schema
	if res {
		log.Print("Database check ...OK")
	} else {
		log.Print("Database check ...ERROR")
	}
	// Insert the project to the database

	// prepare  query.
	stmt, err := db.Prepare("INSERT INTO projects (name, description, author, CreatedOn, ModifiedOn) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Print("Statement preparation failed!")
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	currentTime := getCurrentTime()
	defer stmt.Close()
	// execute statement
	result, err := stmt.Exec(project.Name, project.Description, project.Author, currentTime, currentTime)

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

func GetProjectsHandler(c *gin.Context) {
	// get all projects form database

	db := database.SetupDB()
	defer db.Close()

	rows, err := db.Query("SELECT ID, Name, Description, Author, CreatedOn, ModifiedOn FROM projects")
	if err != nil {
		log.Print("Query execution failed!")
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	projects := []models.Project{}
	for rows.Next() {
		var project models.Project
		err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.Author, &project.CreatedOn, &project.ModifiedOn)
		if err != nil {
			log.Print("Scan failed!")
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// get custom properties of each project
		propRows, err := db.Query("SELECT property_key, property_value FROM project_properties WHERE project_id = ?", project.ID)
		if err != nil {
			log.Print("Properties scan failed!")
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer propRows.Close()

		for propRows.Next() {
			var key, value string
			err := propRows.Scan(&key, &value)
			if err != nil {
				log.Print("Key, value scan failed!")
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			project.CustomProperties = append(project.CustomProperties, models.Property{Key: key, Value: value})
		}

		projects = append(projects, project)
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})

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

func getCurrentTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02-15-04")
	return formattedTime
}
