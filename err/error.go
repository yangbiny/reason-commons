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
