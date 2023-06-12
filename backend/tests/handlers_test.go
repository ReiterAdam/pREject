package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ReiterAdam/pREject/backend/handlers"
	"github.com/ReiterAdam/pREject/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProjectHandler(t *testing.T) {

	// create new gin router
	router := gin.Default()
	// set up route w =ith handler
	router.POST("/projects", handlers.CreateProjectHandler)

	// create a sample

	project := models.Project{
		// ID:               0,
		Name:             "Test project",
		Description:      "This is a sample of test project",
		Author:           "xxxx",
		CustomProperties: []models.Property{{Key: "Property 1", Value: "Value 1"}, {Key: "Property 2", Value: "Value 2"}},
	}

	// convert data to JSON
	payload, _ := json.Marshal(project)

	// create test http request
	req, _ := http.NewRequest("POST", "/projects", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// create a test http response recorder
	res := httptest.NewRecorder()

	// perform http post request
	router.ServeHTTP(res, req)

	// check response satatus code
	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)
	}
}

func TestGetProjectsHandler(t *testing.T) {

	// create new gin router
	router := gin.Default()
	// set up handler
	router.GET("/projects", handlers.GetProjectsHandler)

	req, _ := http.NewRequest("GET", "/projects", nil)

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)
	}
}

func TestGetProjectByIDHandler(t *testing.T) {

	// create new gin router
	router := gin.Default()
	// set up handler
	router.GET("/projects/:id", handlers.GetProjectByIDHandler)

	// create request, response

	req, _ := http.NewRequest("GET", "/projects/2", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)

	}
}

func TestUpdateProjectHandler(t *testing.T) {

	// create new gin router
	router := gin.Default()
	// set up handler
	router.PUT("/projects/:id", handlers.UpdateProjectHandler)

	project := models.Project{
		ID:          3,
		Name:        "Updated project name",
		Description: "Updated project description",
		CustomProperties: []models.Property{
			{Key: "property1", Value: "updated-value1"},
			{Key: "property2", Value: "Updated-value2"},
		},
	}

	payload, _ := json.Marshal(project)

	// create request, response

	req, _ := http.NewRequest("PUT", "/projects/2", bytes.NewBuffer(payload))
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	// if res.Code != http.StatusOK {
	// 	t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)

	// }

	// assert response status code
	assert.Equal(t, http.StatusOK, res.Code)

	// decode response body
	var response struct {
		Message string `json:"message"`
	}

	err := json.Unmarshal(res.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// assert the response message
	expectedMessage := "Project updated successfully"
	assert.Equal(t, expectedMessage, response.Message)
}
