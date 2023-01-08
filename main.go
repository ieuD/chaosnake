package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World")

	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	screen.SetStyle(defStyle)

	snakeBody := SnakeBody{
		X:      5,
		Y:      10,
		XSpeed: 1,
		YSpeed: -1,
	}

	game := Game{
		Screen:    screen,
		snakeBody: snakeBody,
	}

	go game.Run()

	for {
		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				game.Screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				game.snakeBody.ChangeDir(-1, 0)
			} else if event.Key() == tcell.KeyDown {
				game.snakeBody.ChangeDir(1, 0)
			} else if event.Key() == tcell.KeyLeft {
				game.snakeBody.ChangeDir(0, -1)
			} else if event.Key() == tcell.KeyRight {
				game.snakeBody.ChangeDir(0, 1)
			}
		}
	}
}
