package main

import "fmt"

type Player struct {
	Name        string
	HealthPoint int
	Level       int
	NowPosition []int
	Prop        []string
}

func NewPlayer(name string, hp int, level int, np []int, prop []string) *Player {

	return &Player{
		Name:        name,
		HealthPoint: hp,
		Level:       level,
		NowPosition: np,
		Prop:        prop,
	}
}

func (p Player) attack() {
	fmt.Printf("%s发起攻击！\n", p.Name)
}
func (p *Player) attacked() {
	fmt.Printf("%s被攻击！\n", p.Name)
	p.HealthPoint -= 10
	fmt.Println(p.HealthPoint)
}

func (p *Player) buyProp(prop string) {
	p.Prop = append(p.Prop, prop)
	fmt.Printf("%s购买道具！\n", p.Name)
}

func main() {
	player := NewPlayer("lin", 100, 100, []int{10, 100}, nil)
	player.attack()
	player.attacked()
	fmt.Println(player.HealthPoint)
	player.buyProp("魔法石")
	fmt.Println(player.Prop)
	fmt.Println(player)
}
