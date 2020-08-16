// chapter 7 exercise 2 - why would you use an embedded anonymous field instead of a normal named field?

package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person //embedded type anonymous field here
	Model  string
}

func (p *Android) Talk() {
	fmt.Println("Hi, my name is", p.Name, "model", p.Model)
}

func main() {
	per := Person{"Kryten"}
	and := Android{per, "Series 4000"}
	//	per.Talk()
	and.Talk()
}
