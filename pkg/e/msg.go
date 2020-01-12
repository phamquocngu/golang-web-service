package e

var MsgFlags = map[int]string{
	SUCCESS:                         "SUCCESS",
	ERROR:                           "ERROR",
	INVALID_PARAMS:                  "INVALID PARAMS"
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
