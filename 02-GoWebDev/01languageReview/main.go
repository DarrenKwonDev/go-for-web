package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	say()
}

func saySomething(h human) {
	h.say()
}

func (p person) say() {
	fmt.Println(p.fname, `says, good`)
}

func (s secretAgent) intro() {
	fmt.Println(s.lname, `has license`)
}

func main() {
	me := person{"darren", "kwon"}
	me.say()
	saySomething(me)
	james := secretAgent{
		person{"james", "bond"},
		true,
	}
	james.say()
	james.intro()
}
