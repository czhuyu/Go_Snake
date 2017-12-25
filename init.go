package main

import (
	// #include <conio.h>
	"C"

	"fmt"
	//"os/exec"
	"time"
	"strings"
	"strconv"
	//"github.com/nsf/termbox-go"
	//"github.com/go-xorm/builder"
	//"github.com/nsf/termbox-go"
	//"log"
)

var gameMap [20][20]int
var snake [50]string
var orientation = 2
var movecount = 0
var HeadCoordinate = "0,5"
var tailCoordinate = "0,5"

func main(){


	//var orientation int
	//cmd := exec.Command("clear")
	//_ = cmd.Run()

	snake[0] = HeadCoordinate
	//orientation = 2

	for {
		for x:= 0;x<20 ;x++  {
			fmt.Print("一")
		}
		fmt.Println()
		for z:=0; z<50; z++  {
			if("" != snake[z]) {
				tempHeadCoordinate := strings.Split(snake[z], ",")
				CoordinateX, _ := strconv.Atoi(tempHeadCoordinate[0])
				CoordinateY, _ := strconv.Atoi(tempHeadCoordinate[1])
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


		//log.Println(<-key)
		if(C.kbhit() == 1){
			orientation = move(orientation)
		}

		if(orientation == 1){
			//向上移动蛇
			tempHeadCoordinate := strings.Split(snake[0],",")
			tempHeadX,_ := strconv.Atoi(tempHeadCoordinate[0])
			tempHeadY,_ := strconv.Atoi(tempHeadCoordinate[1])
			snake[0] = strconv.Itoa(tempHeadX-1) + "," + strconv.Itoa(tempHeadY)
			HeadCoordinate = snake[0]
			if(movecount >= 1){
				snake[3] = tailCoordinate
			}
			snake[1] = tempHeadCoordinate[0] + ","+ tempHeadCoordinate[1]
			tailCoordinate = snake[1]
		}else if(orientation == 2){
			//向右移动蛇
			tempHeadCoordinate := strings.Split(snake[0],",")
			tempHeadX,_ := strconv.Atoi(tempHeadCoordinate[0])
			tempHeadY,_ := strconv.Atoi(tempHeadCoordinate[1])
			snake[0] = strconv.Itoa(tempHeadX) + "," + strconv.Itoa(tempHeadY + 1)
			HeadCoordinate = snake[0]
			if(movecount >= 1){
				snake[3] = tailCoordinate
			}
			snake[1] = tempHeadCoordinate[0] + ","+ tempHeadCoordinate[1]
			tailCoordinate = snake[1]
		}else if(orientation ==3){
			//向下移动蛇
			tempHeadCoordinate := strings.Split(snake[0],",")
			tempHeadX,_ := strconv.Atoi(tempHeadCoordinate[0])
			tempHeadY,_ := strconv.Atoi(tempHeadCoordinate[1])
			snake[0] = strconv.Itoa(tempHeadX+1) + "," + strconv.Itoa(tempHeadY)
			HeadCoordinate = snake[0]
			if(movecount >= 1){
				snake[3] = tailCoordinate
			}
			snake[1] = tempHeadCoordinate[0] + ","+ tempHeadCoordinate[1]
			tailCoordinate = snake[1]
		}else{
			//向左移动蛇
			tempHeadCoordinate := strings.Split(snake[0],",")
			tempHeadX,_ := strconv.Atoi(tempHeadCoordinate[0])
			tempHeadY,_ := strconv.Atoi(tempHeadCoordinate[1])
			snake[0] = strconv.Itoa(tempHeadX) + "," + strconv.Itoa(tempHeadY - 1)
			HeadCoordinate = snake[0]
			if(movecount >= 1){
				snake[3] = tailCoordinate
			}
			snake[1] = tempHeadCoordinate[0] + ","+ tempHeadCoordinate[1]
			tailCoordinate = snake[1]
		}

		movecount++


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

func move(before int) int {
	key := C.getch()
	if(key == 119 && before != 3){
		return 1
	}else if(key == 100 && before != 4){
		return 2
	}else if(key == 115 && before != 1){
		return 3
	}else if(key == 97 && before !=2){
		return 4
	}else{
		return before
	}
}

func getSnakeLen() int {
	i := 0
	for z:=0; z<50; z++  {
		if("" != snake[z]) {
			i++
		}
	}
	return i
}