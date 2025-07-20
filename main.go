package main

import (
	"Pluralsight-Creating-Custom-Data-Types/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Sulaimon", "Shittu", organization.NewSocialSecurityNumber("123-45-6789"))
	err := p.SetTwitterHandler("@eniolorunmife")

	if err != nil {
		fmt.Print(err.Error())
	}
}

type Name struct {
	first string
	last  string
}
