package builder

type IBuilder interface {
	setMuzzleType()
	setCartridgeCapacity()
	setAmmoType()

	getGun() Gun
}
