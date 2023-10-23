package err

type ApiStatus int

const (
	Success           ApiStatus = 2
	RequestParamError           = 4
	Fail                        = 5
)

type ApiError struct {
	ApiStatus ApiStatus
	Err       error
}

func WhenServiceError(err error) *ApiError {
	return &ApiError{
		ApiStatus: Fail,
		Err:       err,
	}
}

func WhenParamError(err error) *ApiError {
	return &ApiError{
		ApiStatus: RequestParamError,
		Err:       err,
	}
}
