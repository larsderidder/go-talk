package main

import "fmt"

type Team int

const (
	SERVICES Team = iota
	INTERFACES
	CORE
	OPIT
)

var teamToString = map[Team]string{
	SERVICES:   "Services",
	INTERFACES: "Interfaces",
	CORE:       "Core",
	OPIT:       "Operational IT",
}

func (t Team) String() string {
	return teamToString[t]
}

type Paylogician struct {
	name    string
	country string
	team    string
}

func (p *Paylogician) String() string {
	return fmt.Sprintf("%s from %s does some stuff in the %s team",
		p.name, p.country, p.team)
}

type Paylogician2 struct {
	name    string
	country string
	team    Team
}

func (p *Paylogician2) String() string {
	return fmt.Sprintf("%s from %s does some stuff in the %s team",
		p.name, p.country, p.team)
}

type Person struct {
	name    string
	country string
}

func whoAreYou(p *Person) string {
	return p.name
}

type Paylogician3 struct {
	Person
	team Team
}

func (p *Paylogician3) String() string {
	return fmt.Sprintf("%s from %s does some stuff in the %s team",
		p.name, p.country, p.team)
}

type Serializer interface {
	Serialize() string
}

func (p *Paylogician3) Serialize() string {
	return fmt.Sprintf("{ 'name': '%s', 'country': '%s', 'team': '%s' }",
		p.name, p.country, p.team)
}

func serializeMe(s Serializer) string {
	return s.Serialize()
}

func main() {
	oleg := Paylogician{name: "Oleg", country: "UKRAIN", team: "Interfaces"}
	fmt.Println(&oleg)
	siebejan := new(Paylogician)
	fmt.Println(siebejan.name) // This is ""
	siebejan.name = "Siebe Jan"
	siebejan.country = "Fryslân"
	siebejan.team = "Services"
	fmt.Println(siebejan) // No & necessary, new allocates pointer to data

	spyros := Paylogician2{name: "Spyros", country: "Ελλάδα", team: CORE}
	fmt.Println(&spyros)

	egon := Paylogician3{Person{name: "Egon", country: "Nederland"}, OPIT}
	fmt.Println(&egon)

	// name := whoAreYou(egon)

	wouter := Paylogician3{
		Person{name: "Wouter", country: "Nederland"}, SERVICES,
	}
	fmt.Println(serializeMe(&wouter))
}
