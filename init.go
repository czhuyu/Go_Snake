package main

import (
	// #include <conio.h>
	"C"
	"fmt"
	"time"
	"strings"
	"strconv"
	"math/rand"
)

var gameMap [20][20]int
var snake [50]string
var orientation = 2
var movecount = 0
var headCoordinate = "0,5" //初始蛇头
var tailCoordinate = "0,5" //初始蛇尾
var foodCoordinate = "-1,-1"
var snakeLen = 3

func main(){
	snake[0] = headCoordinate
	for {
		drawMap()
		moveSnake()
		movecount++
		resetMap()
	}
}

func changeOrientation(before int) int {
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

func moveSnake(){
	//非阻塞获取键盘输入
	if(C.kbhit() == 1){
		orientation = changeOrientation(orientation)
	}

	if(orientation == 1){
		//向上移动蛇
		tempheadCoordinate := strings.Split(snake[0],",")
		tempHeadX,_ := strconv.Atoi(tempheadCoordinate[0])
		snake[0] = strconv.Itoa(tempHeadX-1) + "," + tempheadCoordinate[1]
		headCoordinate = snake[0]
		if(movecount >= 1){
			snake[2] = tailCoordinate
		}
		snake[1] = tempheadCoordinate[0] + ","+ tempheadCoordinate[1]
		tailCoordinate = snake[1]
		isGetFood()
	}else if(orientation == 2){
		//向右移动蛇
		tempheadCoordinate := strings.Split(snake[0],",")
		tempHeadY,_ := strconv.Atoi(tempheadCoordinate[1])
		snake[0] = tempheadCoordinate[0] + "," + strconv.Itoa(tempHeadY + 1)
		headCoordinate = snake[0]
		if(movecount >= 1){
			snake[2] = tailCoordinate
		}
		snake[1] = tempheadCoordinate[0] + ","+ tempheadCoordinate[1]
		tailCoordinate = snake[1]
		isGetFood()
	}else if(orientation ==3){
		//向下移动蛇
		tempheadCoordinate := strings.Split(snake[0],",")
		tempHeadX,_ := strconv.Atoi(tempheadCoordinate[0])
		snake[0] = strconv.Itoa(tempHeadX+1) + "," + tempheadCoordinate[1]
		headCoordinate = snake[0]
		if(movecount >= 1){
			snake[2] = tailCoordinate
		}
		snake[1] = tempheadCoordinate[0] + ","+ tempheadCoordinate[1]
		tailCoordinate = snake[1]
		isGetFood()
	}else{
		//向左移动蛇
		tempheadCoordinate := strings.Split(snake[0],",")
		tempHeadY,_ := strconv.Atoi(tempheadCoordinate[1])
		snake[0] = tempheadCoordinate[0] + "," + strconv.Itoa(tempHeadY - 1)
		headCoordinate = snake[0]
		if(movecount >= 1){
			snake[2] = tailCoordinate
		}
		snake[1] = tempheadCoordinate[0] + ","+ tempheadCoordinate[1]
		tailCoordinate = snake[1]
		isGetFood()
	}
}

func drawMap(){
	if(movecount >= 2 && foodCoordinate == "-1,-1"){
		drawFood()
	}
	for x:= 0;x<20 ;x++  {
		fmt.Print("一")
	}
	fmt.Println()
	for z:=0; z<50; z++  {
		if("" != snake[z]) {
			tempheadCoordinate := strings.Split(snake[z], ",")
			CoordinateX, _ := strconv.Atoi(tempheadCoordinate[0])
			CoordinateY, _ := strconv.Atoi(tempheadCoordinate[1])
			gameMap[CoordinateX][CoordinateY] = 1
		}
		if(foodCoordinate != "-1,-1"){
			tempheadCoordinate := strings.Split(foodCoordinate, ",")
			CoordinateX, _ := strconv.Atoi(tempheadCoordinate[0])
			CoordinateY, _ := strconv.Atoi(tempheadCoordinate[1])
			gameMap[CoordinateX][CoordinateY] = 2
		}
	}
	for x:=0; x<20; x++ {
		fmt.Print("|")
		for y:=0; y<20; y++ {
			if(gameMap[x][y] == 1){
				fmt.Print("口")
			}else if(gameMap[x][y] == 2){
				fmt.Print("△")
			}else{
				fmt.Print("  ")
			}
		}
		fmt.Println("|")
	}
	for x:= 0;x<20 ;x++  {
		fmt.Print("一")
	}
}

func resetMap(){
	for i:=0; i<20; i++ {
		for j:=0; j<20; j++ {
			gameMap[i][j] = 0
		}
	}
	time.Sleep(time.Second)
	for x:=1; x<20; x++{
		fmt.Println()
	}
}

func drawFood(){
	for {
		randX := generateRandnum()
		randY := generateRandnum()
		if(gameMap[randX][randY] == 0){
			gameMap[randX][randY] = 2
			foodCoordinate = strconv.Itoa(randX) + "," + strconv.Itoa(randY)
			return
		}else{
			continue
		}
	}

}
func generateRandnum() int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(19)
	return randNum
}

func isGetFood(){
	if(headCoordinate == foodCoordinate){
		//吃到了食物
		drawFood()
		//蛇身增长
	}else{
		return
	}
}

func addSnakeLen(){
	snake[snakeLen] = tailCoordinate
}