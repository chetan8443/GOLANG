package main

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmpName string `json:"empName"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}
