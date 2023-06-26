package errorx

var ErrUnKnow = NewError(500, "UN_KNOW", "未知错误, 请联系开发团队")
var ErrBadRequest = NewError(400, "BAD_REQUEST", "违规请求")

var ErrUnAuthorization = NewError(401, "UN_AUTHORIZATION", "未授权")
var ErrPhoneUnAuthorization = NewError(401, "UN_PHONE_AUTHORIZATION", "用户需要先授权登录")
