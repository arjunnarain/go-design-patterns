package main

import (
	"fmt"
	"personal_projects/go-design-pattern/constants"
	"personal_projects/go-design-pattern/creational/builder"
)

func main() {
	ak47Builder, _ := builder.GetBuilder(constants.AK47)
	musketBuilder, _ := builder.GetBuilder(constants.MUSKET)

	manufacturer := builder.NewDirector(ak47Builder)
	ak47 := manufacturer.BuildGun()

	fmt.Printf("AK 47 Muzzle Type: %s", ak47.MuzzleType)
	fmt.Printf("AK 47 Ammo Type: %s", ak47.AmmoType)
	fmt.Printf("Ak 47 Cartridge Capacity: %d", ak47.CartridgeCapacity)

	manufacturer.SetBuilder(musketBuilder)
	musket := manufacturer.BuildGun()

	fmt.Printf("Musket Muzzle Type: %s", musket.MuzzleType)
	fmt.Printf("Musket Ammo Type: %s", musket.AmmoType)
	fmt.Printf("Musket Cartridge Capacity: %d", musket.CartridgeCapacity)
}
