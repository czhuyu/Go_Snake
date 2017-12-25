package main

import (
	// #include <conio.h>
	"C"
	"fmt"
)

func Move() {
	key := C.getch()
	fmt.Println(key)
}

func main() {
	Move()
}

