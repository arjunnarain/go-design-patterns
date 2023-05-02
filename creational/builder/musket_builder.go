package builder

import "personal_projects/go-design-pattern/constants"

type MusketBuilder struct {
	Gun
}

func newMusketBuilder() *MusketBuilder {
	return &MusketBuilder{}
}

func (a *MusketBuilder) setMuzzleType() {
	a.muzzleType = constants.MUSKET_MUZZLE_TYPE
}

func (a *MusketBuilder) setAmmoType() {
	a.ammoType = constants.MUSKET_AMMO_TYPE
}

func (a *MusketBuilder) setCartridgeCapacity() {
	a.cartridgeCapacity = constants.MUSKET_CARTRIDGE_CAPACITY
}

func (a *MusketBuilder) getMusket() Gun {
	return Gun{
		ammoType:          a.ammoType,
		muzzleType:        a.muzzleType,
		cartridgeCapacity: a.cartridgeCapacity,
	}
}
