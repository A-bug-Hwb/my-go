package intercept

import (
	"myGo/config"
	"strings"
)

// NotIp 不拦截的ip(ip白名单)
func NotIp(ip string) bool {
	arr := strings.Split(config.IConf.IpWhite, ",")
	if len(arr) == 0 {
		return true
	}
	//默认本机不拦截
	arr = append(arr, "127.0.0.1", "::1")
	for _, element := range arr {
		if ip == element {
			return true
		}
	}
	return false
}

// NotIntercept 判断路由是否拦截
func NotIntercept(path string) bool {
	//不拦截的地址
	noPath := []string{"/api/v1/user/login", "api/v1/user/register"}
	for _, element := range noPath {
		if path == element {
			return true
		}
	}
	return false
}
