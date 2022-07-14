package game

import (
	"fmt"
	"os"
	"snake/config"

	"github.com/eiannone/keyboard"
)

func DealContorll() {
	keyEvent, err := keyboard.GetKeys(10)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetKeys : %v\n", err)
	}

	for {
		event := <-keyEvent
		if event.Err != nil {
			panic(event.Err)
		}

		if config.GameState == 1 {
			// fmt.Println(string(event.Rune))
			switch event.Key {
			case keyboard.KeyArrowDown:
				MoveDir = 1
			case keyboard.KeyArrowUp:
				MoveDir = 3
			case keyboard.KeyArrowLeft:
				MoveDir = 2
			case keyboard.KeyArrowRight:
				MoveDir = 0
			}
		} else if config.GameState == 0 {
			if event.Key == keyboard.KeyEnter {
				config.GameState = 1
			}
		} else {
			if event.Key == keyboard.KeyEnter {
				config.GameState = 0
			}
		}

	}
}
