package factory

import (
	"personal_projects/go-design-pattern/constants"
	"personal_projects/go-design-pattern/errors"
)

// TODO: Add a function for adding a new gun to the catalogue

func GetGun(gunType string) (IGun, errors.IError) {
	switch gunType {
	case constants.MUSKET:
		return newMusket(), errors.IError{}
	case constants.AK47:
		return newAk47(), errors.IError{}
	default:
		return nil, errors.NewCustomError(errors.InternalInvalidGunType, errors.PublicInvalidGunType, 200)
	}
}
