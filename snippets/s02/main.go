package main

import (
	"fmt"
)

func main() {
	cities := []string{"Dublin", "Berlin", "Rome"}

	for _, city := range cities {
		fmt.Println(city)
	}
}
