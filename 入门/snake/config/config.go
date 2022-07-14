package config

const SnakeHead string = "◯"
const SnakeBody string = "■"
const FoodUi string = "◙"
const Height int = 30
const Weight int = 60

var Move [4][2]int = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
var Speed uint = 250

// 0 代表在开始界面
// 1 代表游戏开始
// 2 游戏结束
var GameState int = 0
