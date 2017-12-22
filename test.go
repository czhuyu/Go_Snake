package main

//import "fmt"
import (
	termbox "github.com/nsf/termbox-go"
	//"golang.org/x/net/html/atom"
	//"flag"
	"log"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	key := make(chan string)

	go func() {
		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyArrowUp:
					key <- "1"
				case termbox.KeyArrowRight:
					key <- "2"
				case termbox.KeyArrowDown:
					key <- "3"
				case termbox.KeyArrowLeft:
					key <- "4"
				default:
					key <- "hhhh"
				}
			}
		}
	}()
	for{
		log.Println(<-key)
	}
}