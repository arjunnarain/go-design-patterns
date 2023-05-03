package builder

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildGun() Gun {
	d.builder.setMuzzleType()
	d.builder.setAmmoType()
	d.builder.setCartridgeCapacity()
	return d.builder.getGun()
}
