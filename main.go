package main

import (
	"Pluralsight-Creating-Custom-Data-Types/organization"
	"fmt"
)

// Course to be started again

func main() {
	p := organization.NewPerson("Sulaimon", "Shittu", organization.NewSocialSecurityNumber("123-45-6789"))
	err := p.SetTwitterHandler("@eniolorunmife")

	if err != nil {
		fmt.Print(err.Error())
	}

	println(p.GetTwitterhandler())
	println(p.GetTwitterhandler().RedirectUrl())
	println(p.ID())

}
