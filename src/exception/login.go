package exception

type InvalidLoginError struct {
}

func NewInvalidLoginError() *InvalidLoginError {
	return &InvalidLoginError{}
}

func (e *InvalidLoginError) Error() string {
	return "Invalid email or password"
}
