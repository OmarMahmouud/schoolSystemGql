package main

import (
	"github.com/graphql-go/graphql"
)

var mutationType = graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createStudent":         createStudentMutation,
		"updateStudent":         updateStudentMutation,
		"deleteStudent":         deleteStudentMutation,
		"createTeacher":         createTeacherMutation,
		"updateTeacher":         updateTeacherMutation,
		"deleteTeacher":         deleteTeacherMutation,
		"createCourse":          createCourseMutation,
		"updateCourse":          updateCourseMutation,
		"deleteCourse":          deleteCourseMutation,
		"createCategory":        createCategoryMutation,
		"updateCategory":        updateCategoryMutation,
		"deleteCategory":        deleteCategoryMutation,
		"createOnLineMaterial":  createOnLineMaterialMutation,
		"updateOnLineMaterial":  updateOnLineMaterialMutation,
		"deleteOnLineMaterial":  deleteOnLineMaterialMutation,
		"createOffLineMaterial": createOffLineMaterialMutation,
		"updateOffLineMaterial": updateOffLineMaterialMutation,
		"deleteOffLineMaterial": deleteOffLineMaterialMutation,
		"createStudentCourse":   createStudentCourseMutation,
		"updateStudentCourse":   updateStudentCourseMutation,
		"deleteStudentCourse":   deleteStudentCourseMutation,
	},
}

var createStudentCourseMutation = &graphql.Field{
	Type: StudentCourseType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: createStudentCourseInput,
		},
	},
	Resolve: addStudentCourse,
}

var updateStudentCourseMutation = &graphql.Field{
	Type: StudentCourseType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: updateStudentCourseInput,
		},
	},
	Resolve: updateStudentCourse,
}
var createCourseMutation = &graphql.Field{
	Type: CourseType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: createCourseInput,
		},
	},
	Resolve: createCourse,
}

var updateCourseMutation = &graphql.Field{
	Type: CourseType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: updateCourseInput,
		},
	},
	Resolve: updateCourse,
}

var deleteCourseMutation = &graphql.Field{
	Type:        deleted,
	Description: "delete course by id",
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(deleteInput),
		},
	},
	Resolve: deleteCourse,
}

var deleteCategoryMutation = &graphql.Field{
	Type:        deleted,
	Description: "delete category by id",
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(deleteInput),
		},
	},
	Resolve: deleteCategory,
}

var createCategoryMutation = &graphql.Field{
	Type: CategoryType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: createCategoryInput,
		},
	},
	Resolve: createCategory,
}

var updateCategoryMutation = &graphql.Field{
	Type: CategoryType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: updateCategoryInput,
		},
	},
	Resolve: updateCategory,
}

var createStudentMutation = &graphql.Field{
	Type: StudentType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: createStudentInput,
		},
	},
	Resolve: createStudent,
}

var createTeacherMutation = &graphql.Field{
	Type: TeacherType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: createTeacherInput,
		},
	},
	Resolve: createTeacher,
}

var updateStudentMutation = &graphql.Field{
	Type:        StudentType,
	Description: "Update Student by id",
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: updateStudentInput,
		},
	},
	Resolve: updateStudent,
}

var updateTeacherMutation = &graphql.Field{
	Type:        TeacherType,
	Description: "Update Teacher by id",
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: updateTeacherInput,
		},
	},
	Resolve: updateTeacher,
}

var deleteStudentMutation = &graphql.Field{
	Type:        deleted,
	Description: "delete Student by id",
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(deleteInput),
		},
	},
	Resolve: deleteStudent,
}

var deleteTeacherMutation = &graphql.Field{
	Type:        deleted,
	Description: "delete Student by id",
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(deleteInput),
		},
	},
	Resolve: deleteTeacher,
}

var GetAllInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "getAll",
		Fields: graphql.InputObjectConfigFieldMap{
			"limit": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	},
)

var createCategoryInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CreateCategoryInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var updateCategoryInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "updateCategoryInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

var createCourseInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CreateCourseInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"category_id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	},
)

var createStudentCourseInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CreateStudentCourseInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"studentID": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"courseID": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			}},
	},
)

var updateStudentCourseInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "updateStudentCourseInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"studentID": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"courseID": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			}},
	},
)

var OnLineMaterialInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CreateOnLineMaterialInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"website": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"price": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"bookName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"courseID": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			}},
	},
)

var OffLineMaterialInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CreateOffLineMaterialInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"library": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"price": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"bookName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"courseID": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			}},
	},
)

var updateCourseInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "updateCourseInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"category_id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	},
)

var createStudentInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "createStudentInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"age": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var createTeacherInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "createTeacherInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"age": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"course_id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	},
)

var updateStudentInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "updateStudentInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"age": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

var updateTeacherInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "updateTeacherInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"age": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"course_id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	},
)

var deleteInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "deleteInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
	},
)

var deleteStudentCourseMutation = &graphql.Field{
	Type:        deleted,
	Description: "delete StudentCourse By ID",
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(deleteInput),
		},
	},
	Resolve: deleteStudentCourse,
}

var createOnLineMaterialMutation = &graphql.Field{
	Type: onLineMaterialType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: OnLineMaterialInput,
		},
	},
	Resolve: addOnLineMaterial,
}

var updateOnLineMaterialMutation = &graphql.Field{
	Type: onLineMaterialType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: OnLineMaterialInput,
		},
	},
	Resolve: updateOnLineMaterial,
}

var deleteOnLineMaterialMutation = &graphql.Field{
	Type:        deleted,
	Description: "delete OnLineMaterial by CourseID",
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(deleteInput),
		},
	},
	Resolve: deleteOnLineMaterial,
}

var createOffLineMaterialMutation = &graphql.Field{
	Type: offLineMaterialType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: OffLineMaterialInput,
		},
	},
	Resolve: addOffLineMaterial,
}

var updateOffLineMaterialMutation = &graphql.Field{
	Type: offLineMaterialType,
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: OffLineMaterialInput,
		},
	},
	Resolve: updateOffLineMaterial,
}

var deleteOffLineMaterialMutation = &graphql.Field{
	Type:        deleted,
	Description: "delete OffLineMaterial by CourseID",
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(deleteInput),
		},
	},
	Resolve: deleteOffLineMaterial,
}
