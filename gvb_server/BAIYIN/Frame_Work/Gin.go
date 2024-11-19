package FK

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gvb_server/BAIYIN/BY"
	"gvb_server/BAIYIN/Frame_Work/FK_Struct"
	"reflect"
)

func (_Gin) GetValidMsg(c *gin.Context, err error, obj any) bool {
	if err == nil {
		return true
	}
	ptr := reflect.TypeOf(obj)          // obj指针
	var errs validator.ValidationErrors // err接口断言为类型
	var message string
	if errors.As(err, &errs) {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误
			f, exits := ptr.Elem().FieldByName(e.Field())
			if exits {
				message = f.Tag.Get("msg")
				if message == "" {
					json := f.Tag.Get("json")
					if BY.AT.Find(e.Error(), "required") {
						message = json + "为必填字段"
					}
				}
			}
		}
	}
	if message == "" {
		message = "Error: " + err.Error()
	}
	BY.JSON(c, 1000, message)
	BY.LogE("绑定错误: " + err.Error())
	return false
}

// BindJSON 绑定参数   json
func (by _Gin) BindJSON(c *gin.Context, obj any) bool {
	return by.GetValidMsg(c, c.ShouldBindJSON(obj), obj)
}

// BindQuery 绑定参数 ? form
func (by _Gin) BindQuery(c *gin.Context, obj any) bool {
	return by.GetValidMsg(c, c.ShouldBindQuery(obj), obj)
}

// BindURI 绑定参数  /
func (by _Gin) BindURI(c *gin.Context, obj any) bool {
	return by.GetValidMsg(c, c.ShouldBindUri(obj), obj)
}

// GetHead   获取请求头
func (_Gin) GetHead(c *gin.Context, str string) string {
	return c.GetHeader(str)
}

// GetHeader 获取Headers
func (_Gin) GetHeader(c *gin.Context, str string) string {
	return c.Request.Header.Get(str)
}

// Param 获取路由参数   /:id
func (_Gin) Param(c *gin.Context, ID string) string {
	return c.Param(ID)
}

// GetAddr 获取IP  属地
func (_Gin) GetAddr(c *gin.Context) FK_Struct.Addr {
	ip := c.ClientIP()
	addr := BY.Other.GetAddr(ip)
	return FK_Struct.Addr{IP: ip, Addr: addr}
}
