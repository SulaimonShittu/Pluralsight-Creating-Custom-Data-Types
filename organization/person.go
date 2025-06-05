package organization

import (
	"errors"
	"fmt"
	"strings"
)

type TwitterHandler string

type Handler struct {
	Handle string
	Name   string
}

func (t TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(t), "@")
	return fmt.Sprintf("https://www.x.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) socialSecurityNumber {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

type Name struct {
	first string
	last  string
}

func (n Name) FullName() string {
	return fmt.Sprintf("Your fullname is : %s %s", n.first, n.last)
}

type Employee struct {
	Name
}

type Person struct {
	Name
	twitterHandler TwitterHandler
	Identifiable
}

func NewPerson(firstname, lastname string, identifiable Identifiable) Person {
	return Person{
		Name: Name{
			first: firstname,
			last:  lastname,
		},
		Identifiable: identifiable,
	}
}

func (e Employee) Print() {
	fmt.Println(e.Name.first)
}

func (p *Person) ID() string {
	return "12345"
}

func (p Person) FullName() string {
	return fmt.Sprintf("Your fullname is : %s %s", p.first, p.last)
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if !strings.HasPrefix(string(handler), "@") && len(handler) > 0 {
		return errors.New("Twitter handler missing important @ symbol.")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) GetTwitterhandler() TwitterHandler {
	return p.twitterHandler
}

func asaa() {

}
