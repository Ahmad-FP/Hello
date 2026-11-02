package main

import (
	"ex21/internal/greet"
	"ex21/internal/mixed"
	"fmt"
)

func main() {
	fmt.Println(greet.Hello("Budi"))
	fmt.Println(mixed.Truncate("budi", 4))
	fmt.Println(mixed.Truncate("budi", 3))
	
}
