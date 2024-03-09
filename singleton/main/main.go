package main

import (
	"design-patterns/singleton"
	clientone "design-patterns/singleton/client-one"
	clienttwo "design-patterns/singleton/client-two"
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			clientone.IncrementAge()
		}()

		go func() {
			defer wg.Done()
			clienttwo.IncrementAge()
		}()
	}
	wg.Wait()

	p := singleton.GetInstance()
	age := p.GetAge()
	log.Println(age)
}
