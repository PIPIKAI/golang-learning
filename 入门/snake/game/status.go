package game

type Cord struct {
	X int
	Y int
}

func (a Cord) equl(b Cord) bool {
	return (a.X == b.X && a.Y == b.Y)
}

var HeadPos Cord
var Bodypos map[Cord]bool
var BodyQueue []Cord
var FoodPos Cord
var MoveDir = 0
var BodyLen = 1
var FoodExist bool = false
