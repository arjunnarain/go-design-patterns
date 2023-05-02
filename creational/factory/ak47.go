package factory

import "personal_projects/go-design-pattern/constants"

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  constants.AK47,
			power: constants.AK47_POWER,
		},
	}
}
