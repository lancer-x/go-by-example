package main

import "fmt"

type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "kitty"}
}

type Warrior struct {
	Name string
}

func NewWarrior(name string) Warrior {
	if name != "" {
		return Warrior{Name: name}
	}
	return Warrior{Name: "wusong"}
}

type Mission struct {
	Warrior Warrior
	Monster Monster
}

func NewMission(w Warrior, m Monster) Mission {
	return Mission{w, m}
}

func (m Mission) start() {
	fmt.Printf("%s defeats %s", m.Warrior.Name, m.Monster.Name)
}

func main() {
	monster := NewMonster()
	warrior := NewWarrior("")
    mission := NewMission(warrior, monster)
    mission.start()

    mission2 := InitMission("武松")
    mission2.start()
}
