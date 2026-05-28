package main

import "fmt"

type Worker interface {
	Work() string
	GetName() string
}

type Programmer struct {
	Name     string
	Language string
}

type Designer struct {
	Name string
	Tool string
}

func (p Programmer) Work() string {
	return "Программист " + p.Name + " пишет код на " + p.Language
}

func (p Programmer) GetName() string {
	return p.Name
}

func (d Designer) Work() string {
	return "Дизайнер " + d.Name + " делает макет в " + d.Tool
}

func (d Designer) GetName() string {
	return d.Name
}

func ShowWork(w Worker) {
	fmt.Println("Name:", w.GetName())
	fmt.Println(w.Work())
	fmt.Println()
}

func main() {
	p1 := Programmer{Name: "Ali", Language: "Go"}
	p2 := Programmer{Name: "Dias", Language: "Java"}

	d1 := Designer{Name: "Alina", Tool: "Figma"}
	d2 := Designer{Name: "Amina", Tool: "Photoshop"}

	ShowWork(p1)
	ShowWork(p2)
	ShowWork(d1)
	ShowWork(d2)
}
