package main

import (
	"fmt"
	"time"
	"github.com/nsf/termbox-go"
)

func main() {


	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	events := make(chan termbox.Event, 1000)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()

	for true {
		select {
		case ev := <-events:
			if ev.Type == termbox.EventKey {
				// exit the game
				if ev.Key == termbox.KeyArrowUp {
					fmt.Println("you press key up")
				}
			}

		default:

		}
		fmt.Println("---")
		time.Sleep(time.Second)
	}

	termbox.Close()
	fmt.Println("over")
}
