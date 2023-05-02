package builder

import "personal_projects/go-design-pattern/constants"

type AK47Builder struct {
	Gun
}

func newAk47Builder() *AK47Builder {
	return &AK47Builder{}
}

func (a *AK47Builder) setMuzzleType() {
	a.muzzleType = constants.AK47_MUZZLE_TYPE
}

func (a *AK47Builder) setAmmoType() {
	a.ammoType = constants.AK47_AMMO_TYPE
}

func (a *AK47Builder) setCartridgeCapacity() {
	a.cartridgeCapacity = constants.AK47_CARTRIDGE_CAPACITY
}

func (a *AK47Builder) getAK47() Gun {
	return Gun{
		ammoType:          a.ammoType,
		muzzleType:        a.muzzleType,
		cartridgeCapacity: a.cartridgeCapacity,
	}
}
