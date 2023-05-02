package factory

// IGun defines the common properties that each gun should have
type IGun interface {
	setName(name string)
	GetName() string
	setPower(power int64)
	GetPower() int64
}
