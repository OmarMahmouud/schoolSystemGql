package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

func GetCourses(params graphql.ResolveParams) (interface{}, error) {
	var courses []Course
	input, ok := params.Args["input"].(map[string]interface{})
	limit := 10
	if input["limit"] != nil {
		limit = input["limit"].(int)
	}
	offset := 0
	if input["offset"] != nil {
		offset = input["offset"].(int)
	}
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	DB.Limit(limit).Offset(offset).Preload("StudentCourse").Preload("offLineMaterialType").Preload("onLineMaterialType").Find(&courses)
	return courses, nil
}

func GetStudents(params graphql.ResolveParams) (interface{}, error) {
	var students []Student
	input, ok := params.Args["input"].(map[string]interface{})
	limit := 10
	if input["limit"] != nil {
		limit = input["limit"].(int)
	}
	offset := 0
	if input["offset"] != nil {
		offset = input["offset"].(int)
	}
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	DB.Limit(limit).Offset(offset).Find(&students)
	return students, nil
}

func GetTeachers(params graphql.ResolveParams) (interface{}, error) {
	var teachers []Teacher
	input, ok := params.Args["input"].(map[string]interface{})
	limit := 10
	if input["limit"] != nil {
		limit = input["limit"].(int)
	}
	offset := 0
	if input["offset"] != nil {
		offset = input["offset"].(int)
	}
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	DB.Limit(limit).Offset(offset).Find(&teachers)
	return teachers, nil
}

func GetStudentCourse(params graphql.ResolveParams) (interface{}, error) {
	var studentCourse []StudentCourse
	input, ok := params.Args["input"].(map[string]interface{})
	limit := 10
	if input["limit"] != nil {
		limit = input["limit"].(int)
	}
	offset := 0
	if input["offset"] != nil {
		offset = input["offset"].(int)
	}
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	DB.Limit(limit).Offset(offset).Find(&studentCourse)
	return studentCourse, nil
}
func GetCategories(params graphql.ResolveParams) (interface{}, error) {
	var categories []Category
	input, ok := params.Args["input"].(map[string]interface{})
	limit := 10
	if input["limit"] != nil {
		limit = input["limit"].(int)
	}
	offset := 0
	if input["offset"] != nil {
		offset = input["offset"].(int)
	}
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	DB.Limit(limit).Offset(offset).Preload("Courses").Find(&categories)

	return categories, nil
}

func GetCourse(p graphql.ResolveParams) (interface{}, error) {
	var course Course
	id, _ := p.Args["id"].(int)
	DB.Where("id = ? ", id).Find(&course)
	return course, nil
}

func createStudent(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	student := Student{
		Name:  input["name"].(string),
		Age:   input["age"].(int),
		Email: input["email"].(string),
	}
	DB.Create(&student)
	return student, nil
}

func createTeacher(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	teacher := Teacher{
		Name:     input["name"].(string),
		Age:      input["age"].(int),
		Email:    input["email"].(string),
		CourseID: uint(input["course_id"].(int)),
	}
	if input["course_id"] != nil {
		teacher.CourseID = uint(input["course_id"].(int))
	}
	DB.Create(&teacher)
	return teacher, nil
}

func createCourse(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	course := Course{
		Name:        input["name"].(string),
		Description: input["description"].(string),
	}
	if input["category_id"] != nil {
		course.CategoryID = uint(input["category_id"].(int))
	}
	DB.Create(&course)
	return course, nil
}

func addStudentCourse(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	studentCourse := StudentCourse{
		CourseID:  uint(input["courseID"].(int)),
		StudentID: uint(input["studentID"].(int)),
	}
	DB.Where("student_id = ? AND course_id = ?", uint(input["studentID"].(int)), uint(input["courseID"].(int))).Unscoped().First(&studentCourse)
	if studentCourse.ID != 0 {
		return nil, nil
	}
	DB.Create(&studentCourse)
	return studentCourse, nil
}

func updateStudentCourse(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	var studentCourse StudentCourse
	studentCourse.ID = uint(input["id"].(int))
	DB.Find(&studentCourse)
	if input["course_id"] != nil {
		studentCourse.CourseID = uint(input["course_id"].(int))
	}
	if input["student_id"] != nil {
		studentCourse.StudentID = uint(input["student_id"].(int))
	}
	DB.Updates(&studentCourse)
	return studentCourse, nil
}

func addOnLineMaterial(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	onLineMaterial := OnLineMaterial{
		CourseID: uint(input["courseID"].(int)),
		Website:  input["website"].(string),
		BookName: input["bookName"].(string),
	}
	if input["price"] != nil {
		onLineMaterial.Price = uint(input["price"].(int))
	}
	DB.Create(&onLineMaterial)
	return onLineMaterial, nil
}

func addOffLineMaterial(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	offLineMaterial := OffLineMaterial{
		CourseID: uint(input["courseID"].(int)),
		Library:  input["library"].(string),
		BookName: input["bookName"].(string),
	}
	if input["price"] != nil {
		offLineMaterial.Price = uint(input["price"].(int))
	}
	DB.Create(&offLineMaterial)
	return offLineMaterial, nil
}

func updateOnLineMaterial(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	onLineMaterial := OnLineMaterial{
		CourseID: uint(input["courseID"].(int)),
		Website:  input["website"].(string),
		BookName: input["bookName"].(string),
	}
	if input["price"] != nil {
		onLineMaterial.Price = uint(input["price"].(int))
	}
	DB.Where("course_id = ? ", uint(input["courseID"].(int))).Updates(&onLineMaterial)
	return onLineMaterial, nil
}

func updateOffLineMaterial(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	offLineMaterial := OffLineMaterial{
		CourseID: uint(input["courseID"].(int)),
		Library:  input["library"].(string),
		BookName: input["bookName"].(string),
	}
	if input["price"] != nil {
		offLineMaterial.Price = uint(input["price"].(int))
	}
	DB.Where("course_id = ? ", uint(input["courseID"].(int))).Updates(&offLineMaterial)
	return offLineMaterial, nil
}

func updateCourse(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	var course Course
	course.ID = uint(input["id"].(int))
	if input["name"] != nil {
		course.Name = input["name"].(string)
	}
	if input["description"] != nil {
		course.Description = input["description"].(string)
	}
	if input["category_id"] != nil {
		fmt.Println("error")
		course.CategoryID = uint(input["category_id"].(int))
	}
	DB.Where("id = ?", course.ID).Updates(&course)
	DB.Preload("Category").Find(&course)
	return course, nil
}

func createCategory(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	category := Category{
		Name:        input["name"].(string),
		Description: input["description"].(string),
	}
	DB.Create(&category)
	return category, nil
}

func updateCategory(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}

	var category Category
	category.ID = uint(input["id"].(int))
	DB.Find(&category)
	if input["name"] != nil {
		category.Name = input["name"].(string)
	}
	if input["description"] != nil {
		category.Description = input["description"].(string)
	}
	DB.Updates(&category)
	return category, nil
}

func updateStudent(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	var student Student
	student.ID = uint(input["id"].(int))
	DB.Find(&student)
	if input["name"] != nil {
		student.Name = input["name"].(string)
	}
	if input["age"] != nil {
		student.Age = input["age"].(int)
	}
	if input["email"] != nil {
		student.Email = input["email"].(string)
	}
	DB.Updates(&student)
	return student, nil
}

func updateTeacher(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	var teacher Teacher
	teacher.ID = uint(input["id"].(int))
	DB.Find(&teacher)
	if input["name"] != nil {
		teacher.Name = input["name"].(string)
	}
	if input["age"] != nil {
		teacher.Age = input["age"].(int)
	}
	if input["email"] != nil {
		teacher.Email = input["email"].(string)
	}
	if input["course_id"] != nil {
		teacher.CourseID = input["course_id"].(uint)
	}
	DB.Updates(&teacher)
	return teacher, nil
}

func deleteStudent(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	DB.Where("id = ?", input["id"]).Delete(&Student{})
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}

func deleteTeacher(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	DB.Where("id = ?", input["id"]).Delete(&Teacher{})
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}

func deleteCourse(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	DB.Where("id = ?", input["id"]).Delete(&Course{})
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}

func deleteStudentCourse(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	DB.Where("id = ?", input["id"]).Delete(&StudentCourse{})
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}

func deleteOnLineMaterial(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	DB.Where("id = ?", input["id"]).Unscoped().Delete(&OnLineMaterial{})
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}

func deleteOffLineMaterial(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	DB.Where("id = ?", input["id"]).Unscoped().Delete(&OffLineMaterial{})
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}

func deleteCategory(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input type")
	}
	DB.Where("id = ?", input["id"]).Delete(&Category{})
	isDeleted := isDeleted{
		deleted: "yes",
	}
	return isDeleted, nil
}
