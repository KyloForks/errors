package code

import (
	"github.com/marmotedu/errors"
	"github.com/novalagung/gubrak"
)

// ErrCode implements `github.com/marmotedu/errors`.Coder interface.
// 一种错误类型包括业务码、HTTP Status Code、对外展示的 Message，以及跟该错误匹配的帮助文档。
// 使用 MustRegister 或 Register 注册。
type ErrCode struct {
	// C refers to the code of the ErrCode.
	C int

	// HTTP status that should be used for the associated error code.
	HTTP int

	// External (user) facing error text.
	Ext string

	// Ref specify the reference document.
	Ref string
}

// Code returns the integer code of ErrCode.
func (coder ErrCode) Code() int {
	return coder.C
}

// String implements stringer. String returns the external error message,
// if any.
func (coder ErrCode) String() string {
	return coder.Ext
}

// Reference returns the reference document.
func (coder ErrCode) Reference() string {
	return coder.Ref
}

// HTTPStatus returns the associated HTTP status code, if any. Otherwise,
// returns 200.
func (coder ErrCode) HTTPStatus() int {
	if coder.HTTP == 0 {
		return 500
	}
	return coder.HTTP
}

func register(code int, httpStatus int, message string, refs ...string) {
	found, _ := gubrak.Includes([]int{200, 400, 401, 403, 404, 500}, httpStatus)
	if !found {
		panic("http code not in `200, 400, 401, 403, 404, 500`")
	}

	var reference string
	if len(refs) > 0 {
		reference = refs[0]
	}

	coder := &ErrCode{
		C:    code,
		HTTP: httpStatus,
		Ext:  message,
		Ref:  reference,
	}

	errors.MustRegister(coder)
}
