package main

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"math/rand"
	"github.com/nsf/termbox-go"
)

var gameMap [20][20]int
var snake [50]string
var orientation = 2
var score = 0
var headCoordinate = "0,0" //初始蛇头
var foodCoordinate = "-1,-1"
var snakeLen = 3

func main() {
	snake[0] = headCoordinate
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

	for {
		select {
		case ev := <-events:
			if ev.Type == termbox.EventKey {
				if ev.Key == termbox.KeyArrowRight && orientation != 4 {
					orientation = 2
				}
				if ev.Key == termbox.KeyArrowLeft && orientation != 2 {
					orientation = 4
				}
				if ev.Key == termbox.KeyArrowUp && orientation != 3 {
					orientation = 1
				}
				if ev.Key == termbox.KeyArrowDown && orientation != 1 {
					orientation = 3
				}
			}

		default:

		}

		drawMap()
		moveSnake()
		score++
		resetMap()
	}
}

func moveSnake() {
	if orientation == 1 {
		//向上移动蛇
		tempheadCoordinate := strings.Split(snake[0], ",")
		tempHeadX, _ := strconv.Atoi(tempheadCoordinate[0])
		snake[0] = strconv.Itoa(tempHeadX-1) + "," + tempheadCoordinate[1]
		headCoordinate = snake[0]
		if score >= 1 {
			if snakeLen >= 3 {
				if isGetFood() {
					snakeLen++
					drawFood()
				}
				for i := 1; i <= snakeLen-3; i++ {
					snake[snakeLen-i] = snake[snakeLen-(i+1)]
				}
			}
			snake[2] = snake[1]
		}
		snake[1] = tempheadCoordinate[0] + "," + tempheadCoordinate[1]
	} else if orientation == 2 {
		//向右移动蛇
		tempheadCoordinate := strings.Split(snake[0], ",")
		tempHeadY, _ := strconv.Atoi(tempheadCoordinate[1])
		snake[0] = tempheadCoordinate[0] + "," + strconv.Itoa(tempHeadY+1)
		headCoordinate = snake[0]
		if score >= 1 {
			if snakeLen >= 3 {
				if isGetFood() {
					snakeLen++
					drawFood()
				}
				for i := 1; i <= snakeLen-3; i++ {
					snake[snakeLen-i] = snake[snakeLen-(i+1)]
				}
			}
			snake[2] = snake[1]
		}
		snake[1] = tempheadCoordinate[0] + "," + tempheadCoordinate[1]
	} else if orientation == 3 {
		//向下移动蛇
		tempheadCoordinate := strings.Split(snake[0], ",")
		tempHeadX, _ := strconv.Atoi(tempheadCoordinate[0])
		snake[0] = strconv.Itoa(tempHeadX+1) + "," + tempheadCoordinate[1]
		headCoordinate = snake[0]
		if score >= 1 {
			if snakeLen >= 3 {
				if isGetFood() {
					snakeLen++
					drawFood()
				}
				for i := 1; i <= snakeLen-3; i++ {
					snake[snakeLen-i] = snake[snakeLen-(i+1)]
				}
			}
			snake[2] = snake[1]
		}
		snake[1] = tempheadCoordinate[0] + "," + tempheadCoordinate[1]
	} else {
		//向左移动蛇
		tempheadCoordinate := strings.Split(snake[0], ",")
		tempHeadY, _ := strconv.Atoi(tempheadCoordinate[1])
		snake[0] = tempheadCoordinate[0] + "," + strconv.Itoa(tempHeadY-1)
		headCoordinate = snake[0]
		if score >= 1 {
			if snakeLen >= 3 {
				if isGetFood() {
					snakeLen++
					drawFood()
				}
				for i := 1; i <= snakeLen-3; i++ {
					snake[snakeLen-i] = snake[snakeLen-(i+1)]
				}
			}
			snake[2] = snake[1]
		}
		snake[1] = tempheadCoordinate[0] + "," + tempheadCoordinate[1]
	}
}

func drawMap() {
	if score >= 2 && foodCoordinate == "-1,-1" {
		drawFood()
	}
	for x := 0; x < 20; x++ {
		fmt.Print("一")
	}
	fmt.Println()
	for z := 0; z < 50; z++ {
		if "" != snake[z] {
			tempheadCoordinate := strings.Split(snake[z], ",")
			CoordinateX, _ := strconv.Atoi(tempheadCoordinate[0])
			CoordinateY, _ := strconv.Atoi(tempheadCoordinate[1])
			gameMap[CoordinateX][CoordinateY] = 1
		}
		if foodCoordinate != "-1,-1" {
			tempheadCoordinate := strings.Split(foodCoordinate, ",")
			CoordinateX, _ := strconv.Atoi(tempheadCoordinate[0])
			CoordinateY, _ := strconv.Atoi(tempheadCoordinate[1])
			gameMap[CoordinateX][CoordinateY] = 2
		}
	}
	for x := 0; x < 20; x++ {
		fmt.Print("|")
		for y := 0; y < 20; y++ {
			if gameMap[x][y] == 1 {
				fmt.Print("口")
			} else if gameMap[x][y] == 2 {
				fmt.Print("△")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("|")
	}
	for x := 0; x < 20; x++ {
		fmt.Print("一")
	}
}

func resetMap() {
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			gameMap[i][j] = 0
		}
	}
	time.Sleep(time.Second)
	for x := 1; x < 10; x++ {
		fmt.Println()
	}
}

func drawFood() {
	for {
		randX := generateRandnum()
		randY := generateRandnum()
		if gameMap[randX][randY] == 0 {
			gameMap[randX][randY] = 2
			foodCoordinate = strconv.Itoa(randX) + "," + strconv.Itoa(randY)
			return
		} else {
			continue
		}
	}

}
func generateRandnum() int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(19)
	return randNum
}

func isGetFood() bool {
	if headCoordinate == foodCoordinate {
		//吃到了食物
		return true
	} else {
		return false
	}
}
