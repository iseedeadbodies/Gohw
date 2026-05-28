package main

import "fmt"

type Course struct {
	Id    int
	Title string
	Price int
}

type Student struct {
	Name    string
	Courses []Course
}

var catalog = make(map[int]Course)
var students = make(map[string]Student)

func AddCourse(c Course) {
	catalog[c.Id] = c
}

func AddStudent(name string) {
	students[name] = Student{
		Name:    name,
		Courses: []Course{},
	}
}

func Enroll(name string, courseId int) {

	student, ok1 := students[name]
	course, ok2 := catalog[courseId]

	if !ok1 || !ok2 {
		return
	}

	student.Courses = append(student.Courses, course)

	students[name] = student
}

func TotalCost(name string) int {

	student, ok := students[name]

	if !ok {
		return 0
	}

	total := 0

	for _, course := range student.Courses {
		total += course.Price
	}

	return total
}

func PrintStudent(name string) {

	student, ok := students[name]

	if !ok {
		return
	}

	fmt.Println("Student:", student.Name)
	fmt.Println("Courses:")

	for _, course := range student.Courses {
		fmt.Println(course.Title, "-", course.Price)
	}

	fmt.Println("Total:", TotalCost(name))
	fmt.Println()
}

func main() {

	AddCourse(Course{
		Id:    1,
		Title: "Go Basics",
		Price: 20000,
	})

	AddCourse(Course{
		Id:    2,
		Title: "SQL",
		Price: 15000,
	})

	AddCourse(Course{
		Id:    3,
		Title: "Java",
		Price: 25000,
	})

	AddStudent("Ali")
	AddStudent("Dias")

	Enroll("Ali", 1)
	Enroll("Ali", 2)

	Enroll("Dias", 2)
	Enroll("Dias", 3)

	PrintStudent("Ali")
	PrintStudent("Dias")
}
