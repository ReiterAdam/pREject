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
	// perform database check, get schema
	if res {
		log.Print("Database check ...OK")
	} else {
		log.Print("Database check ...ERROR")
	}
	// Insert the project to the database

	// prepare  query.
	stmt, err := db.Prepare("INSERT INTO projects (name, description, author, CreatedOn, ModifiedOn) VALUES (?, ?, ?, ?, ?)")

	if checkErr(c, err, "Statement preparation failed!") {
		return
	}
	currentTime := getCurrentTime()
	defer stmt.Close()
	// execute statement
	result, err := stmt.Exec(project.Name, project.Description, project.Author, currentTime, currentTime)
	if checkErr(c, err, "Statement execution failed!") {
		return
	}

	// Get ID of the freshly added project
	projectID, _ := result.LastInsertId()

	// add additional properties of the project
	for _, prop := range project.CustomProperties {
		stmt, err = db.Prepare("INSERT INTO project_properties (project_id, property_key, property_value) VALUES (?, ?, ?)")

		if checkErr(c, err, "Statement preparation failed!") {
			return
		}
		_, err = stmt.Exec(projectID, prop.Key, prop.Value)
		// if err != nil {
		// 	log.Printf("Statement execution failed! Values: \nProjectID: %v\nProperty: %v\nValue: %v", projectID, prop.Key, prop.Value)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
		// }

		if checkErr(c, err, "Statement preparation failed!") {
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project created succesfully"})

}

func GetProjectsHandler(c *gin.Context) {
	// get all projects form database

	db := database.SetupDB()
	defer db.Close()

	rows, err := db.Query("SELECT Name, Description, Author, CreatedOn, ModifiedOn FROM projects")
	if checkErr(c, err, "Query execution failed!") {
		return
	}

	defer rows.Close()

	projects := []models.Project{}
	for rows.Next() {
		var project models.Project
		err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.Author, &project.CreatedOn, &project.ModifiedOn)
		if checkErr(c, err, "Scan execution failed!") {
			return
		}

		// get custom properties of each project
		propRows, err := db.Query("SELECT property_key, property_value FROM project_properties WHERE project_id = ?", project.ID)
		if checkErr(c, err, "Properties execution failed!") {
			return
		}

		defer propRows.Close()
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})

}

func getProjectByIDHandler(c *gin.Context) {
	// get projectID from request
	projectID := c.Param("id")

	db := database.SetupDB()
	defer db.Close()

	var project models.Project
	err := db.QueryRow("SELECT (Name, Description, Author, ModifiedOn, CreatedOn) FROM projects WHERE id = ?", projectID).Scan(&project.ID, &project.Name, &project.Description, &project.Author, &project.ModifiedOn, &project.CreatedOn)
	if checkErr(c, err, "Error while searching for project!") {
		return
	}

	// get custom properties of each project
	propRows, err := db.Query("SELECT property_key, property_value FROM project_properties WHERE project_id = ?", projectID)
	if checkErr(c, err, "Error while searching for project properties!") {
		return
	}
	defer propRows.Close()

	for propRows.Next() {
		var key, value string
		err := propRows.Scan(&key, &value)
		if checkErr(c, err, "Project properties scan failed!") {
			return
		}
		project.CustomProperties = append(project.CustomProperties, models.Property{Key: key, Value: value})
	}

	c.JSON(http.StatusOK, gin.H{"project": project})

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

func checkErr(c *gin.Context, err error, message string) bool {
	if err != nil {
		log.Print(message)
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return true
	}
	return false
}
