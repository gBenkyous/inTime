package main

import (
	"fmt"
	"log"

	"intimeServer/pkg/keisan"
)

func main() {
	var i int = 1
	i = keisan.Tasizan(1, 2)
	fmt.Println(i)
	log.Println("OK")
}
