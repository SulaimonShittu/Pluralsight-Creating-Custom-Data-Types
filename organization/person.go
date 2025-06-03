package organization

import (
	"errors"
	"fmt"
	"strings"
)

type TwitterHandler Handler

type Handler struct {
	Handle string
	Name   string
}

func (t *TwitterHandler) RedirectUrl() string {
	return ""
}

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName      string
	lastName       string
	twitterHandler TwitterHandler
}

func NewPerson(firstname, lastname string) Person {
	return Person{
		firstName: firstname,
		lastName:  lastname,
	}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("Your fullname is : %s %s", p.firstName, p.lastName)
}

func (p *Person) ID() string {
	return "12345"
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if !strings.HasPrefix(string(handler.Handle), "@") && len(handler.Handle) > 0 {
		return errors.New("Twitter handler missing important @ symbol.")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) GetTwitterhandler() TwitterHandler {
	return p.twitterHandler
}
