package runner

func Notice(msg string) *NoticeError {
	return &NoticeError{msg}
}

type NoticeError struct {
	msg string
}

func (e *NoticeError) Error() string {
	return e.msg
}