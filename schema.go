package main

import "github.com/graphql-go/graphql"

// Schema
var fields = graphql.Fields{

	"courses": &graphql.Field{
		Type: graphql.NewList(CourseType),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: GetAllInput,
			},
		},
		Resolve: GetCourses,
	},
	"students": &graphql.Field{
		Type: graphql.NewList(StudentType),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: GetAllInput,
			},
		},
		Resolve: GetStudents,
	},

	"teachers": &graphql.Field{
		Type: graphql.NewList(TeacherType),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: GetAllInput,
			},
		},
		Resolve: GetTeachers,
	},

	"studentsCourses": &graphql.Field{
		Type: graphql.NewList(StudentCourseType),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: GetAllInput,
			},
		},
		Resolve: GetStudentCourse,
	},

	"course": &graphql.Field{
		Type: CourseType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: GetCourse,
	},

	"categories": &graphql.Field{
		Type: graphql.NewList(CategoryType),
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: GetAllInput,
			},
		},
		Resolve: GetCategories,
	},
}

var StudentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Student",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(Student).ID, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var TeacherType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Teacher",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(Teacher).ID, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.Int,
		},
		"courseID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(Teacher).CourseID, nil
			}},
	},
})

var deleted = graphql.NewObject(graphql.ObjectConfig{
	Name: "isDeleted",
	Fields: graphql.Fields{
		"deleted": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(isDeleted).deleted, nil
			},
		},
	},
})

var CourseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Course",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(Course).ID, nil
			},
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"start_date": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(Course).StartDate, nil
			}},
		"end_date": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(Course).EndDate, nil
			}},
		"category_id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(Course).CategoryID, nil
			}},
		"studentCourse": &graphql.Field{
			Type: graphql.NewList(StudentCourseType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(Course).StudentCourse, nil
			},
		},
		"on_line_material": &graphql.Field{
			Type: onLineMaterialType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(Course).OnLineMaterial, nil
			}},
		"off_line_material": &graphql.Field{
			Type: offLineMaterialType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(Course).OffLineMaterial, nil
			}},
	},
})

var StudentCourseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "studentCourse",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(StudentCourse).ID, nil
			},
		},
		"course_id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(StudentCourse).CourseID, nil
			},
		},
		"student_id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(StudentCourse).StudentID, nil
			},
		},
	},
})

var CategoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Category",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(Category).ID, nil
			},
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"courses": &graphql.Field{
			Type: graphql.NewList(CourseType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(Category).Courses, nil
			},
		},
	},
})

var onLineMaterialType = graphql.NewObject(graphql.ObjectConfig{
	Name: "onLineMaterial",
	Fields: graphql.Fields{
		"courseID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(OnLineMaterial).CourseID, nil
			},
		},
		"website": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(OnLineMaterial).Website, nil
			},
		},
		"bookName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(OnLineMaterial).BookName, nil
			},
		},
		"price": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(OnLineMaterial).Price, nil
			},
		},
	},
})

var offLineMaterialType = graphql.NewObject(graphql.ObjectConfig{
	Name: "offLineMaterial",
	Fields: graphql.Fields{
		"courseID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(OffLineMaterial).CourseID, nil
			},
		},
		"library": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(OffLineMaterial).Library, nil
			},
		},
		"bookName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(OffLineMaterial).BookName, nil
			},
		},
		"price": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(OffLineMaterial).Price, nil
			},
		},
	},
})
