package structs

import "fmt"

type PlatziCourse struct {
	Name   string
	Slug   string
	Skills []string
}

type PlatziCareer struct {
	PlatziCourse
	Courses []PlatziCourse
}

func (p PlatziCourse) Subscribe(name string) {
	fmt.Printf("La persona %s se acaba de subscribir al curso %s\n", name, p.Name)
}
