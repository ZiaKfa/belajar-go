package main

import "fmt"

type Player interface {
	Attack() int
	GetName() string
}

type Mage struct {
	Name   string
	damage int
}

type Assassin struct {
	Name   string
	damage int
	speed  int
}

func (m Mage) Attack() int { // Mage implements Player interface
	return m.damage * 2
}

func (m Mage) GetName() string { // Mage implements Player interface
	return m.Name
}

func (a Assassin) Attack() int { // Assassin implements Player interface
	return a.damage * a.speed
}

func (a Assassin) GetName() string { // Assassin implements Player interface
	return a.Name
}

func CastSpell(p Player) {
	fmt.Println(p.GetName(), "casted a spell and dealt", p.Attack(), "damage")
}

func Backstab(p Player) {
	fmt.Println(p.GetName(), "backstabbed and dealt", p.Attack(), "damage")
}

//empty interface
func Omg() interface{} {
	return "omg"
}

func main() {
	wizard := Mage{
		Name:   "Gandalf",
		damage: 12,
	}
	CastSpell(wizard)

	rogue := Assassin{
		Name:   "Rikimaru",
		damage: 10,
		speed:  3,
	}
	Backstab(rogue)
	var omg interface{} = Omg()
	fmt.Println(omg)
}
