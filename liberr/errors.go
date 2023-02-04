package liberr

type Error struct {
	HttpStatus int    `json:"-"`
	Errno      int    `json:"errno"`
	Msg        string `json:"msg"`
}

func (t *Error) Error() string {
	return t.Msg
}

func (t *Error) String() string {
	return t.Msg
}

func FormatError(err error) *Error {
	if v, ok := err.(*Error); ok {
		return v
	}
	res := *UnknownError
	res.Msg = err.Error()
	return &res
}

var (
	NotFound   = &Error{404, 404, "unsupport api"}
	PanicError = &Error{500, 500, "server error"}

	ParamError        = &Error{200, 1000, "param error"}
	DBError           = &Error{200, 1030, "unknown error"}
	UnknownError      = &Error{200, 1060, "unknown error"}
	ErrorFlagValError = &Error{200, 1101, "unknown error"}
)
