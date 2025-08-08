package constants

type ErrorCodeType struct {
	Code string
	Msg  string
}

var ErrorCodesMap = map[string]ErrorCodeType{
	"min": {
		Code: "MIN_[X]",
		Msg:  "Minimal [X]",
	},
	"max": {
		Code: "MAX_[X]",
		Msg:  "Maximal [X]",
	},
	"email": {
		Code: "EMAIL_NOT_VALID",
		Msg:  "[F] not valid",
	},
	"required": {
		Code: "REQUIRED",
		Msg:  "[F] required",
	},
}
