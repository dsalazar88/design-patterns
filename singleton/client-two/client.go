package clienttwo

import "design-patterns/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
