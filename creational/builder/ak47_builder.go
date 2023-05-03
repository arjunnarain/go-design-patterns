package builder

import "personal_projects/go-design-pattern/constants"

type AK47Builder struct {
	Gun
}

func newAk47Builder() *AK47Builder {
	return &AK47Builder{}
}

func (a *AK47Builder) setMuzzleType() {
	a.MuzzleType = constants.AK47_MUZZLE_TYPE
}

func (a *AK47Builder) setAmmoType() {
	a.AmmoType = constants.AK47_AMMO_TYPE
}

func (a *AK47Builder) setCartridgeCapacity() {
	a.CartridgeCapacity = constants.AK47_CARTRIDGE_CAPACITY
}

func (a *AK47Builder) getGun() Gun {
	return Gun{
		AmmoType:          a.AmmoType,
		MuzzleType:        a.MuzzleType,
		CartridgeCapacity: a.CartridgeCapacity,
	}
}
