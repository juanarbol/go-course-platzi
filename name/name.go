package name

import "fmt"

// GetName Es una función muy linda
func GetName() string {
	fmt.Print("Hola, dime tu nombre")

	var name string
	fmt.Scanf("%s", &name)
	return name
}
