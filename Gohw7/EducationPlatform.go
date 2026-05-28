package main

import "fmt"

type Participant interface {
	Info()
	DoAction()
}

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) User {
	u := User{}
	u.SetName(name)
	u.SetAge(age)
	return u
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetAge() int {
	return u.age
}

func (u *User) SetName(name string) {
	if name == "" {
		fmt.Println("Error: name cannot be empty")
		return
	}

	u.name = name
}

func (u *User) SetAge(age int) {
	if age < 0 {
		fmt.Println("Error: age cannot be negative")
		return
	}

	u.age = age
}

type Student struct {
	User
	Score int
}

func NewStudent(name string, age int, score int) *Student {
	if score < 0 {
		score = 0
	}

	return &Student{
		User:  NewUser(name, age),
		Score: score,
	}
}

func (s Student) Info() {
	fmt.Println("Student:", s.GetName())
	fmt.Println("Age:", s.GetAge())
	fmt.Println("Score:", s.Score)
}

func (s Student) DoAction() {
	fmt.Println(s.GetName(), "studies on the platform")
}

type Teacher struct {
	User
	Subject string
}

func NewTeacher(name string, age int, subject string) *Teacher {
	if subject == "" {
		subject = "Unknown"
	}

	return &Teacher{
		User:    NewUser(name, age),
		Subject: subject,
	}
}

func (t Teacher) Info() {
	fmt.Println("Teacher:", t.GetName())
	fmt.Println("Age:", t.GetAge())
	fmt.Println("Subject:", t.Subject)
}

func (t Teacher) DoAction() {
	fmt.Println(t.GetName(), "teaches", t.Subject)
}

func (t Teacher) GradeStudent(student *Student, points int, course Course) {
	if points <= 0 {
		fmt.Println("Error: points must be positive")
		return
	}

	student.Score += points

	if student.Score > course.MaxScore {
		student.Score = course.MaxScore
	}
}

type Course struct {
	Title       string
	MaxScore    int
	MaxStudents int
	Students    []*Student
	Teacher     *Teacher
}

func NewCourse(title string, maxScore int, maxStudents int, teacher *Teacher) Course {
	return Course{
		Title:       title,
		MaxScore:    maxScore,
		MaxStudents: maxStudents,
		Teacher:     teacher,
		Students:    []*Student{},
	}
}

func (c *Course) AddStudent(student *Student) {
	if len(c.Students) >= c.MaxStudents {
		fmt.Println("Cannot add student:", student.GetName())
		fmt.Println("Course is full")
		return
	}

	c.Students = append(c.Students, student)
}

func (c Course) Info() {
	fmt.Println("Course:", c.Title)
	fmt.Println("Max score:", c.MaxScore)
	fmt.Println("Max students:", c.MaxStudents)
	fmt.Println("Teacher:", c.Teacher.GetName())
	fmt.Println("Students:")

	for _, student := range c.Students {
		fmt.Println("-", student.GetName())
	}
}

func AverageScore(students []*Student) float64 {
	if len(students) == 0 {
		return 0
	}

	sum := 0

	for _, student := range students {
		sum += student.Score
	}

	return float64(sum) / float64(len(students))
}

func BestStudent(students []*Student) *Student {
	if len(students) == 0 {
		return nil
	}

	best := students[0]

	for _, student := range students {
		if student.Score > best.Score {
			best = student
		}
	}

	return best
}

func main() {
	teacher := NewTeacher("Aibek", 35, "Go")

	course := NewCourse("Go Programming", 100, 3, teacher)

	student1 := NewStudent("Ali", 18, 40)
	student2 := NewStudent("Dias", 19, 60)
	student3 := NewStudent("Amina", 18, 80)
	student4 := NewStudent("Sanzhar", 20, 50)

	course.AddStudent(student1)
	course.AddStudent(student2)
	course.AddStudent(student3)
	course.AddStudent(student4)

	fmt.Println("COURSE INFO")
	course.Info()

	fmt.Println("\nPARTICIPANTS INFO")

	participants := []Participant{teacher, student1, student2, student3}

	for _, p := range participants {
		p.Info()
		p.DoAction()
		fmt.Println()
	}

	fmt.Println("GRADING STUDENTS")

	teacher.GradeStudent(student1, 30, course)
	teacher.GradeStudent(student2, 50, course)
	teacher.GradeStudent(student3, 10, course)

	student1.Info()
	fmt.Println()

	student2.Info()
	fmt.Println()

	student3.Info()
	fmt.Println()

	fmt.Println("Average score:", AverageScore(course.Students))

	best := BestStudent(course.Students)

	if best != nil {
		fmt.Println("Best student:", best.GetName())
		fmt.Println("Score:", best.Score)
	}
}
