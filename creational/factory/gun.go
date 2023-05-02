package factory

// Gun defines a normal gun
type Gun struct {
	name  string
	power int64
}

func (g *Gun) GetName() string {
	return g.name
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) GetPower() int64 {
	return g.power
}

func (g *Gun) setPower(power int64) {
	g.power = power
}
