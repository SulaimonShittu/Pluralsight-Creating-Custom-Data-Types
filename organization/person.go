package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName      string
	lastName       string
	twitterHandler string
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

func (p *Person) SetTwitterHandler(handler string) error {
	if !strings.HasPrefix(handler, "@") && len(handler) > 0 {
		return errors.New("Twitter handler missing important @ symbol.")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) GetTwitterhandler() string {
	return p.twitterHandler
}
