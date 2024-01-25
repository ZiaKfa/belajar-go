package main

import "fmt"

type Pokemon struct {
	name              string
	maxhp, hp, attack int
}

func (p *Pokemon) receiveDamage(damage int) {
	p.hp -= damage
	fmt.Println(p.name, "got", damage, "damage")
}

func (p *Pokemon) heal(amount int) {
	p.hp += amount
	if p.hp > p.maxhp {
		p.hp = p.maxhp
	}
	fmt.Println(p.name, "got", amount, "heal")
}

func (p *Pokemon) attackEnemy(enemy *Pokemon) {
	enemy.receiveDamage(p.attack)
	fmt.Println(p.name, "attack", enemy.name, "with", p.attack, "damage")
}

func main() {
	var pikachu = Pokemon{"Pikachu", 100, 100, 20}
	fmt.Println(pikachu.name, "hp:", pikachu.hp, "attack:", pikachu.attack)
	pikachu.receiveDamage(10)
	fmt.Println(pikachu.name, "hp:", pikachu.hp, "attack:", pikachu.attack)
	pikachu.heal(30)
	fmt.Println(pikachu.name, "hp:", pikachu.hp, "attack:", pikachu.attack)

	var meowth = Pokemon{"Meowth", 100, 100, 10}
	fmt.Println(meowth.name, "hp:", meowth.hp, "attack:", meowth.attack)
	meowth.attackEnemy(&pikachu)
	fmt.Println(pikachu.name, "hp:", pikachu.hp, "attack:", pikachu.attack)
}
