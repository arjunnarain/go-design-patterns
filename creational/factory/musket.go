package factory

import "personal_projects/go-design-pattern/constants"

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  constants.MUSKET,
			power: constants.MUSKET_POWER,
		},
	}
}
