package main

import (
	"fmt"

	"intimeServer/pkg/keisan"
)

func main() {
	var i int = 1
	i = keisan.Tasizan(1, 2)
	fmt.Println(i)
}
