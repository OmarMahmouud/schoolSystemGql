package main

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	gorm.Model
	Name            string          `json:"name" binding:"required"  gorm:"unique"`
	Description     string          `json:"description" binding:"required"`
	StartDate       time.Time       `json:"start_date"`
	EndDate         time.Time       `json:"end_date"`
	CategoryID      uint            `json:"category_id" validate:"required"`
	Category        Category        `json:"category"`
	StudentCourse   []StudentCourse `json:"student_course"`
	OnLineMaterial  OnLineMaterial
	OffLineMaterial OffLineMaterial
}

type Category struct {
	gorm.Model
	Name        string   `json:"name" validate:"required" gorm:"type:varchar(50);unique"`
	Description string   `json:"description" validate:"required"  gorm:"type:varchar(500)"`
	Courses     []Course `json:"courses"`
}

type Student struct {
	gorm.Model
	Name  string `json:"name" validate:"required"  gorm:"type:varchar(50)"`
	Age   int    `json:"age"`
	Email string `json:"email" validate:"email,required"  gorm:"unique"`
}

type StudentCourse struct {
	gorm.Model
	StudentID uint `json:"student_id"`
	CourseID  uint `json:"course_id"`
}

type isDeleted struct {
	deleted string
}

type Teacher struct {
	gorm.Model
	Name     string `json:"name" validate:"required"  gorm:"type:varchar(50)"`
	Age      int    `json:"age"`
	Email    string `json:"email" validate:"email,required"  gorm:"unique"`
	CourseID uint   `json:"course_id" gorm:"unique"`
}

type OnLineMaterial struct {
	gorm.Model
	Website  string `json:"website" validate:"required" `
	Price    uint   `json:"price"`
	BookName string `json:"book_name" validate:"required" `
	CourseID uint   `json:"course_id" validate:"required" gorm:"unique"`
}

type OffLineMaterial struct {
	gorm.Model
	Library  string `json:"library" validate:"required"`
	BookName string `json:"book_name" validate:"required" `
	Price    uint   `json:"price"`
	CourseID uint   `json:"course_id" validate:"required" gorm:"unique"`
}
