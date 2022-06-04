package user

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) GetName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) SetName(firstName, lastName string) {
	p.FirstName = firstName
	p.LastName = lastName
}
