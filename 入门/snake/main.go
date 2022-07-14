package main

import (
	// "os"
	"os"
	"os/exec"
	"snake/config"
	"snake/game"
	"snake/gui"
	"time"
)

func main() {

	go game.DealContorll()
	for {
		c := exec.Command("cmd", "/c", "cls")
		c.Stdout = os.Stdout
		c.Run()
		game.Run()
		gui.Draw()
		time.Sleep(time.Duration(config.Speed) * time.Millisecond)

	}

}
