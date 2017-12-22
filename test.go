package main

import (
	//"fmt"
	"os/exec"
	//"time"
)

func main(){

	cmd := exec.Command("cls")
	_ = cmd.Run()
}