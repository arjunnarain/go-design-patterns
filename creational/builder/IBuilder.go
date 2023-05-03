package builder

import (
	"personal_projects/go-design-pattern/constants"
	"personal_projects/go-design-pattern/errors"
)

type IBuilder interface {
	setMuzzleType()
	setCartridgeCapacity()
	setAmmoType()

	getGun() Gun
}

func GetBuilder(builderType string) (IBuilder, errors.IError) {
	switch builderType {
	case constants.AK47:
		return newAk47Builder(), errors.IError{}
	case constants.MUSKET:
		return newMusketBuilder(), errors.IError{}
	default:
		return nil, errors.NewCustomError(errors.InternalInvalidGunType, errors.PublicInvalidGunType, 400)
	}
}
