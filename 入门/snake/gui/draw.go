package gui

import (
	"fmt"
	"snake/config"
	"snake/game"
)

func Draw() {
	SimpGui := []string{}
	for i := 0; i < config.Height+1; i++ {
		temp := ""
		for j := 0; j < config.Weight+1; j++ {
			add := ""
			if j == 0 || j == config.Weight {
				add = "┃"
			} else if i == 0 {
				add = "▔"
			} else if i == config.Height {
				add = "▁"
			} else {
				add = " "
			}
			switch config.GameState {
			case 0:
				if i == config.Height/2 && j == config.Weight/2 {
					SimpGui = append(SimpGui, "            按enter开始游戏            ")
					i++
				}
			case 2:
				if i == config.Height/2 && j == config.Weight/2 {
					SimpGui = append(SimpGui, "            GAME OVER            ")
					i++
				}
			case 1:
				if game.Bodypos[game.Cord{X: i, Y: j}] {
					add = config.SnakeBody
				}
				if i == game.FoodPos.X && j == game.FoodPos.Y {
					add = config.FoodUi
				}
				if i == game.HeadPos.X && j == game.HeadPos.Y {
					add = config.SnakeHead
				}
			}
			temp += add
		}
		SimpGui = append(SimpGui, temp)
	}

	for _, s := range SimpGui {
		fmt.Println(s)
	}
}
