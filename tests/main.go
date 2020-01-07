package main

import (
	"fmt"
	"github.com/aszecowka/modexamples/component"
)

func main() {
	cli := component.Client{}
	fmt.Println(cli.DoIt())
}
