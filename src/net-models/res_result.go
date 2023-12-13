package net_models

var code = map[string]uint{
	"SECCEED": 200,
	"ERROR":   400,
}

// ResResult 响应体
type ResResult struct {
	code uint
	msg  string
	Data interface{}
}

func (response *ResResult) GetData() interface{} {
	return response.Data
}

func ResSucceed(data interface{}) ResResult {
	return ResResult{
		code: code["SECCEED"],
		msg:  "succeed",
		Data: data,
	}
}

func ResErr(msg string, state string) ResResult {
	return ResResult{
		code: code[state],
		msg:  msg,
		Data: nil,
	}
}

func Pending(data interface{}) ResResult {
	return ResResult{
		code: 0,
		msg:  "pending",
		Data: data,
	}
}
