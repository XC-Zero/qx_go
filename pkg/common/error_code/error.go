package ec

/* 为了解决出现错误，无法将程序本身的错误消息和需要返回给前端的错误消息区分开来的问题。定义 Error 结构体，其中包含 LogMessage 表示需要记录到日志文件的错误，实现了 error 接口。
   还包括 DisplayMessage 表示需要返回给前端的错误，实现了 error 接口。
*/

import (
	"fmt"
	"io"
)

// Error DisplayMessage might be error_code.RomeError or error_code.VisaError
type Error struct {
	LogError     error
	DisplayError DisplayError
}

func (r *Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v', 's':
		io.WriteString(s, r.Log())
		return
	case 'q':
		fmt.Fprintf(s, "%q", r.LogError.Error())
	}
}

// Error ...
func (r *Error) Error() string {
	return r.DisplayError.Error() + ";" + r.LogError.Error()
}

// Log return message to be written into log
func (r *Error) Log() string {
	return fmt.Sprintf("%+v", r.LogError)
}

// Display show for front-end
func (r *Error) Display() string {
	return r.DisplayError.Error()
}
