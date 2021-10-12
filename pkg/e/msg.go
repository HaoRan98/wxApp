package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	INVALID_PARSE_FORM:    "解析绑定表单错误",
	INVALID_PARAMS_VERIFY: "参数校验错误",

	ERROR_USERNAME_PASSWORD:        "用户名密码不正确",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时,请重新登录",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	ERROR_UPLOAD_SAVE_FILE_FAIL:    "保存文件失败",
	ERROR_UPLOAD_CHECK_FILE_FAIL:   "检查文件失败",
	ERROR_UPLOAD_CHECK_FILE_FORMAT: "校验文件错误，文件格式不正确",
	ERROR_UPLOAD_CHECK_FILE_SIZE:   "校验文件错误，文件大小超出限制",

	ERROR_USERNAME_EXIST: "用户名已存在",

	ERROR_GET_DEPARTMENT_FAIL: 	"获取部门列表失败",
	ERROR_GET_USER_FAIL:       	"获取部门用户列表失败",

	WX_LOGIN_ERR:				"登陆请求失败",
	WRITE_IN_FAIL:				"数据写入失败",

	ADDRESS_CREATE_ERROR:      	"地址创建错误",
	ADDRESS_GET_ERROR:      	"地址获取错误",

	ORDER_CREATE_ERROR:      	"订单创建失败",
	ORDER_GET_ERROR:      		"订单获取错误",
	ORDER_EDIT_ERROR:      		"订单修改错误",
	ORDER_DELETE_ERROR:      	"订单删除错误",


	PRODUCT_CREATE_ERROR:      	"商品上传失败",
	PRODUCT_EDIT_ERROR:      	"商品修改失败",
	PRODUCT_GET_ERROR:      	"商品获取错误",
	PRODUCT_DELETE_ERROR:      	"商品删除错误",


	QRCODE_GET_ERROR:      		"获取二维码失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
