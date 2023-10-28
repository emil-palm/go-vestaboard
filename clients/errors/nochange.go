package errors

type NoChangeError struct{}

func NewNoChangeError() NoChangeError {
	return NoChangeError{}
}

func (nce NoChangeError) Error() string {
	return "No change of content on board"
}
