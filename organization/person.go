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

type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "United States of America"
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

func NewEuropeanUnionIdentifier(id, country string) Citizen {
	return europeanUnionIdentifier{
		id:      id,
		country: country,
	}
}

func (eu europeanUnionIdentifier) ID() string {
	return eu.id
}

func (eu europeanUnionIdentifier) Country() string {
	return eu.country
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
	Citizen
}

func NewPerson(firstname, lastname string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstname,
			last:  lastname,
		},
		Citizen: citizen,
	}
}

func (e Employee) Print() {
	fmt.Println(e.Name.first)
}

func (p *Person) ID() string {
	return fmt.Sprintf("Person's identifier; %s", p.Citizen.ID())
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
