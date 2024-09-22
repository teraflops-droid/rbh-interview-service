package exception

type NotFoundError struct {
	Message string
}

type UnauthorizedError struct {
	Message string
}

type ValidationError struct {
	Message string
}

func PanicLogging(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func (notFoundError NotFoundError) Error() string {
	return notFoundError.Message
}

func (unauthorizedError UnauthorizedError) Error() string {
	return unauthorizedError.Message
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}
