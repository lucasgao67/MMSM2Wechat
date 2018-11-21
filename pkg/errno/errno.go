package errno

import (
	"fmt"
)

// 通过 errno.go 来对自定义的错误进行处理
type Errno struct {
	Code    int
	Message string
}

// 结构方法
func (err Errno) Error() string {
	return err.Message
}

type Err struct {
	Code    int
	Message string
	Err     error
}

// 结构方法
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func New(errno *Errno, err error) *Err {
	return &Err{errno.Code, errno.Message, err}
}

func (err *Err) Add(message string) error {
	err.Message += "" + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += "" + fmt.Sprintf(format, args)
	return err
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}
	return InternalServerError.Code, err.Error()
}
