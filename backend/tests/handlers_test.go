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
