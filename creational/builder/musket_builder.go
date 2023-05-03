package builder

import "personal_projects/go-design-pattern/constants"

type MusketBuilder struct {
	Gun
}

func newMusketBuilder() *MusketBuilder {
	return &MusketBuilder{}
}

func (a *MusketBuilder) setMuzzleType() {
	a.MuzzleType = constants.MUSKET_MUZZLE_TYPE
}

func (a *MusketBuilder) setAmmoType() {
	a.AmmoType = constants.MUSKET_AMMO_TYPE
}

func (a *MusketBuilder) setCartridgeCapacity() {
	a.CartridgeCapacity = constants.MUSKET_CARTRIDGE_CAPACITY
}

func (a *MusketBuilder) getGun() Gun {
	return Gun{
		AmmoType:          a.AmmoType,
		MuzzleType:        a.MuzzleType,
		CartridgeCapacity: a.CartridgeCapacity,
	}
}
