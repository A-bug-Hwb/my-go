package result

import "net/http"

// SUCCESS 成功
const SUCCESS = http.StatusOK

// BAD_REQUEST 参数有误
const BAD_REQUEST = http.StatusBadRequest

// UNAUTHORIZED 未授权
const UNAUTHORIZED = http.StatusUnauthorized

// FORBIDDEN 授权过期
const FORBIDDEN = http.StatusForbidden

// IP_NOT_AUTHORIZED ip禁止访问
const IP_NOT_AUTHORIZED = http.StatusProxyAuthRequired

// ERROR 系统错误
const ERROR = http.StatusInternalServerError

// WARN 系统警告
const WARN = 601
