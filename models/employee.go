package models

import ( 
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	// Employeeid int `gorm:"primaryKey;size:255;uniqueIndex:empid_indx"`
	FirstName string   
	LastName string 
	Designation string
	Salary float64
	Country string
}


func (Employee) TableName() string {
    return "anandhu2.employee_new"
}

