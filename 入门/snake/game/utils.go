package game

import (
	"math/rand"
	"snake/config"
)

func Init() {
	HeadPos = Cord{config.Height / 2, config.Weight / 2}
	BodyQueue = append(BodyQueue, Cord{config.Height / 2, config.Weight/2 - 1})
	Bodypos = make(map[Cord]bool)
	Bodypos[BodyQueue[len(BodyQueue)-1]] = true
	MoveDir = 0
	BodyLen = 1
	FoodPos = GetRandPos()
	FoodExist = true
}
func GetRandPos() Cord {
	x := (rand.Int() % config.Height) + 1
	y := (rand.Int() % config.Weight) + 1
	if Bodypos[Cord{x, y}] {
		return GetRandPos()
	}
	return Cord{x, y}
}

func OneStep() {
	BodyQueue = append(BodyQueue, HeadPos)
	Bodypos[HeadPos] = true
	HeadPos.X += config.Move[MoveDir][0]
	HeadPos.Y += config.Move[MoveDir][1]
}
func IsEated() bool {
	if FoodExist {
		return HeadPos.equl(FoodPos)
	}
	return false
}
func IsDead() bool {
	if HeadPos.X >= config.Height || HeadPos.X <= 0 || HeadPos.Y >= config.Weight || HeadPos.Y <= 0 {
		return true
	}
	return false
}
func Run() {

	if config.GameState != 1 {
		Init()
	} else {
		OneStep()
		if IsEated() {
			BodyLen++
			FoodPos = GetRandPos()
		} else {
			Bodypos[BodyQueue[0]] = false
			BodyQueue = BodyQueue[1:]
		}
		if IsDead() {
			config.GameState = 2
		}
	}
}
