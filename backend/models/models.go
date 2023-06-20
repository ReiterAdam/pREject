package models

type Project struct {
	ID               int
	Name             string
	Description      string
	Author           []User
	LastModifiedBy   []User
	CreatedOn        string
	ModifiedOn       string
	CustomProperties []Property
}

type Position struct {
	ID               int
	Title            string
	Project          []Project
	Author           []User
	LastModifiedBy   []User
	CustomProperties []Property
}

type Property struct {
	Key   string
	Value string
}

type User struct {
	ID       int
	Email    string
	Password string
	Role     string
}
