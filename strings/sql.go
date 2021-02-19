package strings

import (
	"net/http"
	"strings"
)

// LikeFieldEscape 转义 SQL 的 like 模糊查询时字段值为通配符的值
func LikeFieldEscape(value string) string {
	value = strings.Replace(value, ";", "\\;", -1)
	value = strings.Replace(value, "\"", "\\\"", -1)
	value = strings.Replace(value, "'", "\\'", -1)
	value = strings.Replace(value, "--", "\\--", -1)
	value = strings.Replace(value, "/", "\\/", -1)
	value = strings.Replace(value, "%", "\\%", -1)
	value = strings.Replace(value, "_", "\\_", -1)
	return value
}

//sql字符串过滤
func Guolv(str string) string {
	str = strings.ToLower(str)
	str = strings.Replace(str, `'`, `\'`, -1)
	str = strings.Replace(str, `"`, `\"`, -1)
	str = strings.Replace(str, "`", "\\`", -1)
	str = strings.Replace(str, `;`, `\;`, -1)
	str = strings.Replace(str, `,`, `\,`, -1)
	str = strings.Replace(str, `:`, `\:`, -1)
	str = strings.Replace(str, `:`, `\:`, -1)
	str = strings.Replace(str, `#`, `\#`, -1)
	str = strings.Replace(str, `*`, `\*`, -1)
	str = strings.Replace(str, `/`, `\/`, -1)
	str = strings.Replace(str, `[`, `\[`, -1)
	str = strings.Replace(str, `]`, `\]`, -1)
	str = strings.Replace(str, `=`, `\=`, -1)
	str = strings.Replace(str, `or`, `\or`, -1)
	str = strings.Replace(str, `(`, `\(`, -1)
	return str
}

func PreventSQLInjection(sqlStr string) string {
	sqlStr = strings.Replace(sqlStr, "--", "", -1)
	sqlStr = strings.Replace(sqlStr, "#", "", -1)
	sqlStr = strings.Replace(sqlStr, ";", "", -1)
	sqlStr = strings.Replace(sqlStr, "\n", "", -1)
	sqlStr = strings.Replace(sqlStr, "\r", "", -1)
	sqlStr = strings.Replace(sqlStr, "''", "", -1)
	sqlStr = strings.Replace(sqlStr, "\"", "", -1)

	return sqlStr
}

func IsAjaxRequest(request *http.Request) bool {
	if len(request.Header["X-Requested-With"]) != 0 && request.Header["X-Requested-With"][0] == "XMLHttpRequest" {
		return true
	} else {
		return false
	}
}
