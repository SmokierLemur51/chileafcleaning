package data

import (
	// "database/sql"
	// "errors"
	"time"
)

// this file defines data models as well as the logic
// to run CRUD ops on the database 



// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 
// * * * * * * * * * * * * * * Pages * * * * * * * * * * * * * * * * * * * * * 
// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 
type PublicPageData struct {
	Page 		string
	Title		string
	Content 	string
	Endpoints   PublicEndPoints
}

type PublicEndPoints struct {
	Index 		string
	About 		string
	Services    string
}

// client portal
type ClientPageData struct {}


// admin pages
type DirectorPageData struct {}


// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 
// * * * * * * * * * * * * * * Client System * * * * * * * * * * * * * * * * *
// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 
type ContactInfo struct {
	ID int
	CompanyID int
	Name string
	PhoneNumber string
	Email string
}

type Address struct {
	ID int
	Street1 string
	Street2 string
	City string
	State string
	Zip string
}

type Client struct {
	ID   	int
	Contact ContactInfo
	Address Address
}


// func GetClientByID(id int) (*Client, error) {
// 	// database query

// 	if clientNotFound {
// 		return nil, errors.New("Client not found.")
// 	}

// 	return &client, nil	
// }

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 
// * * * * * * * * * * * * * * Project  System * * * * * * * * * * * * * * * * 
// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 
type Status struct {
	ID int
	Name string
	Description string
	Active bool
}

type Project struct {
	ID   				int
	Name    			string
	PrivateDescription  string
	PublicDescription 	string
	Client				Client
	Url   				string
	Status  			Status
	Tickets  			[]Ticket
	Notes				[]Note
	CreatedAt			time.Time
	CompletedAt  	    time.Time
	SellingPrice	    float32
	HourlyRate	    float32
	TotalHoursWorked    float32

}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 
// * * * * * * * * * * * * * * * Ticket System * * * * * * * * * * * * * * * * 
// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * 
type Ticket struct {
	ID   	int
	ProjectID int
	CreatedAt time.Time
	CompleteAt time.time
	Status	Status
	HourlyRate float32
	TotalHours float32
	Title string
	Content string
	Tasks []Task
	Notes []Note
}

type Task struct {
	ID int
	TaskId int
	Complete bool
	Message string
}

type Note struct {
	ID int
	TaskId int
	Title string
	Content string
}
