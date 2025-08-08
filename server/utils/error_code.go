package utils

import (
	"strings"

	"github.com/krazy-code/devlink/constants"
)

// ErrorCode returns a dynamic error code string for a given key.
func ErrorCode(tag string, p string, f string) string {

	if codeMap, ok := constants.ErrorCodesMap[tag]; ok {
		if strings.Contains(codeMap.Code, "_[X]") {
			return strings.Replace(codeMap.Code, "_[X]", "_"+p, 1)
		}
		if strings.Contains(codeMap.Code, "[F]") {
			return strings.Replace(codeMap.Code, "[F]", strings.ToUpper(f)+"_", 1)
		}

		return codeMap.Code
	}
	return "ERR_UNKNOWN"
}
