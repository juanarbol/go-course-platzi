package structs

import "fmt"

type Platzi interface {
	Subscribe(name string)
}

func deferTest() {
	fmt.Println("InterfaceTest has been finished!")
}

func InterfaceTest(name string) {
	defer deferTest()
	course1 := PlatziCourse{
		Name:   "Curso de Go",
		Slug:   "/go",
		Skills: []string{"Hola", "mundo"},
	}

	career := new(PlatziCareer)
	career.Name = "Curso de GO"
	career.Slug = "/golang"

	callSubsCribe(course1, name)
	callSubsCribe(career, name)
}

func callSubsCribe(p Platzi, name string) {
	p.Subscribe(name)
}
