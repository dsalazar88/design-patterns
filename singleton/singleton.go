package singleton

type person struct {
	name string
	age  int
}

func (p *person) setName(name string) {
	p.name = name
}

func (p *person) setAge(age int) {
	p.age = age
}

func (p *person) GetName() string {
	return p.name
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) IncrementAge() {
	p.age++
}
