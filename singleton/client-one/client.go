package clientone

import "design-patterns/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
