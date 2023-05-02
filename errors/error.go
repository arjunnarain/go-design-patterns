package errors

type IError struct {
	InternalMessage string
	PublicMessage   string
	Code            int64
}

func NewCustomError(im string, pm string, c int64) (customError IError) {
	customError.PublicMessage = pm
	customError.InternalMessage = im
	customError.Code = c

	return
}
