package context

import (
	"RouteHub.Service.Dashboard/ent"
)

const (
	personContextKey = "person"
)

func (cc *ServerEchoContext) GetPerson() *ent.Person {

	if cc.Get(personContextKey) == nil {
		return nil
	}

	return cc.Get(personContextKey).(*ent.Person)
}

func (cc *ServerEchoContext) SetPerson(person *ent.Person) {
	cc.Set(personContextKey, person)
}
