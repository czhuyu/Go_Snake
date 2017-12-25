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

func main(){

	var gameMap [20][20]int
	var snake [50]string
	var orientation int
	//var orientation int
	//cmd := exec.Command("clear")
	//_ = cmd.Run()

	snake[0] = "0,5"
	//orientation = 2

	for {
		for x:= 0;x<20 ;x++  {
			fmt.Print("一")
		}
		fmt.Println()
		for z:=0; z<50; z++  {
			if("" != snake[z]) {
				tempCoordinate := strings.Split(snake[z], ",")
				CoordinateX, _ := strconv.Atoi(tempCoordinate[0])
				CoordinateY, _ := strconv.Atoi(tempCoordinate[1])
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
			orientation = move()
		}else{
			orientation = 2
		}

		if(orientation == 1){
			//向上移动蛇
			tempCoordinate := strings.Split(snake[0],",")
			tempX,_ := strconv.Atoi(tempCoordinate[0])
			tempY,_ := strconv.Atoi(tempCoordinate[1])
			snake[0] = strconv.Itoa(tempX-1) + "," + strconv.Itoa(tempY)
		}else if(orientation == 2){
			//向右移动蛇
			tempCoordinate := strings.Split(snake[0],",")
			tempX,_ := strconv.Atoi(tempCoordinate[0])
			tempY,_ := strconv.Atoi(tempCoordinate[1])
			snake[0] = strconv.Itoa(tempX) + "," + strconv.Itoa(tempY + 1)
		}else if(orientation ==3){
			//向下移动蛇
			tempCoordinate := strings.Split(snake[0],",")
			tempX,_ := strconv.Atoi(tempCoordinate[0])
			tempY,_ := strconv.Atoi(tempCoordinate[1])
			snake[0] = strconv.Itoa(tempX+1) + "," + strconv.Itoa(tempY)
		}else{
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

func move() int {
	key := C.getch()
	if(key == 119){
		return 1
	}else if(key == 100){
		return 2
	}else if(key == 115){
		return 3
	}else if(key == 97){
		return 4
	}else{
		return 2
	}
}