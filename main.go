package main

import (
	"Pluralsight-Creating-Custom-Data-Types/organization"
	"fmt"
)

// Course to be started again

func main() {
	p := organization.NewPerson("Sulaimon", "Shittu")
	err := p.SetTwitterHandler("")
	if err != nil {
		fmt.Print(err.Error())
	}

	println(p.ID())
	println(p.FullName())
	println(p.GetTwitterhandler())
}
