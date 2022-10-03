package main

import (
	"fmt"

	"github.com/gBenkyous/inTime/tree/master/server/pkg/keisan"
)

//上記に関してはgithubに上げてみんなに使ってほしいわけではないのでリモート用にする必要性はないのかもしれない。

func main() {
	var i int = 1
	i = keisan.Tasizan(1, 2)
	fmt.Println(i)
}
