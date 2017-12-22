package main

import (
	"fmt"
	//"os/exec"
	"time"
	"strings"
	"strconv"
     //termbox "github.com/nsf/termbox-go"
	//"github.com/go-xorm/builder"
	"github.com/nsf/termbox-go"
	"log"
)

func main(){

	var gameMap [20][20]int
	var snake [50]string
	//var orientation int
	//cmd := exec.Command("clear")
	//_ = cmd.Run()

	snake[0] = "0,0"
	//orientation = 2

	for {
		for x:= 0;x<20 ;x++  {
			fmt.Print("一")
		}
		fmt.Println()
		for z:=0; z<50; z++  {
			if("" != snake[z]){
				tempCoordinate := strings.Split(snake[z],",")
				CoordinateX,_ := strconv.Atoi(tempCoordinate[0])
				CoordinateY,_ := strconv.Atoi(tempCoordinate[1])
				gameMap[CoordinateX][CoordinateY] = 1
			}
		}
		for x:=0; x<20; x++ {
			fmt.Print("|")
			for y:=0; y<20; y++ {
				if(gameMap[x][y] == 1){
					fmt.Print("口")
				}else if(gameMap[x][y] == 0){
					fmt.Print("  ")
				}
			}
			fmt.Println("|")
		}
		for x:= 0;x<20 ;x++  {
			fmt.Print("一")
		}



		key := make(chan string)
		go func() {
			err := termbox.Init()
			if err != nil {
				panic(err)
			}
			defer termbox.Close()
		Loop:
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
						key <- "2"
						break Loop
					}
				default:
					key <- "2"
					break Loop
				}
				log.Println(key)
			}
		}()

		orientation := <-key

		log.Println(orientation)

		if(orientation == "1"){
			//向上移动蛇
			tempCoordinate := strings.Split(snake[0],",")
			tempX,_ := strconv.Atoi(tempCoordinate[0])
			tempY,_ := strconv.Atoi(tempCoordinate[1])
			snake[0] = strconv.Itoa(tempX-1) + "," + strconv.Itoa(tempY)
		}
		if(orientation == "2"){
			//向右移动蛇
			tempCoordinate := strings.Split(snake[0],",")
			tempX,_ := strconv.Atoi(tempCoordinate[0])
			tempY,_ := strconv.Atoi(tempCoordinate[1])
			snake[0] = strconv.Itoa(tempX) + "," + strconv.Itoa(tempY + 1)
		}
		if(orientation == "3"){
			//向下移动蛇
			tempCoordinate := strings.Split(snake[0],",")
			tempX,_ := strconv.Atoi(tempCoordinate[0])
			tempY,_ := strconv.Atoi(tempCoordinate[1])
			snake[0] = strconv.Itoa(tempX+1) + "," + strconv.Itoa(tempY)
		}
		if(orientation == "4"){
			//向左移动蛇
			tempCoordinate := strings.Split(snake[0],",")
			tempX,_ := strconv.Atoi(tempCoordinate[0])
			tempY,_ := strconv.Atoi(tempCoordinate[1])
			snake[0] = strconv.Itoa(tempX) + "," + strconv.Itoa(tempY - 1)
		}


		for i:=0; i<20; i++ {
			for j:=0; j<20; j++ {
				gameMap[i][j] = 0
			}
		}


		time.Sleep(time.Second)
		//  cmd := exec.Command("cls")
		//_ = cmd.Run()
		for x:=1; x<20; x++{
			fmt.Println()
		}
	}
}