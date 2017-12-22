package main

import (
	"fmt"
	//"os/exec"
	"time"
)

func main(){

	var Map [20][20]int
	//cmd := exec.Command("clear")
	//_ = cmd.Run()

	for {
		for x:= 0;x<20 ;x++  {
			fmt.Print("一")
		}
		fmt.Println()
		for x:=0; x<20; x++ {
			fmt.Print("|")
			for y:=0; y<20; y++ {
				if(Map[x][y] == 1){
					fmt.Print("口")
				}else if(Map[x][y] == 0){
					fmt.Print("  ")
				}
			}
			fmt.Println("|")
		}
		for x:= 0;x<20 ;x++  {
			fmt.Print("一")
		}
		time.Sleep(time.Second)
		//cmd := exec.Command("cls")
		//_ = cmd.Run()
		for x:=1; x<20; x++{
			fmt.Println()
		}
	}
}