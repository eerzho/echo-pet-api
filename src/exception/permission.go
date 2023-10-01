package exception

type PermissionDenied struct {
}

func NewPermissionDenied() *PermissionDenied {
	return &PermissionDenied{}
}

func (e *PermissionDenied) Error() string {
	return "Permission denied"
}
