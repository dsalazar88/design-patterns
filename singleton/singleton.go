package singleton

import (
	"sync"
)

var (
	p    *person
	once sync.Once
)

func GetInstance() *person {
	once.Do(func() {
		p = &person{}
	})

	return p
}

type person struct {
	name string
	age  int
	mux  sync.Mutex
}

func (p *person) SetName(name string) {
	p.mux.Lock()
	p.name = name
	p.mux.Unlock()
}

func (p *person) SetAge(age int) {
	p.mux.Lock()
	p.age = age
	p.mux.Unlock()
}

func (p *person) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *person) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}

func (p *person) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}
