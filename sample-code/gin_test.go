package code

//import (
//	"github.com/marmotedu/errors"
//	code "github.com/marmotedu/sample-code"
//	"net/http"
//)
//
//// Response defines project response format which in marmotedu organization.
//type Response struct {
//	Code      errors.Code `json:"code,omitempty"`
//	Message   string      `json:"message,omitempty"`
//	Reference string      `json:"reference,omitempty"`
//	Data      interface{} `json:"data,omitempty"`
//}
//
//// WriteResponse used to write an error and JSON data into response.
//func WriteResponse(c *gin.Context, err error, data interface{}) {
//	if err != nil {
//		coder := errors.ParseCoder(err)
//
//		c.JSON(coder.HTTPStatus(), Response{
//			Code:      coder.Code(),
//			Message:   coder.String(),
//			Reference: coder.Reference(),
//			Data:      data,
//		})
//	}
//
//	c.JSON(http.StatusOK, Response{Data: data})
//}
//
//func GetUser(c *gin.Context) {
//	log.Info("get user function called.", "X-Request-Id", requestid.Get(c))
//	// Get the user by the `username` from the database.
//	user, err := store.Client().Users().Get(c.Param("username"), metav1.GetOptions{})
//	if err != nil {
//		core.WriteResponse(c, errors.WithCode(code.ErrUserNotFound, err.Error()), nil)
//		return
//	}
//
//	core.WriteResponse(c, nil, user)
//}

// 通过 WriteResponse 统一处理错误：
// 如果 err != nil，则从 error 中解析出 Coder，并调用 Coder 提供的方法，、
// 获取错误相关的 Http Status Code、int 类型的业务码、暴露给用户的信息、错误的参考文档链接，并返回 JSON 格式的信息。
// 如果 err == nil 则返回 200 和数据。
