package main

import (
	"fmt"
	"personal_projects/go-design-pattern/constants"
	"personal_projects/go-design-pattern/creational/factory"
)

func main() {
	ak47, _ := factory.GetGun(constants.AK47)
	musket, _ := factory.GetGun(constants.MUSKET)

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g factory.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
