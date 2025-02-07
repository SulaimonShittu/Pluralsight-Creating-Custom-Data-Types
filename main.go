package main

import (
	"Pluralsight-Creating-Custom-Data-Types/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Sula", "Shittu")
	err := p.SetTwitterHandler("@eniolorunmife")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting the Twitter handler: %s\n", err.Error())
	}
	fmt.Println(p.GetTwitterHandler())
	fmt.Println(p.GetTwitterHandler().RedirectUrl())
	fmt.Println(p.ID())
	fmt.Println(p.FullName())
}
