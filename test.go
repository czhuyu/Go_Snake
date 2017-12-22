package main

import (
	"fmt"
	//"os/exec"
	//"time"
	"strings"
	"reflect"
)

func main(){
	var snake [50]string
	//cmd := exec.Command("clear")
	//_ = cmd.Run()
	snake[0] = "0,0"
	s := strings.Split(snake[0],",")
	x, y := s[0], s[1]
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(y))
}